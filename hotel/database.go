package hotel

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

const AllowedCORSDomain = "http://localhost:"

var _ = godotenv.Load(".env")

// Load environment variables
var (
	dbUser     = os.Getenv("USER")
	dbPassword = os.Getenv("PASSWORD")
	dbHost     = os.Getenv("HOST")
	dbPort     = os.Getenv("PORT")
	dbName     = os.Getenv("DB_NAME")
)

// Create MySQL Path
var DbPath = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPassword, dbHost, dbPort, dbName)
