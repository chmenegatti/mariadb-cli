package utils

import (
	"reflect"
	"time"
)

func ParseError(erro *string) string {

	var err string

	if erro != nil {
		if len(*erro) > 15 {
			err = *erro
			err = err[:15] + "..."
		}
		return err
	} else {
		return ""
	}
}

func ParseDate(date string) string {
	h, _ := time.Parse("2006-01-02 15:04:05 -0700 MST", date)
	return h.Format("02/01/06 15:04")
}

func ParseHeader(data interface{}) []string {
	x := reflect.TypeOf(data)

	var headers []string

	for i := 0; i < x.NumField(); i++ {
		field := x.Field(i)
		if field.Name == "LoadBalanceSize" {
			field.Name = "LB Size"
		}
		if field.Name == "PhysicalFirewall" {
			field.Name = "Phys FW"
		}
		headers = append(headers, field.Name)
	}

	return headers
}
