package main

import (
	"github.com/gorilla/mux"
	"github.com/widgetapps/blockchain-golang-example/chain"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", chain.RootEndPoint).Methods("GET")
	r.HandleFunc("/transactions/new", chain.NewTransactionEndPoint).Methods("POST")
	r.HandleFunc("/chain", chain.GetChainEndPoint).Methods("GET")
	r.HandleFunc("/mine", chain.MineEndPoint).Methods("GET")
	r.HandleFunc("/nodes/register", chain.RegisterNodeEndPoint).Methods("POST")
	r.HandleFunc("/consensus", chain.ConsensusEndPoint).Methods("GET")

	if err := http.ListenAndServe(":3000", r); err != nil {
		log.Fatal(err)
	}
}
