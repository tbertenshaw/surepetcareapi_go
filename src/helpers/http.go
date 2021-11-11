package helpers

import (
	"fmt"

	"net/http"

	"time"
)

func GetResponse(url string,bearerToken string ) (reponse *http.Response, error error) {
	c := http.Client{Timeout: time.Duration(1) * time.Second}
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Set("Authorization", bearerToken)
	if err != nil {
		fmt.Printf("Error %s", err)
		return
	}

	return c.Do(req)
}
