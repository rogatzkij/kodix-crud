package handler

import (
	"github.com/gorilla/mux"
	"github.com/rogatzkij/kodix-crud/internal/core"
	"net/http"
)

type handlerStore struct {
	core *core.Core
}

func NewRouter(core *core.Core) *mux.Router {
	hs := handlerStore{
		core: core,
	}

	majorRouter := mux.NewRouter()
	majorRouter.StrictSlash(true)

	apiRouter := majorRouter.PathPrefix("/api/v1").Subrouter()

	// /api/v1/autos
	autoRouter := apiRouter.PathPrefix("/autos").Subrouter()
	autoRouter.HandleFunc("/", hs.readAutoHandler).Methods(http.MethodGet)
	autoRouter.HandleFunc("/", hs.createAutoHandler).Methods(http.MethodPost)
	autoRouter.HandleFunc("/{auto_id}", hs.updateAutoHandler).Methods(http.MethodPut)
	autoRouter.HandleFunc("/{auto_id}", hs.deleteAutoHandler).Methods(http.MethodDelete)

	// /api/v1/brands
	brandRouter := apiRouter.PathPrefix("/brands").Subrouter()
	brandRouter.HandleFunc("/", hs.createBrandHandler).Methods(http.MethodPost)
	brandRouter.HandleFunc("/", hs.deleteBrandHandler).Methods(http.MethodDelete)

	brandRouter.HandleFunc("/{brand}/models/{model}", hs.createModelHandler).Methods(http.MethodPost)
	brandRouter.HandleFunc("/{brand}/models/{model}", hs.deleteModelHandler).Methods(http.MethodDelete)

	return majorRouter
}

func (hs *handlerStore) readAutoHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
	w.Write([]byte("readAutoHandler isn't implemented"))
}

func (hs *handlerStore) createAutoHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
	w.Write([]byte("createAutoHandler isn't implemented"))
}

func (hs *handlerStore) updateAutoHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
	w.Write([]byte("updateAutoHandler isn't implemented"))
}

func (hs *handlerStore) deleteAutoHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
	w.Write([]byte("deleteAutoHandler isn't implemented"))
}

func (hs *handlerStore) createBrandHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
	w.Write([]byte("createBrandHandler isn't implemented"))
}

func (hs *handlerStore) deleteBrandHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
	w.Write([]byte("deleteBrandHandler isn't implemented"))
}

func (hs *handlerStore) createModelHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
	w.Write([]byte("createModelHandler isn't implemented"))
}

func (hs *handlerStore) deleteModelHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
	w.Write([]byte("deleteModelHandler isn't implemented"))
}
