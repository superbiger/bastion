package controller

import (
	"bastion/internal/response"
	"bastion/pkg"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"io/ioutil"
)

type System struct {
}

var startTime = pkg.NowTime()

// @Summary 系统信息
// @Produce json
// @Success 200 {object} response.Response
// @Router /go/sys/info [get]
func (s *System) Info(c *gin.Context) {
	var msg struct {
		Start   string `json:"start"`
		StartZh string `json:"startZh"`
		Env     string `json:"env"`
	}

	msg.Start = startTime

	str, _ := pkg.PastFromNow(startTime)
	msg.StartZh = str

	msg.Env = viper.GetString("base.env")

	response.Success(c, msg)
	return
}

func (s *System) Error(c *gin.Context) {
	_, err := ioutil.ReadFile("filepath")
	panic(err)
	return
}
