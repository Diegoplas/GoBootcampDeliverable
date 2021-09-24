package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
)

type PokemonData struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Type1     string `json:"type_1"`
	Type2     string `json:"type_2"`
	Legendary bool   `json:"legendary"`
}

func main() {
	router := getRouter()
	methods := handlers.AllowedMethods([]string{"GET"})
	log.Fatal(http.ListenAndServe(":8000", handlers.CORS(methods)(router)))
}

func getRouter() (router *mux.Router) {
	router = mux.NewRouter()
	router.HandleFunc("/pokedex/{pokemon_id}", getPokemon).Methods("GET")
	return router
}

func getPokemon(w http.ResponseWriter, r *http.Request) {

	requestedIndex := mux.Vars(r)["pokemon_id"]

	wantedIndex, _ := strconv.Atoi(requestedIndex)

	if wantedIndex < 1 || wantedIndex > 151 {
		http.Error(w, "Please introduce a valid pokemon ID from first gen. (1-151)", http.StatusBadRequest)
		return
	}

	response := csvReader(wantedIndex)

	render.New().JSON(w, http.StatusOK, &response)
}

func csvReader(index int) PokemonData {

	// open the file
	csvFile, err := os.Open("./pokedex.csv")
	if err != nil {
		fmt.Printf("error encountered opening csv file: %v", err.Error())
	}

	defer csvFile.Close()

	csvLines, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		fmt.Printf("error encountered opening csv file: %v", err.Error())
	}

	for _, line := range csvLines {

		pokemonID, _ := strconv.Atoi(line[0])

		if pokemonID == index {

			if line[3] == "" {
				line[3] = " - "
			}

			legendary := false
			if line[4] == "TRUE" {
				legendary = true
			}

			pokemon := PokemonData{
				pokemonID,
				line[1],
				line[2],
				line[3],
				legendary,
			}
			return pokemon
		}
	}
	return (PokemonData{})
}
