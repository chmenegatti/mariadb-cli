package models

import "time"

type VirtualMachines struct {
	Id              string    `json:",omitempty"`
	Organization    string    `json:",omitempty"`
	Name            string    `json:",omitempty"`
	Description     string    `json:",omitempty"`
	Network         string    `json:",omitempty"`
	DiskSize        int       `json:",omitempty"`
	PrivateIp       string    `json:",omitempty"`
	FloatingIp      string    `json:",omitempty"`
	TypeSO          *string   `json:",omitempty"`
	Topology        string    `json:",omitempty"`
	TypeApplication string    `json:",omitempty"`
	State           string    `json:",omitempty"`
	Status          string    `json:",omitempty"`
	Error           *string   `json:",omitempty"`
	Created         time.Time `json:",omitempty"`
}
