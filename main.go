package main

import (
	"cars-api/config"
	"cars-api/managers"
	"cars-api/models"
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	"github.com/sirupsen/logrus"
	"log"
	"net/http"
)

type Env struct {
	Configuration config.Config
	CarsManager   managers.CarsManager
}

func main() {
	// Creates a global instance `logger`
	logger := logrus.New()

	// JSONFormatter formats logs into parsable JSON
	logger.Formatter = &logrus.JSONFormatter{
		DisableTimestamp: true,
	}

	// Global instance base on Env struct
	env := &Env{}

	// Fetching database configs from passed JSON
	configuration, err := config.GetConfig()
	if err != nil {
		panic(err.Error())
	}

	// Init a new DB connection
	connection, err := models.ConnectToDB(configuration)
	if err != nil {
		panic(err.Error())
	}

	// Migrating models
	connection.DB.AutoMigrate(
		&models.Car{})
	connection.DB.LogMode(true)
	env.CarsManager = managers.InitCarsManager(connection)

	// Connect to managers env blah-blah

	// Create a router without any middleware
	r := chi.NewRouter()

	// Base middleware stack
	// RequestID is a middleware that injects a `request_id` into context
	// of each request
	r.Use(middleware.RequestID)

	// RealIP is a middleware that sets a http.Request's RemoteAddr to the
	// results of parsing either the x-Forwarded-For header of the x-Real-IP header.
	r.Use(middleware.RealIP)

	// Logger middleware writes a logs the start and the end of each request,
	// along with useful data, what the response status was, and how long
	// it took to return
	r.Use(middleware.Logger)

	// Recoverer middleware recovers from any panic and returns a 500 Code
	// if there was one
	r.Use(middleware.Recoverer)

	// SetContentType is a middleware that forces response Content-Type
	r.Use(render.SetContentType(render.ContentTypeJSON))

	// Routes stack
	r.Route("/api/v1/cars", func(r chi.Router) {
		r.Post("/", env.CreateCar)
		r.Get("/", env.GetCars)
		r.Put("/", env.UpdateCar)
		r.Delete("/{id}", env.DeleteCar)
		r.Get("/{id}", env.GetCar)
	})

	// Start Listening :5432
	fmt.Println("Start listening http at port 8090...")
	err = http.ListenAndServe(":8090", r)
	if err != nil {
		log.Print(err.Error())
		return
	}
}
