package config

import (
	"os"
	"path/filepath"
	"runtime"
	"strconv"

	"github.com/joho/godotenv"
)

type AppServerModel struct {
	IP   string
	Port string
}

type AppDBModel struct {
	DbHost string
	DbPort string
	DbUser string
	DbPass string
	DbName string
}

type AppJwtModel struct {
	Secret string
	Expiry int
}

type AppConfigModel struct {
	Server AppServerModel
	Db     AppDBModel
	Jwt    AppJwtModel
}

func Load() (*AppConfigModel, error) {

	// Get the path to the current file (server/dev/main.go)
	_, currentFile, _, _ := runtime.Caller(0)
	currentDir := filepath.Dir(currentFile)

	// Navigate up to project root (adjust the number of ".." as needed)
	projectRoot := filepath.Join(currentDir, "..", "..")
	envPath := filepath.Join(projectRoot, ".env")

	// Load the .env file
	err := godotenv.Load(envPath)
	if err != nil {
		return nil, err
	}

	server := AppServerModel{
		IP:   os.Getenv("APP_IP"),
		Port: os.Getenv("APP_PORT"),
	}

	db := AppDBModel{
		DbHost: os.Getenv("DB_HOST"),
		DbPort: os.Getenv("DB_PORT"),

		DbUser: os.Getenv("DB_USER"),
		DbPass: os.Getenv("DB_PASSWORD"),
		DbName: os.Getenv("DB_NAME"),
	}

	expiry, _ := strconv.ParseInt(os.Getenv("JWT_EXPIRY"), 10, 64)

	jwt := AppJwtModel{
		Secret: os.Getenv("JWT_SECRET"),
		Expiry: int(expiry),
	}

	appConfig := AppConfigModel{
		Server: server, Db: db, Jwt: jwt,
	}

	return &appConfig, nil
}
