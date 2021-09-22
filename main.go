package main

// Dudas----------------------------------------
//   Go mod e init
//   Acomodo de carpetas
//   Directorio de 0
// 	 regresar como mapa o string con formato
//   Como agregar name Y ID
//   Usar ID como string o como int???

import (
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

type PokemonData struct {
	ID        int
	Name      string
	Type1     string
	Type2     string
	Legendary bool
}

func main() {
	router := getRouter()
	methods := handlers.AllowedMethods([]string{"GET"})
	log.Fatal(http.ListenAndServe(":8000", handlers.CORS(methods)(router)))
}

func getRouter() (router *mux.Router) {
	router = mux.NewRouter()
	router.HandleFunc("/pokedex/{wanted_pokemon}", getPokemon).Methods("GET")
	return router
}

func getPokemon(w http.ResponseWriter, r *http.Request) {
	wantedIndex := ""
	formatedResponse := "No pokemon found!"
	pokeData := PokemonData{}
	wantedIndex = mux.Vars(r)["wanted_pokemon"]
	fmt.Println(wantedIndex)
	if wantedIndex == "" {
		http.Error(w, "missing parameter pokemon Name or ID", http.StatusBadRequest)
	}

	pokeData = csvReader(wantedIndex)

	if pokeData.Type2 != "" {
		formatedResponse = fmt.Sprintf("ID: %v Name: %s, Type1: %s, Type2: %s, Legendary: %t",
			pokeData.ID, pokeData.Name, pokeData.Type1, pokeData.Type2, pokeData.Legendary)
	} else {
		formatedResponse = fmt.Sprintf("ID: %v Name: %s, Type1: %s, Legendary: %t",
			pokeData.ID, pokeData.Name, pokeData.Type1, pokeData.Legendary)
	}
	w.Write([]byte(formatedResponse))
}

func csvReader(index string) PokemonData {

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

	legendary := false

	for _, line := range csvLines {

		if line[0] == index {
			if line[4] == "TRUE" {
				legendary = true
			}
			intID, err := strconv.Atoi(line[0])
			if err != nil {
				fmt.Printf("error converting id to int: %v", err.Error())
			}
			pokemon := PokemonData{
				intID,
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
