package main

import (
	"catering-api/internal/auth"
	"catering-api/internal/config"
	"catering-api/internal/database"
	"catering-api/internal/meal"
	"catering-api/internal/testimonial"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main(){
	cfg := config.Load()

	db, err := database.NewPostgresDB(cfg.DatabaseUrl)
	if err != nil {
		log.Fatal("failed to connect database", err)
	}

	err = db.AutoMigrate(
		&auth.User{},
		&meal.Meal{},
		&testimonial.Testimonial{},
	)
	if err != nil {
		log.Fatal("failed to run migration", err)
	}

	r := chi.NewRouter()
	
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	authRepo := auth.NewRepository(db)
	authService := auth.NewService(authRepo, cfg.JWTSecret, cfg.JWTExpiresHour)
	authHandler := auth.NewHandler(authService)

	authHandler.RegisterRoutes(r)

	mealRepo := meal.NewRepository(db)
	mealService := meal.NewService(mealRepo)
	mealHandler := meal.NewHandler(mealService)

	testimonialRepo := testimonial.NewRepository(db)
	testimonialService := testimonial.NewService(testimonialRepo)
	testimonialHandler := testimonial.NewHandler(testimonialService)

	r.Group(func(r chi.Router) {
		r.Use(auth.Middleware(cfg.JWTSecret))

		mealHandler.RegisterRoutes(r)
		testimonialHandler.RegisterRoutes(r)
	})

	addr := ":" + cfg.AppPort

	log.Println("server running on port " + addr)

	err = http.ListenAndServe(addr, r)
	if err != nil {
		log.Fatal(err)
	}
}