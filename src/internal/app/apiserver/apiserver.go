package apiserver

import (
	"encoding/json"
	"fmt"
	"mazano-server/src/internal/app/api/search"
	"mazano-server/src/internal/app/models"

	"mazano-server/src/internal/app/recommendations/get_games"
	"mazano-server/src/internal/app/recommendations/get_movies"
	"mazano-server/src/internal/app/recommendations/get_series"

	decode_request "mazano-server/src/internal/app/request_decoder"
	"net/http"

	"mazano-server/src/internal/app/auth/change_password"
	"mazano-server/src/internal/app/auth/change_username"
	"mazano-server/src/internal/app/auth/get_profile"
	"mazano-server/src/internal/app/auth/recover_password/confirm_code"
	"mazano-server/src/internal/app/auth/recover_password/send_email"
	"mazano-server/src/internal/app/auth/set_avatar"
	"mazano-server/src/internal/app/auth/sign_in"
	"mazano-server/src/internal/app/auth/sign_up"

	"github.com/gorilla/mux"
)

type APIServer struct {
	config *Config
	router *mux.Router
}

func New(config *Config) *APIServer {
	return &APIServer{
		config: config,
		router: mux.NewRouter(),
	}
}

func (server *APIServer) Start() error {
	server.ConfigRouter()

	fmt.Println("Starting API-Server " + server.config.BindAddr)

	return http.ListenAndServe(server.config.BindAddr, server.router)
}

func (server *APIServer) ConfigRouter() {
	server.router.HandleFunc("/auth/register", server.HandleSignUp())
	server.router.HandleFunc("/auth/login", server.HandleSignIn())
	server.router.HandleFunc("/auth/profile", server.HandleGetProfile())

	server.router.HandleFunc("/auth/send-email", server.HandleSendEmail())
	server.router.HandleFunc("/auth/verify-code", server.HandleValidateCode())
	server.router.HandleFunc("/auth/change-password", server.HandleChangePassword())

	server.router.HandleFunc("/auth/set-avatar", server.HandleSetAvatar())
	server.router.HandleFunc("/auth/change-username", server.HandleChangeUsername())

	server.router.HandleFunc("/api/search", server.HandleSearch())

	server.router.HandleFunc("/recommendations/movies", server.HandleGetMovies())
	server.router.HandleFunc("/recommendations/series", server.HandleGetSeries())
	server.router.HandleFunc("/recommendations/games", server.HandleGetGames())

}

func (server *APIServer) HandleSignUp() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "POST":
			{
				var data models.NewUserRequestData
				decode_request.Decode_Request(r, &data)

				res, _ := json.Marshal(decode_request.SnakeCaseEncode(sign_up.Sign_Up(data)))
				w.Write(res)
			}
		}
	}
}

func (server *APIServer) HandleSignIn() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "POST":
			{
				var data models.LoginUserRequestData
				decode_request.Decode_Request(r, &data)

				res, _ := json.Marshal(decode_request.SnakeCaseEncode(sign_in.Sign_In(data)))
				w.Write(res)
			}
		}
	}
}

func (server *APIServer) HandleSendEmail() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "POST":
			{
				var data models.SendEmailRequest
				decode_request.Decode_Request(r, &data)

				res, _ := json.Marshal(decode_request.SnakeCaseEncode(send_email.Send_Email(data)))
				w.Write(res)
			}
		}
	}
}

func (server *APIServer) HandleValidateCode() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "POST":
			{
				token := r.URL.Query().Get("token")

				var data models.ConfirmCodeRequest
				decode_request.Decode_Request(r, &data)

				data.Token = token

				res, _ := json.Marshal(decode_request.SnakeCaseEncode(confirm_code.ConfirmCode(data)))
				w.Write(res)
			}
		}
	}
}

func (server *APIServer) HandleChangePassword() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "POST":
			{
				token := r.URL.Query().Get("token")

				var data models.ChangePasswordRequest
				decode_request.Decode_Request(r, &data)

				data.Token = token

				res, _ := json.Marshal(decode_request.SnakeCaseEncode(change_password.ChangePassword(data)))
				w.Write(res)
			}
		}
	}
}

func (server *APIServer) HandleGetProfile() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			{
				var data models.GetProfileRequest
				decode_request.Decode_Request(r, &data)

				res, _ := json.Marshal(decode_request.SnakeCaseEncode(get_profile.GetProfile(data)))
				w.Write(res)
			}
		}
	}
}

func (server *APIServer) HandleSetAvatar() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "PUT":
			{
				var data models.SetAvatarRequest
				decode_request.Decode_Request(r, &data)

				res, _ := json.Marshal(decode_request.SnakeCaseEncode(set_avatar.SetAvatar(data)))
				w.Write(res)
			}
		}
	}
}

func (server *APIServer) HandleChangeUsername() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "PUT":
			{
				var data models.ChangeUsernameRequest
				decode_request.Decode_Request(r, &data)

				res, _ := json.Marshal(decode_request.SnakeCaseEncode(change_username.ChangeUsername(data)))
				w.Write(res)
			}
		}
	}
}

func (server *APIServer) HandleSearch() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			{
				var data models.SearchRequest
				decode_request.Decode_Request(r, &data)

				res, _ := json.Marshal(decode_request.SnakeCaseEncode(search.Search(data)))
				w.Write(res)
			}
		}
	}
}

func (server *APIServer) HandleGetMovies() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			{
				var data models.GetMoviesRequest
				decode_request.Decode_Request(r, &data)

				res, _ := json.Marshal(decode_request.SnakeCaseEncode(get_movies.GetMovies(data)))
				w.Write(res)
			}
		}
	}
}

func (server *APIServer) HandleGetSeries() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			{
				var data models.GetSeriesRequest
				decode_request.Decode_Request(r, &data)

				res, _ := json.Marshal(decode_request.SnakeCaseEncode(get_series.GetSeries(data)))
				w.Write(res)
			}
		}
	}
}

func (server *APIServer) HandleGetGames() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			{
				var data models.GetGamesRequest
				decode_request.Decode_Request(r, &data)

				res, _ := json.Marshal(decode_request.SnakeCaseEncode(get_games.GetGames(data)))
				w.Write(res)
			}
		}
	}
}
