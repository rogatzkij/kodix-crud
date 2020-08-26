package handler

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/rogatzkij/kodix-crud/internal/core"
	"github.com/rogatzkij/kodix-crud/model"
	"github.com/rs/zerolog/log"
	"io/ioutil"
	"net/http"
	"strconv"
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
	majorRouter.Use(LogMiddleware)

	apiRouter := majorRouter.PathPrefix("/api/v1").Subrouter()

	// /api/v1/autos
	autoRouter := apiRouter.PathPrefix("/autos").Subrouter()
	autoRouter.HandleFunc("/", hs.readAutosHandler).Methods(http.MethodGet)
	autoRouter.HandleFunc("/", hs.createAutoHandler).Methods(http.MethodPost)
	autoRouter.HandleFunc("/{auto_id}", hs.readAutoByIDHandler).Methods(http.MethodGet)
	autoRouter.HandleFunc("/{auto_id}", hs.updateAutoHandler).Methods(http.MethodPut)
	autoRouter.HandleFunc("/{auto_id}", hs.deleteAutoHandler).Methods(http.MethodDelete)

	// /api/v1/brands
	brandRouter := apiRouter.PathPrefix("/brands").Subrouter()
	brandRouter.HandleFunc("/", hs.createBrandHandler).Methods(http.MethodPost)
	brandRouter.HandleFunc("/{brandname}", hs.deleteBrandHandler).Methods(http.MethodDelete)

	brandRouter.HandleFunc("/{brandname}/models/{automodel}", hs.createModelHandler).Methods(http.MethodPost)
	brandRouter.HandleFunc("/{brandname}/models/{automodel}", hs.deleteModelHandler).Methods(http.MethodDelete)

	return majorRouter
}

func LogMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Info().
			Str("path", r.URL.Path).
			Str("method", r.Method).
			Str("remote addr", r.RemoteAddr).
			Str("user agent", r.UserAgent()).
			Msg("Поступил запрос")

		h.ServeHTTP(w, r)
	})
}

func (hs *handlerStore) readAutosHandler(w http.ResponseWriter, r *http.Request) {
	var err error
	var offset uint64 = 0
	var limit uint64 = 10

	if param := r.URL.Query().Get("limit"); param != "" {
		limit, err = strconv.ParseUint(param, 10, 64)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	}

	if param := r.URL.Query().Get("offset"); param != "" {
		offset, err = strconv.ParseUint(param, 10, 64)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	}

	autos, err := hs.core.Auto.GetAutos(uint(limit), uint(offset))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	body, err := json.Marshal(autos)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(body)
}

func (hs *handlerStore) readAutoByIDHandler(w http.ResponseWriter, r *http.Request) {
	autoID, err := strconv.ParseUint(mux.Vars(r)["auto_id"], 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	auto, err := hs.core.Auto.GetAutoByID(uint(autoID))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	body, err := json.Marshal(auto)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(body)
}

func (hs *handlerStore) createAutoHandler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	auto := &model.Auto{}
	err = json.Unmarshal(body, auto)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	_, err = hs.core.Auto.CreateAuto(*auto)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (hs *handlerStore) updateAutoHandler(w http.ResponseWriter, r *http.Request) {
	autoID, err := strconv.ParseUint(mux.Vars(r)["auto_id"], 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	auto := &model.Auto{}
	err = json.Unmarshal(body, auto)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = hs.core.Auto.UpdateAutoByID(uint(autoID), *auto)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (hs *handlerStore) deleteAutoHandler(w http.ResponseWriter, r *http.Request) {
	autoID, err := strconv.ParseUint(mux.Vars(r)["auto_id"], 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = hs.core.Auto.DeleteAutoByID(uint(autoID))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (hs *handlerStore) createBrandHandler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	brand := &model.Brand{}
	err = json.Unmarshal(body, brand)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = hs.core.Brand.CreateBrand(*brand)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (hs *handlerStore) deleteBrandHandler(w http.ResponseWriter, r *http.Request) {
	brand := mux.Vars(r)["brandname"]

	err := hs.core.Brand.DeleteBrand(brand)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (hs *handlerStore) createModelHandler(w http.ResponseWriter, r *http.Request) {
	brandname := mux.Vars(r)["brandname"]
	automodel := mux.Vars(r)["automodel"]

	err := hs.core.Brand.CreateModel(brandname, automodel)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (hs *handlerStore) deleteModelHandler(w http.ResponseWriter, r *http.Request) {
	brandname := mux.Vars(r)["brandname"]
	automodel := mux.Vars(r)["automodel"]

	err := hs.core.Brand.DeleteModel(brandname, automodel)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}
