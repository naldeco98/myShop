package main

import (
	"flag"
	"log"
	"os"

	"github.com/TwiN/go-color"
	"github.com/joho/godotenv"
	"github.com/nialdeco98/myShop/cmd/api/api"
	"github.com/nialdeco98/myShop/internal/driver"
	"github.com/nialdeco98/myShop/internal/models"
)

const (
	version = "1.0.0"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	var cfg api.Config

	flag.IntVar(&cfg.Port, "port", 4001, "Server port to listen on")
	flag.StringVar(&cfg.Env, "env", "dev", "Application enviroment {dev|prod|mant}")
	flag.StringVar(&cfg.Db.Dsn, "dsn", "nialdeco98:@tcp(localhost:3306)/widgets?parseTime=true&tls=false", "DSN to connect database")

	flag.Parse()

	cfg.Stripe.Key = os.Getenv("STRIPE_KEY")
	cfg.Stripe.Secret = os.Getenv("STRIPE_SECRET")

	cfg.InfoLog = log.New(os.Stdout, color.InBold("INFO \t"), log.Ldate|log.Ltime)
	cfg.ErrorLog = log.New(os.Stdout, color.InRed("ERROR \t"), log.Ldate|log.Ltime|log.Lshortfile)

	conn, err := driver.OpenDB(cfg.Db.Dsn)
	if err != nil {
		cfg.ErrorLog.Fatalln(err)
	}
	defer conn.Close()

	app := api.New(cfg)

	app.Version = version
	app.DB = models.DBModel{DB: conn}

	if err := app.Serve(); err != nil {
		app.ErrorLog.Println(err)
		log.Fatal()
	}
}
