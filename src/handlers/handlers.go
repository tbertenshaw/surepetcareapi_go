package handlers

import (
	"fmt"
	"net/http"
	"log"
)

func InitHandlers() {
	http.HandleFunc("/", handlerRoot)
	http.HandleFunc("/door", handlerDoor)
	http.HandleFunc("/pet", handlerPet)
	http.HandleFunc("/petlocation", handlerSetPetLocation)
}

func handlerRoot(w http.ResponseWriter, r *http.Request) {
	logIncomingCall(r)
	fmt.Fprint(w, "No functions here!")
}

func handlerDoor(w http.ResponseWriter, r *http.Request) {
	logIncomingCall(r)
	var response = getDoorResponse()

	fmt.Fprintf(w, "Penguin portal Status is : %s", response)
}

func handlerPet(w http.ResponseWriter, r *http.Request) {
	logIncomingCall(r)
	var response = getPetResponse()
	log.Printf("call to pet: %s", response)
	fmt.Fprintf(w, " %s", response)
}

func handlerSetPetLocation(w http.ResponseWriter, r *http.Request) {
	logIncomingCall(r)
	var response = postPetLocationResponse()
	log.Printf("call to pet: %s", response)
	fmt.Fprintf(w, " %s", response)
}

func logIncomingCall( r *http.Request){
	method := r.Method
	log.Printf("%s from %s, ", method, r.RemoteAddr)
}
