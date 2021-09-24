package routes

import "github.com/gorilla/mux"

func getRouter() (router *mux.Router) {
	router = mux.NewRouter()
	router.HandleFunc("/pokedex/{pokemon_id}", getPokemon).Methods("GET")
	return router
}
