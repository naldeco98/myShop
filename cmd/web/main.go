package main

import (
	"flag"
	"html/template"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/nialdeco98/myShop/cmd/web/app"
)

const (
	version    = "1.0.0"
	cssVersion = "1"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	var cfg app.Config

	flag.IntVar(&cfg.Port, "port", 4000, "Server port to listen on")
	flag.StringVar(&cfg.Env, "env", "dev", "Application enviroment {dev|prod}")
	flag.StringVar(&cfg.Api, "api", "http://localhost:4001", "Url to api")

	flag.Parse()

	cfg.Stripe.Key = os.Getenv("STRIPE_KEY")
	cfg.Stripe.Secret = os.Getenv("STRIPE_SECRET")

	cfg.InfoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	cfg.ErrorLog = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	application := app.New(cfg)

	application.TemplateCache = make(map[string]*template.Template)
	application.Version = version

	if err := application.Serve(); err != nil {
		application.ErrorLog.Println(err)
		log.Fatal()
	}
}
