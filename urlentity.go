package main

import "time"

type urlEntity struct {
	ID         int       `json:"ID"`
	URL        string    `json:"URL"`
	CREATED_AT time.Time `json:"CREATED_AT"`
}
