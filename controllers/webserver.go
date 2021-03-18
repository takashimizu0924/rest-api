package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Work struct {
	
}
type Aircon struct {
	Name string `json:"name"`
	Price int	`json:"price"`
	}
type Antena struct {
	Name string
	Price int
}
var aircon []*Aircon


// 全アイテム取得
func getAllData(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(aircon)
}

// アイテム登録
func createData(w http.ResponseWriter, r *http.Request){
	reqBody,_ := ioutil.ReadAll(r.Body)

	var item Aircon
	if err := json.Unmarshal(reqBody, &item); err != nil {
		log.Fatal(err,"here")
	}
	aircon = append(aircon, &item)
	json.NewEncoder(w).Encode(item)
}

func StartWebServer() error {
	fmt.Println("Rest Api Server Start......")
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/items",getAllData).Methods("GET")
	router.HandleFunc("/item",createData).Methods("POST")
	return http.ListenAndServe(fmt.Sprintf(":%d",8080),router)
}

func init(){
	aircon = []*Aircon{
		&Aircon{
			Name: "標準工事",
			Price: 10000,
		},
		&Aircon{
			Name: "9.0kw工事",
			Price: 15000,
		},
	}
}