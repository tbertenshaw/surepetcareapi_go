package main

import (
	"fmt"
	//"helpers"
	"log"
	"net/http"
	"strings"
	"tbertenshaw/surepetapi_go/helpers"
	"tbertenshaw/surepetapi_go/handlers"
)

var bearer string
var urlDoorRoot string
var urlPetStatusRoot string
var urlPetLocationRoot string
var appConfig helpers.Cfg

func main() {
	initConfig()
	handlers.InitHandlers()
	log.Fatal(http.ListenAndServe("0.0.0.0:9090", nil))
}

func initConfig() {
	appConfig = helpers.AppConfig.ReadConfig()
	bearer = appConfig.BearerToken
	urlDoorRoot = appConfig.UrlDoorRoot
	urlPetStatusRoot = appConfig.UrlPetStatusRoot
	urlPetLocationRoot = appConfig.UrlPetLocationRoot
}



func postPetLocationResponse() (status string) {
	petResponse := PetResponse{}
	petLocation := PetLocation{Where: Outside, Since: "2021-11-11 16:24:01"}

	petResponse.PostPetLocation(bearer, Tia, petLocation)

	return "nice"
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
