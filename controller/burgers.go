package controller

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"burgers-api/entity"
	"burgers-api/service"

	"github.com/gorilla/mux"
)

func FetchBurgers(service service.BurgerServiceInterface) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		res, err := service.Fetch()
		if err != nil {
			WriteResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		WriteResponse(w, http.StatusOK, res)

	}
}

func GetBurger(service service.BurgerServiceInterface) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		name := params["name"]

		res, err := service.GetBurger(name)
		if err != nil {
			WriteResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		WriteResponse(w, http.StatusOK, res)
	}
}

func GetRandomBurger(service service.BurgerServiceInterface) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res, err := service.GetRandomBurger()
		if err != nil {
			WriteResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		WriteResponse(w, http.StatusOK, res)
	}
}

func CreateBurger(service service.BurgerServiceInterface) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		burger := entity.Burger{}

		body, err := ioutil.ReadAll(r.Body)

		if err != nil {
			WriteResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		err = json.Unmarshal(body, &burger)
		if err != nil {
			WriteResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		res, err := service.Create(burger)
		if err != nil {
			WriteResponse(w, http.StatusBadRequest, err.Error())
			return
		}
		WriteResponse(w, http.StatusOK, res)
	}
}

func WriteResponse(w http.ResponseWriter, status int, res interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(res)
}
