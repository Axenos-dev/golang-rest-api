package save_code

import (
	config_auth_db "mazano-server/src/internal/app/auth/config_db"
	"mazano-server/src/internal/app/connect_to_mongodb"
	"mazano-server/src/internal/app/models"
)

type filter struct {
	Email string
}

func Save_Code(code models.ValidationCode) {
	config := config_auth_db.GetConfig("recover-password")

	collection := connect_to_mongodb.CreateConnection(config)

	collection.DeleteData(filter{Email: code.Email})
	collection.InsertData(code)
}
