package main

import (
	"burgers-api/config"
	"burgers-api/controller"
	"burgers-api/database"
	"burgers-api/service"
	"context"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	conf := config.GetConfig()
	ctx := context.TODO()

	db := database.ConnectDB(ctx, conf.Mongo)
	collection := db.Collection(conf.Mongo.Collection)

	client := &service.BurgerClient{
		Col: *collection,
		Ctx: ctx,
	}

	r := mux.NewRouter()
	s := r.PathPrefix("/api").Subrouter()

	s.HandleFunc("/burgers", controller.CreateBurger(client)).Methods("POST")
	s.HandleFunc("/burgers", controller.FetchBurgers(client)).Methods("GET")
	s.HandleFunc("/burgers/{id}", controller.GetBurger(client)).Methods("GET")
	s.HandleFunc("/burgers/random", controller.GetRandomBurger((client))).Methods("GET")

	http.ListenAndServe(":3000", r)
}
