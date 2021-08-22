package main

import (
	"net/http"
	"time"
	"log"
	"encoding/json"
)

func main(){
	handler:= http.NewServexMux()

	handler.HandleFunc("/hello", helloHandler)



	s := http.Server{
		Addr: ":8080",
		Handler: handler,
		ReadTimeout: 10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout: 10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Fatal(s.ListenAndServe())


}

type Resp struct{
	Message string
	Error string
}
func helloHandler(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "applictation/json")

	resp:= Resp{
		Message: "hello",
	}

	respJson, _:= json.Marshal(resp)

	w.WriteHeader(http.StatusOK)

	w.Write(respJson)
}
