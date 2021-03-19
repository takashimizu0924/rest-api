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
	CompletedDate string `json:"completedDate"`
	RecieptNumber int    `json:"recieptNumber"`
	Name          string `json:"name"`
	Quantity      int    `json:"quantity"`
	Price         int	 `json:"price"`
	}
type Antena struct {
	Name string
	Price int
}
var aircon []*Aircon


// 全アイテム取得
func getAllData(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(aircon)
}

// アイテム登録
func createData(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	reqBody,_ := ioutil.ReadAll(r.Body)

	var item Aircon
	if err := json.Unmarshal(reqBody, &item); err != nil {
		log.Fatal(err,"here")
	}
	aircon = append(aircon, &item)
	json.NewEncoder(w).Encode(item)
}

func handleCORS(handle http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		log.Println(r)
		w.Header().Set("Access-Control-Allow-Headers","Content-Type")
		w.Header().Set("Access-Control-Allow-origin","*")
		w.Header().Set("Access-Control-Allow-Methods","GET, POST,PUT, DELETE, OPTIONS")
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		handle.ServeHTTP(w, r)
		return
	})
}

func StartWebServer() error {
	fmt.Println("Rest Api Server Start......")
	router := mux.NewRouter().StrictSlash(true)
	router.Use(handleCORS)
	router.HandleFunc("/items",getAllData).Methods("GET")
	router.HandleFunc("/item",createData)
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