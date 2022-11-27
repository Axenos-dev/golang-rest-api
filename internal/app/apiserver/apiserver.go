package apiserver

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	recoverpassword "mazano-server/internal/app/recover_password"
	"mazano-server/internal/app/registration"
	search_main "mazano-server/internal/app/search"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
)

type APIServer struct {
	config *Config
	logger *logrus.Logger
	router *mux.Router
}

type Handler func(bson.M) bson.M

func New(config *Config) *APIServer {
	return &APIServer{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}

func (s *APIServer) Start() error {
	if err := s.ConfigLogger(); err != nil {
		return err
	}

	s.ConfigRouter()

	s.logger.Info("Starting API-Server " + s.config.BindAddr)

	return http.ListenAndServe(s.config.BindAddr, s.router)
}

func (s *APIServer) ConfigLogger() error {
	level, err := logrus.ParseLevel(s.config.LogLevel)

	if err != nil {
		return err
	}

	s.logger.SetLevel(level)

	return nil
}

func (s *APIServer) ConfigRouter() {
	s.router.HandleFunc("/auth/register", s.HandleRegister())
	s.router.HandleFunc("/auth/login", s.HandleLogin())

	s.router.HandleFunc("/auth/recover-password", s.HandleSendEmail())
	s.router.HandleFunc("/auth/verify-code", s.HandleValidCode())
	s.router.HandleFunc("/auth/change-password", s.HandleChangePassword())

	s.router.HandleFunc("/auth/profile", s.HandleGetProfile())
	s.router.HandleFunc("/auth/change-username", s.HandleChangeUsername())
	s.router.HandleFunc("/auth/set-avatar", s.HandleChangeAvatar())

	s.router.HandleFunc("/api/search", s.HandleSearch())
}

func HandleResponse(w http.ResponseWriter, r *http.Request, method Handler, extra_params [2]string) {

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	var data bson.M
	bson.UnmarshalExtJSON(body, true, &data)

	if extra_params[0] != "none" {
		data["request_type"] = extra_params[0]
		data["method"] = extra_params[1]
	}
	res, _ := json.Marshal(method(data))
	w.Write(res)
}

func (s *APIServer) HandleRegister() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		switch r.Method {
		case "POST":
			{
				HandleResponse(w, r, registration.MongoRequest, [2]string{"INSERT", "register"})
			}
		}
	}
}

func (s *APIServer) HandleLogin() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "POST":
			{
				HandleResponse(w, r, registration.MongoRequest, [2]string{"GET", "login"})
			}
		}
	}
}

func (s *APIServer) HandleSendEmail() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "POST":
			{
				body, err := ioutil.ReadAll(r.Body)
				if err != nil {
					log.Fatal(err)
				}

				var data bson.M
				bson.UnmarshalExtJSON(body, true, &data)

				res, _ := json.Marshal(recoverpassword.MongoRequestRecoverPassword("Send-Email", data))
				fmt.Fprintf(w, string(res))
			}
		}
	}
}

func (s *APIServer) HandleValidCode() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "POST":
			{
				token := r.URL.Query().Get("token")

				body, err := ioutil.ReadAll(r.Body)
				if err != nil {
					log.Fatal(err)
				}

				data := bson.M{}
				bson.UnmarshalExtJSON(body, true, &data)

				data["token"] = token

				res, _ := json.Marshal(recoverpassword.MongoRequestRecoverPassword("Code-Validation", data))
				fmt.Fprintf(w, string(res))
			}
		}
	}
}

func (s *APIServer) HandleChangePassword() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "POST":
			{
				token := r.URL.Query().Get("token")

				body, err := ioutil.ReadAll(r.Body)
				if err != nil {
					log.Fatal(err)
				}

				data := bson.M{}
				bson.UnmarshalExtJSON(body, true, &data)

				data["token"] = token

				res, _ := json.Marshal(recoverpassword.MongoRequestRecoverPassword("Recover-Password", data))
				fmt.Fprintf(w, string(res))
			}
		}
	}
}

func (s *APIServer) HandleGetProfile() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			{
				HandleResponse(w, r, registration.MongoRequest, [2]string{"GET", "profile"})
			}
		}
	}
}

func (s *APIServer) HandleChangeUsername() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "PUT":
			{
				HandleResponse(w, r, registration.MongoRequest, [2]string{"SET", "change-username"})
			}
		}
	}
}

func (s *APIServer) HandleChangeAvatar() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "PUT":
			{
				HandleResponse(w, r, registration.MongoRequest, [2]string{"SET", "set-avatar"})
			}
		}
	}
}

func (s *APIServer) HandleSearch() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			{
				HandleResponse(w, r, search_main.Search_MazanoAPI, [2]string{"none"})
			}
		}
	}
}
