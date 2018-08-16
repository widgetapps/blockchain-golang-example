package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

type Transaction struct {
	Sender    string
	Recipient string
	amount    float64
}

type Block struct {
	Index        uint64
	Timestamp    time.Time
	Transactions []Transaction
	Proof        uint64
	PreviousHash string
}

var transactions []Transaction
var chain []Block

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", RootEndPoint).Methods("GET")
	r.HandleFunc("/transactions/new", NewTransactionEndPoint).Methods("POST")
	r.HandleFunc("/chain", GetChainEndPoint).Methods("GET")
	r.HandleFunc("/mine", MineEndPoint).Methods("GET")
	r.HandleFunc("/nodes/register", RegisterNodeEndPoint).Methods("POST")
	r.HandleFunc("/consensus", ConsensusEndPoint).Methods("GET")

	if err := http.ListenAndServe(":3000", r); err != nil {
		log.Fatal(err)
	}
}

func RootEndPoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "This is an example blockchain node written in GoLang. Isn't that weird?")
}

func NewTransactionEndPoint(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create a new transaction.")

	transaction := Transaction{}
	err := json.NewDecoder(r.Body).Decode(&transaction)
	if err != nil {
		panic(err)
	}

	transactions = append(transactions, transaction)

	transJson, err := json.Marshal(transaction)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(transJson)
}

func GetChainEndPoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Get the chain.")
}

func MineEndPoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Mine a block.")
}

func RegisterNodeEndPoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Register a new node.")
}

func ConsensusEndPoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Consensus.")
}
