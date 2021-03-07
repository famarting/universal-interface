package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/famartinrh/uninterface/pkg/config"
	"github.com/famartinrh/uninterface/pkg/documentstore"
	"github.com/gorilla/mux"
	"gocloud.dev/gcerrors"
)

type documentStoreHandler struct {
	stores map[string]*documentstore.DocumentStore
}

func NewDocumentStoreHandler(appConfig *config.Config) (*documentStoreHandler, error) {
	var docstoresConfig map[string]*documentstore.DocumentStore = make(map[string]*documentstore.DocumentStore)
	fmt.Println("Configuring documentstores...")
	for _, store := range appConfig.Docstores {
		docstore, err := documentstore.NewDocumentStore(store)
		if err != nil {
			return nil, err
		}
		fmt.Println("New documentstore: " + store.Name)
		docstoresConfig[store.Name] = docstore
	}
	return &documentStoreHandler{
		stores: docstoresConfig,
	}, nil
}

func (h documentStoreHandler) Create(w http.ResponseWriter, r *http.Request) {

	storeName := mux.Vars(r)["store"]

	store := h.stores[storeName]
	if store == nil {
		handleError(r.Context(), w, httpError{err: nil, code: 404, message: "Store not found"})
		return
	}

	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		handleError(r.Context(), w, httpError{err: err, code: 400, message: "Error reading request body"})
		return
	}

	fmt.Println(string(bytes))

	var document map[string]interface{} = make(map[string]interface{})
	err = json.Unmarshal(bytes, &document)
	if err != nil {
		fmt.Println(err)
		handleError(r.Context(), w, httpError{err: err, code: 400, message: "Error unmarshalling document"})
		return
	}

	err = store.Create(document)
	if err != nil {
		fmt.Println(err)
		handleError(r.Context(), w, httpError{err: err, code: 500, message: "Error creating document"})
		return
	}
	writeJSONResponse(w, 200, document)
}

func (h documentStoreHandler) Get(w http.ResponseWriter, r *http.Request) {
	h.handleByIdOperation(w, r, func(store *documentstore.DocumentStore, document map[string]interface{}) (interface{}, *httpError) {
		err := store.Get(document)
		if err != nil {
			errorCode := gcerrors.Code(err)
			code := 500
			if errorCode == gcerrors.NotFound {
				code = 404
			}
			return nil, &httpError{err: err, code: code, message: ""}
		}
		return document, nil
	})
}

func (h documentStoreHandler) Delete(w http.ResponseWriter, r *http.Request) {
	h.handleByIdOperation(w, r, func(store *documentstore.DocumentStore, document map[string]interface{}) (interface{}, *httpError) {
		err := store.Delete(document)
		if err != nil {
			errorCode := gcerrors.Code(err)
			code := 500
			if errorCode == gcerrors.NotFound {
				code = 404
			}
			return nil, &httpError{err: err, code: code, message: ""}
		}
		return nil, nil
	})
}

type operation func(store *documentstore.DocumentStore, document map[string]interface{}) (interface{}, *httpError)

func (h documentStoreHandler) handleByIdOperation(w http.ResponseWriter, r *http.Request, op operation) {
	storeName := mux.Vars(r)["store"]

	store := h.stores[storeName]
	if store == nil {
		handleError(r.Context(), w, httpError{err: nil, code: 404, message: "Store not found"})
		return
	}

	idKey := r.FormValue("idKey")
	if idKey == "" {
		idKey = store.IDField
	}
	if idKey == "" {
		handleError(r.Context(), w, httpError{err: nil, code: 400, message: "Missing idField"})
		return
	}

	idValue := mux.Vars(r)["id"]
	if idValue == "" {
		idValue = r.FormValue("idValue")
	}
	if idValue == "" {
		handleError(r.Context(), w, httpError{err: nil, code: 400, message: "Missing idValue"})
		return
	}

	var iddoc map[string]interface{} = make(map[string]interface{})
	iddoc[idKey] = idValue

	fmt.Println(iddoc)

	result, httpErr := op(store, iddoc)
	if httpErr != nil {
		fmt.Println(httpErr)
		handleError(r.Context(), w, *httpErr)
		return
	}
	writeJSONResponse(w, 200, result)
}
