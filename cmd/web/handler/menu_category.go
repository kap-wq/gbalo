package handler

import (
	"app/cmd/web/view"
	"app/cmd/web/viewmodel"
	"net/http"

	"github.com/socle-lab/render"
)

func (h *Handler) Menu_categoryHandler(w http.ResponseWriter, r *http.Request) {

	var err error

	vm := viewmodel.NewIndexViewModel("Modules de recherche", nil)

	// err = h.render(w, r, render.PageOptions{
	// 	Data: views.Home(vm),
	// })

	err = h.render(w, r, render.PageOptions{
		ComponentFunc: view.Menu_category,
		ViewModel:     &vm,
		Data:          nil,
	})
	if err != nil {
		h.log("error", err)
	}
}
