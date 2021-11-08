package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)
var bearer string
var urlDoorRoot string
var urlPetStatusRoot string

func main() {
	bearer = AppConfig.BearerToken
	urlDoorRoot = AppConfig.UrlDoorRoot
	urlPetStatusRoot = AppConfig.UrlPetStatusRoot
	http.HandleFunc("/", handlerRoot)
	http.HandleFunc("/door", handlerDoor)
	http.HandleFunc("/pet", handlerPet)
	log.Fatal(http.ListenAndServe("0.0.0.0:9090", nil))
}

func handlerRoot(w http.ResponseWriter, r *http.Request) {
	method := r.Method
	log.Printf("%s from %s, ",method, r.RemoteAddr)
	fmt.Fprint(w, "No functions here!")
}

func handlerDoor(w http.ResponseWriter, r *http.Request) {
	method := r.Method
	log.Printf("%s from %s, ",method, r.RemoteAddr)
	var response = getDoorResponse()

	fmt.Fprintf(w, "Penguin portal Status is : %s", response)
}

func handlerPet(w http.ResponseWriter, r *http.Request) {
	method := r.Method
	log.Printf("%s from %s, ",method, r.RemoteAddr)
	var response = getPetResponse()
	log.Printf("call to pet: %s", response)
	fmt.Fprintf(w, " %s", response)
}

func getDoorResponse() (status string) {

	doorResponse := DoorResponse{}
	var result = doorResponse.GetDoorStatus(bearer)

	switch DoorState(result) {
	case Open:
		status = "Door is Open"
	case LockedIn:
		status = "Door is Inwards Only"
	case LockedOut:
		status = "Door is Outwards Only"
	case LockedBoth:
		status = "Door is Locked"
	}
	//status = getCurfewStatus(status)
	return
}

func getPetResponse() (status string) {

	petResponse := PetResponse{}
	var result = petResponse.GetPetResponse(bearer)
	var builder strings.Builder
	for _, pet := range result.Data {

		switch Where(pet.Status.Activity.Where) {
		case Outside:
			log.Printf("\n%s is Outside", pet.Name)
			builder.WriteString(fmt.Sprintf("\n%s is Outside", pet.Name))

		case Inside:
			log.Printf("\n%s is Inside", pet.Name)
			builder.WriteString(fmt.Sprintf("\n%s is Inside", pet.Name))

		}
	}
	builder.WriteString("\n")
	status = builder.String()
	return
}

// func getCurfewStatus(status string) (statusout string) {

// }
