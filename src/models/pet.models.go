package main

import "time"

type Gender int8

const (
	Female Gender = 0
	Male   Gender = 1
)

type Cat int

const (
	Tux Cat = 206802
	Tia Cat = 206803
)

type Pet struct {
	Id           Cat    `json:"id"`
	Name         string `json:"name"`
	Gender       Gender `json:"gender"`
	Comments     string `json:"comments"`
	Household_id int    `json:"household_id"`
	Photo_id     int    `json:"photo_id"`
	Species_id   int    `json:"species_id"`
	Tag_id       int    `json:"tag_id"`
	Version      string `json:"version"`
	Created_at   string `json:"created_at"`
	Updated_at   string `json:"updated_at"`
	Status       Status `json:"status"`
}

type Status struct {
	Activity Activity `json:"activity"`
	Feeding  Feeding  `json:"feeding"`
	Drinking Drinking `json:"drinking"`
}

type Activity struct {
	Tag_id    int    `json:"tag_id"`
	Device_id int    `json:"device_id"`
	Where     Where  `json:"where"`
	Since     string `json:"since"`
}

type Where int8

const (
	Inside  Where = 1
	Outside Where = 2
)

type Feeding struct {
	Tag_id    int       `json:"tag_id"`
	Device_id int       `json:"device_id"`
	Change    []float32 `json:"change"`
	At        string    `json:"at"`
}

type Drinking struct {
	Tag_id    int       `json:"tag_id"`
	Device_id int       `json:"device_id"`
	Change    []float32 `json:"change"`
	At        string    `json:"at"`
}

type PetResponse struct {
	Data []Pet `json:"data"`
}

type PetWhere struct {
	Where Where
	Since time.Time
}
