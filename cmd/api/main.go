package main

import (
	"flag"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/nialdeco98/myShop/cmd/api/api"
)

const (
	version = "1.0.0"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	var cfg api.Config

	flag.IntVar(&cfg.Port, "port", 4001, "Server port to listen on")
	flag.StringVar(&cfg.Env, "env", "dev", "Application enviroment {dev|prod|mant}")

	flag.Parse()

	cfg.Stripe.Key = os.Getenv("STRIPE_KEY")
	cfg.Stripe.Secret = os.Getenv("STRIPE_SECRET")

	cfg.InfoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	cfg.ErrorLog = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	app := api.New(cfg)

	app.Version = version

	if err := app.Serve(); err != nil {
		app.ErrorLog.Println(err)
		log.Fatal()
	}
}
