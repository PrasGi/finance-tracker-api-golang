package helpers

import "github.com/joho/godotenv"

func LoadEnv() {
	err := godotenv.Load(".env")
	PanicIfErrSystem(err)
}
