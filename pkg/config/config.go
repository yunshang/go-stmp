package config
import (
	"os"
)

// Config mail varialbes.
type Config struct {
	Host     string
	Port     string
	Password string
	Sender   string
	To       string
	Subject  string
	Body     string
}

// New a mail config
func New() Config {
	port := os.Getenv("PORT")

	host := os.Getenv("HOST")

	password := os.Getenv("PASSWORD")

	return Config{
		Port:     port,
		Host:     host,
		Password: password,
	}
}
