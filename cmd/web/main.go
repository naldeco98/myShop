package main

import (
	"flag"
	"html/template"
	"log"
	"os"

	"github.com/TwiN/go-color"
	"github.com/joho/godotenv"
	"github.com/nialdeco98/myShop/cmd/web/app"
	"github.com/nialdeco98/myShop/internal/driver"
	"github.com/nialdeco98/myShop/internal/models"
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

	{
		flag.IntVar(&cfg.Port, "port", 4000, "Server port to listen on")
		flag.StringVar(&cfg.Env, "env", "dev", "Application enviroment {dev|prod}")
		flag.StringVar(&cfg.Db.Dsn, "dsn", "nialdeco98:@tcp(localhost:3306)/widgets?parseTime=true&tls=false", "DSN to connect database")
		flag.StringVar(&cfg.Api, "api", "http://localhost:4001", "Url to api")
		flag.Parse()
	}

	cfg.Stripe.Key = os.Getenv("STRIPE_KEY")
	cfg.Stripe.Secret = os.Getenv("STRIPE_SECRET")

	cfg.InfoLog = log.New(os.Stdout, color.InBold("INFO \t"), log.Ldate|log.Ltime)
	cfg.ErrorLog = log.New(os.Stdout, color.InRed("ERROR \t"), log.Ldate|log.Ltime|log.Lshortfile)

	conn, err := driver.OpenDB(cfg.Db.Dsn)
	if err != nil {
		cfg.ErrorLog.Fatalln(err)
	}
	defer conn.Close()

	application := app.New(cfg)

	application.TemplateCache = make(map[string]*template.Template)
	application.Version = version
	application.DB = models.DBModel{DB: conn}

	if err := application.Serve(); err != nil {
		application.ErrorLog.Println(err)
		log.Fatal()
	}
}
