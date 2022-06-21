package enviroment

import (
	"os"

	"github.com/joho/godotenv"
)

var DB_URI string
var PORT string

func GetEnvVars() {
	err := godotenv.Load(".env")

	DB_URI = os.Getenv("DB_URI")
	PORT = os.Getenv("PORT")

	if err != nil {
		panic(err)
	}
}
