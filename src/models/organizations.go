package models

import "time"

type Organization struct {
	Id               string    `json:",omitempty"`
	Name             string    `json:",omitempty"`
	TierProvider     string    `json:",omitempty"`
	BackupCluster    string    `json:",omitempty"`
	PhysicalFirewall string    `json:",omitempty"`
	VirtualFirewall  string    `json:",omitempty"`
	LoadBalanceSize  string    `json:",omitempty"`
	Status           string    `json:",omitempty"`
	Error            *string   `json:",omitempty"`
	Created          time.Time `json:",omitempty"`
}
