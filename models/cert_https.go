package models

import (
	"fmt"

	"dvij.geoloc/conf"
	"github.com/kabukky/httpscerts"
)

// MakeHTTPSCertV1 If cert files are not available, generate new ones.
func MakeHTTPSCertV1(nameCert string, nameKey string, hostName string) *conf.ApiError {
	err := httpscerts.Check(nameCert, nameKey)
	if err != nil {
		err = httpscerts.Generate(nameCert, nameKey, hostName)
		if err != nil {
			fmt.Print("Error: Couldn't create https certs.")
			return conf.ErrHttpsCert
		}
	}
	return nil
}
