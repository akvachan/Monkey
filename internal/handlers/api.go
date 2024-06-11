package handlers

import (
    "github.com/go-chi/chi"
    chimiddle "github.com/go-chi/chi/middleware"
    "github.com/akvachan/Monkey/internal/middleware"
)

func Handler(r *chi.Max) {
    r.Use(chimiddle.StripSlashes)
    r.Route("/predictions", func(router chi.Router) {
        
        router.Use(middleware.Authorization)

        router.Get("/prediction", GetPrediction)
    })
}
