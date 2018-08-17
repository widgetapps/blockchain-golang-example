package chain

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Transaction struct {
	Sender    string  `json:"sender"`
	Recipient string  `json:"recipient"`
	Amount    float64 `json:"amount"`
}

type Block struct {
	Index        uint
	Timestamp    time.Time
	Transactions []Transaction
	Proof        uint
	PreviousHash string
}

var transactions []Transaction
var chain []Block

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
	fmt.Printf("Number of transactions: %d\n", len(transactions))

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
