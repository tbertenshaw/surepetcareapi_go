package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"tbertenshaw/surepetapi_go/helpers"
	"time"
)

type DoorResponse struct {
	Data DoorResponseData `json:"data"`
}

type DoorResponseData struct {
	Curfews      []Curfew  `json:"curfew"`
	Locking      DoorState `json:"locking"`
	Fast_polling bool      `json:"fast_polling"`
}

type Curfew struct {
	Enabled     bool   `json:"enabled"`
	Lock_time   string `json:"lock_time"`
	Unlock_time string `json:"unlock_time"`
}

type DoorState = int8

const (
	Open       DoorState = 0
	LockedIn   DoorState = 1
	LockedOut  DoorState = 2
	LockedBoth DoorState = 3
)

func (a *DoorResponse) GetDoorStatus(bearer string) (lockstatus DoorState) {

	resp, err := helpers.GetResponse(urlDoorRoot,bearer )

	if err != nil {
		fmt.Printf("Error %s", err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error %s", err)
		return
	}

	if err := json.Unmarshal(body, &a); err != nil {
		fmt.Println("failed to unmarshall:", err)
	} else {
		fmt.Println(a.Data.Locking)
		lockstatus = a.Data.Locking
		curfewLocked := getCurfewStatus(a.Data.Curfews)
		if curfewLocked {
			if lockstatus == Open {
				lockstatus = LockedIn
			}
			if lockstatus == LockedOut {
				lockstatus = LockedBoth
			}
		}
	}
	return
}

func getCurfewStatus(curfews []Curfew) (locked bool) {
	for _, curfew := range curfews {
		if !curfew.Enabled {
			continue
		}
		hr, min, _ := time.Now().Clock()
		timeNow := (hr *100) + min

		locktime,_ := strconv.Atoi(strings.ReplaceAll(curfew.Lock_time,":",""))

		unlocktime,_ := strconv.Atoi(strings.ReplaceAll(curfew.Unlock_time,":",""))

		if locktime < unlocktime {
			if timeNow > locktime && timeNow <unlocktime {
				//locked
				locked = true
				return
			}
		} else {
			if timeNow >locktime || timeNow < unlocktime {
				//locked
				locked = true
				return
			}
		}
	}
	return
}
