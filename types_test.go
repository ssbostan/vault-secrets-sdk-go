package vaultsecrets

import (
	"reflect"
	"testing"
)

func TestClient(t *testing.T) {
	client := Client{}
	reflectValue := reflect.ValueOf(client)
	for _, v := range []string{
		"OrganizationID",
		"ProjectID",
		"ApplicationName",
		"ClientID",
		"ClientSecret",
		"AccessToken",
	} {
		if !reflectValue.FieldByName(v).IsValid() {
			t.Errorf("type Client struct: field '%s' not found.", v)
		}
	}
}
