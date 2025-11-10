package handler

import (
	"app/cmd/web/view"
	"app/cmd/web/viewmodel"
	"net/http"

	"github.com/socle-lab/render"
)

func (h *Handler) SigninHandler(w http.ResponseWriter, r *http.Request) {

	var err error

	vm := viewmodel.NewIndexViewModel("Modules de recherche", nil)

	// err = h.render(w, r, render.PageOptions{
	// 	Data: views.SignIn(vm),
	// })

	err = h.render(w, r, render.PageOptions{
		ComponentFunc: view.SignIn,
		ViewModel:     &vm,
		Data:          nil,
	})
	if err != nil {
		h.log("error", err)
	}
}

func (h *Handler) SigninPostHandler(w http.ResponseWriter, r *http.Request) {
	// Parse form data
	err := r.ParseForm()
	if err != nil {
		h.log("error", err)
		http.Error(w, "Error parsing form", http.StatusBadRequest)
		return
	}

	// Get form values
	email := r.FormValue("email")
	password := r.FormValue("password")
	rememberMe := r.FormValue("remember_me")

	// TODO: Add your authentication logic here
	// For now, we'll just log the values and redirect
	_ = password // Will be used for authentication validation
	h.log("info", "Login attempt:", "email:", email, "remember:", rememberMe)

	// TODO: Validate credentials against database
	// If valid, create session and redirect
	// If invalid, show error message

	// For now, redirect to dashboard (you should add proper authentication)
	http.Redirect(w, r, "/navigation/dashboard", http.StatusSeeOther)
}
