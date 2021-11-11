package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"
	"bytes"
	"tbertenshaw/surepetapi_go/helpers"
)

func (a *PetResponse) GetPetResponse(bearer string) (response PetResponse) {

	resp, err := helpers.GetResponse(urlPetStatusRoot,bearer )

	if err != nil {
		log.Printf("Error %s", err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error %s", err)
		return
	}

	if err := json.Unmarshal(body, &response); err != nil {
		log.Println("failed to unmarshall:", err)
	} else {
		log.Println("deserialised")

	}
	return
}

func (a *PetResponse) PostPetLocation(bearer string,cat Cat,petLocation PetLocation) {
	
	c := http.Client{Timeout: time.Duration(1) * time.Second}
	payloadBuf := new(bytes.Buffer)
	json.NewEncoder(payloadBuf).Encode(petLocation)
	req, err := http.NewRequest("POST", urlPetLocationRoot, payloadBuf)
	req.Header.Set("Authorization", bearer)
	if err != nil {
		log.Printf("Error %s", err)
		return
	}

	resp, err := c.Do(req)

	if err != nil {
		log.Printf("Error %s", err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error %s", err)
		return
	}

response :=""

	if err := json.Unmarshal(body, &response); err != nil {
		log.Println("failed to unmarshall:", err)
	} else {
		log.Println("deserialised")

	}
	//return
}

type PetLocation struct {
	Where Where `json:"where"`
	Since string `json:"since"`
}
