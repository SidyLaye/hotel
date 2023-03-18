package hotel

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

// Config est une structure qui contient les variables d'environnement pour la configuration de la base de données
type Config struct {
	DBHost string
	DBPort string
	DBUser string
	DBPass string
	DBName string
}

func Conf() (*Config, error) {
	// Charger les variables d'environnement à partir d'un fichier .env
	err := godotenv.Load(".env")
	if err != nil {
		return nil, fmt.Errorf("Impossible de charger le fichier .env : %v", err)
	}

	// Créer une nouvelle instance de la structure Config avec les variables d'environnement chargées
	conf := &Config{
		DBHost: os.Getenv("DB_HOST"),
		DBPort: os.Getenv("DB_PORT"),
		DBUser: os.Getenv("DB_USER"),
		DBPass: os.Getenv("DB_PASS"),
		DBName: os.Getenv("DB_NAME"),
	}

	return conf, nil
}
