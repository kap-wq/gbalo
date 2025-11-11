package handler

import (
	"app/cmd/web/view"
	"app/cmd/web/viewmodel"
	"net/http"

	"github.com/socle-lab/render"
)

func (h *Handler) DashboardHandler(w http.ResponseWriter, r *http.Request) {

	var err error

	vm := viewmodel.NewIndexViewModel("Modules de recherche", nil)

	// err = h.render(w, r, render.PageOptions{
	// 	Data: views.Dashboard(vm),
	// })

	err = h.render(w, r, render.PageOptions{
		ComponentFunc: view.Dashboard,
		ViewModel:     &vm,
		Data:          nil,
	})
	if err != nil {
		h.log("error", err)
	}
}
