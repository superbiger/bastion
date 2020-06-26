package service

import (
	"bastion/controller/validate"
	"bastion/internal/setup"
	"bastion/pkg"
	"testing"
)

func init() {
	setup.InitTest()
}

func TestCreateErrors(t *testing.T) {

}

func TestFindErrorsByAppId(t *testing.T) {
	query := validate.ErrorsQuery{}
	query.AppId = "1"

	data, total, err := FindErrors(query, 10, 1, "")

	pkg.Must(err)
	pkg.PrintJsonString(total)
	pkg.PrintJsonString(data)
}
