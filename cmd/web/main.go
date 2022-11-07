package main

import (
	"flag"
	"html/template"
	"log"
	"os"

	"github.com/nialdeco98/myShop/cmd/web/app"
)

const (
	version    = "1.0.0"
	cssVersion = "1"
)

func main() {
	var cfg app.Config

	flag.IntVar(&cfg.Port, "port", 4000, "Server port to listen on")
	flag.StringVar(&cfg.Env, "env", "dev", "Application enviroment {dev|prod}")
	flag.StringVar(&cfg.Api, "api", "http://localhost:4001", "Url to api")

	flag.Parse()

	cfg.Stripe.Key = os.Getenv("STRIPE_KEY")
	cfg.Stripe.Secret = os.Getenv("STRIPE_SECRET")

	cfg.InfoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	cfg.ErrorLog = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	app := app.New(cfg)

	app.TemplateCache = make(map[string]*template.Template)
	app.Version = version

	if err := app.Serve(); err != nil {
		app.ErrorLog.Println(err)
		log.Fatal()
	}
}
