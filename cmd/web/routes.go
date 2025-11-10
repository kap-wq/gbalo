package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	socle "github.com/socle-lab/core"
	"github.com/socle-lab/core/pkg/env"
)

func (app *application) routes() http.Handler {
	r := chi.NewRouter()
	// middleware must come before any routes
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{env.GetString("CORS_ALLOWED_ORIGIN", "*")},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	r.Route("/", func(r chi.Router) {
		// add routes here
		// r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		// 	w.Header().Set("Content-Type", "application/json")
		// 	w.WriteHeader(http.StatusOK)
		// 	w.Write([]byte(`{"data":"API 📺 Up and Running"}`))
		// })
		r.Get("/", app.Handler.HomeHandler)
	
	})

	r.Route("/restaurant", func(r chi.Router) {
		// add routes here
		// r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		// 	w.Header().Set("Content-Type", "application/json")
		// 	w.WriteHeader(http.StatusOK)
		// 	w.Write([]byte(`{"data":"API 📺 Up and Running"}`))
		// })
		r.Get("/menu", app.Handler.MenuHandler)
		r.Get("/menu_category", app.Handler.Menu_categoryHandler)
		r.Get("/table", app.Handler.TableHandler)
		r.Get("/valeur_stock", app.Handler.Valeur_stockHandler)


		r.Get("/stock", app.Handler.StockHandler)
		r.Get("/ajustement_stock", app.Handler.Ajustement_stockHandler)
		r.Get("/transfert_stock", app.Handler.Transfert_stockHandler)
		r.Get("/inventaire", app.Handler.InventaireHandler)
		r.Get("/new_stock_commande", app.Handler.New_stock_commandeHandler)
		r.Get("/manage_commande_request", app.Handler.Manage_commande_requestHandler)




	})

	r.Route("/personnes", func(r chi.Router) {
	// navigation routes
	r.Route("/navigation", func(r chi.Router) {
		// add routes here
		// r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		// 	w.Header().Set("Content-Type", "application/json")
		// 	w.WriteHeader(http.StatusOK)
		// 	w.Write([]byte(`{"data":"API 📺 Up and Running"}`))
		// })
		r.Get("/client", app.Handler.ClientHandler)
		r.Get("/fournisseur", app.Handler.FournisseurHandler)
		r.Get("/employe", app.Handler.EmployeHandler)
		r.Get("/employe_affectation", app.Handler.Employe_affectationHandler)
		

	})

	
		r.Get("/dashboard", app.Handler.DashboardHandler)
		r.Get("/analytics", app.Handler.AnalyticsHandler)

	})

	// Hotel routes
	r.Route("/hotel", func(r chi.Router) {
		// add routes here
		// r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		// 	w.Header().Set("Content-Type", "application/json")
		// 	w.WriteHeader(http.StatusOK)
		// 	w.Write([]byte(`{"data":"API 📺 Up and Running"}`))
		// })
		r.Get("/rooms", app.Handler.RoomsHandler)
		r.Get("/reservation", app.Handler.ReservationHandler)

	})

	// SignIn routes
	r.Route("/connexion", func(r chi.Router) {
		// add routes here
		r.Get("/signin", app.Handler.SigninHandler)
		r.Post("/signin", app.Handler.SigninPostHandler)
	})

	// static routes
	fileServer := http.FileServer(http.Dir("./public"))
	r.Handle("/public/*", http.StripPrefix("/public", fileServer))

	// routes from socle
	r.Mount("/socle", socle.Routes())

	// routes from web inner admin
	r.Mount("/admin-console", app.AdminRoutes())

	// routes from web inner api
	//r.Mount("/api", app.ApiRoutes())
	return r
}
