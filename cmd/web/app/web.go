package app

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"

	"github.com/nialdeco98/myShop/internal/models"
)

// Config - struct containing configuration prefrences
type Config struct {
	Port int
	Env  string
	Api  string
	Db   struct {
		Dsn string
	}
	Stripe struct {
		Secret string
		Key    string
	}
	InfoLog  *log.Logger
	ErrorLog *log.Logger
}

type application struct {
	Config
	TemplateCache map[string]*template.Template
	Version       string
	DB            models.DBModel
}

// New get a new application instance
func New(cfg Config) *application {
	return &application{
		Config: cfg,
	}
}

func (app *application) Serve() error {
	srv := &http.Server{
		Addr:              fmt.Sprintf(":%d", app.Port),
		Handler:           app.routes(),
		IdleTimeout:       30 * time.Second,
		ReadTimeout:       10 * time.Second,
		ReadHeaderTimeout: 5 * time.Second,
		WriteTimeout:      5 * time.Second,
	}

	app.InfoLog.Printf("Starting Frontend server in \"%s\" mode on port %d\n", app.Env, app.Port)
	return srv.ListenAndServe()
}
