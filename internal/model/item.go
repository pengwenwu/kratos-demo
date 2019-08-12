package model

import "time"

// go_server items table
type Items struct {
	ID        int       `json:"id"`
	Appkey    string    `json:"appkey"`
	Channel   int       `json:"channel"`
	Name      string    `json:"name"`
	Photo     string    `json:"photo"`
	Datail    string    `json:"detail"`
	LastDated time.Time `json:"last_dated"`
	Dated     time.Time `json:"dated"`
}