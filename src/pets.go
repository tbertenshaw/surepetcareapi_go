package main

import (
	"encoding/json"
	"log"
	"io/ioutil"
	"net/http"
	"time"
)



func (a *PetResponse) GetPetResponse(bearer string) (response PetResponse) {
		c := http.Client{Timeout: time.Duration(1) * time.Second}
	req, err := http.NewRequest("GET", urlPetStatusRoot, nil)
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

	if err := json.Unmarshal(body, &response); err != nil {
		log.Println("failed to unmarshall:", err)
	} else {
		log.Println("deserialised")
		
	}
	return
}
