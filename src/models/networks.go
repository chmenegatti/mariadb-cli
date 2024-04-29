package models

import "time"

type Networks struct {
	Id                      string    `json:",omitempty"`
	Name                    string    `json:",omitempty"`
	Description             string    `json:",omitempty"`
	Address                 string    `json:",omitempty"`
	EnableSideCommunication bool      `json:",omitempty"`
	Status                  string    `json:",omitempty"`
	Error                   *string   `json:",omitempty"`
	Created                 time.Time `json:",omitempty"`
}
