package dao

import (
	"bastion/controller/validate"
	"bastion/internal/setup"
	"bastion/utils"
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

	utils.Must(err)
	utils.PrintJsonString(total)
	utils.PrintJsonString(data)
}
