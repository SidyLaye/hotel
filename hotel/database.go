package hotel

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

const AllowedCORSDomain = "http://localhost:11111"

var _ = godotenv.Load(".env")

// Load environment variables
var (
	dbUser     = os.Getenv("DB_USER")
	dbPassword = os.Getenv("DB_PASSWORD")
	dbHost     = os.Getenv("DB_HOST")
	dbPort     = os.Getenv("DB_PORT")
	dbName     = os.Getenv("DB_NAME")
)

// Create MySQL Path
var DbPath = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPassword, dbHost, dbPort, dbName)
