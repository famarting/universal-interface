package main

import (
	"log"
	"net/http"

	"github.com/famartinrh/uninterface/pkg/config"
	"github.com/famartinrh/uninterface/pkg/handlers"
	"github.com/gorilla/mux"
)

func runServer(appConfig *config.Config) error {

	r, err := makeRouter(appConfig)
	if err != nil {
		return err
	}

	var addr string = ":7878"
	log.Println("Running server on addr " + addr)
	return http.ListenAndServe(addr, r)
}

func makeRouter(appConfig *config.Config) (*mux.Router, error) {

	docstoresHandler, err := handlers.NewDocumentStoreHandler(appConfig)
	if err != nil {
		return nil, err
	}

	r := mux.NewRouter()

	apiRouter := r.PathPrefix("/api").Subrouter()

	// /api/documentstores
	documentstoresV1Router := apiRouter.PathPrefix("/documentstores/v1").Subrouter()

	documentstoresV1Router.HandleFunc("/{store}", docstoresHandler.Get).Queries("idKey", "", "idValue", "").Methods(http.MethodGet)
	documentstoresV1Router.HandleFunc("/{store}/{id}", docstoresHandler.Get).Methods(http.MethodGet)

	documentstoresV1Router.HandleFunc("/{store}", docstoresHandler.Delete).Queries("idKey", "", "idValue", "").Methods(http.MethodDelete)
	documentstoresV1Router.HandleFunc("/{store}/{id}", docstoresHandler.Delete).Methods(http.MethodDelete)

	documentstoresV1Router.HandleFunc("/{store}", docstoresHandler.Create).Methods(http.MethodPost)
	// documentstoresV1Router.HandleFunc("", docstoresHandler.List).Methods(http.MethodGet)

	return r, nil
}
