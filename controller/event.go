package controller

import (
	"bastion/common/errno"
	"bastion/controller/validate"
	"bastion/dao"
	"bastion/internal/response"
	"bastion/models"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)


type Monitor struct {
}


// @Summary 上报错误
// @Produce json
// @Router /api/stat/errors [post]
func (s *Monitor) CreateError(c *gin.Context) {
	p := validate.ErrorData{}
	_ = p.BindingValidParams(c)

	data := models.StatError{
		Appid:       p.Appid,
		Path:        p.Path,
		ErrorMsg:    p.Msg,
		ErrorString: p.Desc,
		UId:         p.UId,
	}

	e := dao.CreateErrors(data)
	if e != nil {
		c.String(http.StatusOK, e.Error())
		return
	}
	c.String(http.StatusOK, "ok")
}

// @Summary errlock上报
// @Produce json
// @Router /api/stat/errors [post]
func (s *Monitor) ImgCreateError(c *gin.Context) {
	p := validate.ErrorData{}
	_ = p.BindingValidParams(c)

	// 设备信息
	if p.BrowserUa != "" {
		data := models.StatDevice{
			UId:           p.UId,
			BrowserUa:     p.BrowserUa,
			BrowserResult: p.BrowserResult,
		}
		if e := dao.CreateIfNotExistDevice(data); e != nil {
			c.String(http.StatusOK, e.Error())
			return
		}
	}

	// 创建错误
	if p.Msg != "" {
		data := models.StatError{
			Appid:       p.Appid,
			Tag:         p.Tag,
			Path:        p.Path,
			ErrorMsg:    p.Msg,
			ErrorString: p.Desc,
			FileURL:     p.FileUrl,
			Lineno:      p.Lineno,
			Colno:       p.Colno,
			UId:         p.UId,
		}
		if e := dao.CreateErrors(data); e != nil {
			c.String(http.StatusOK, e.Error())
			return
		}
	}
	c.String(http.StatusOK, "ok")
}

// @Summary 查找所有错误
// @Produce json
// @Router /api/stat/errors [get]
func (s *Monitor) FindErrorsWithParams(c *gin.Context) {
	p, err := validate.CheckPagination(c)
	if err != nil {
		response.Fail(c, errno.InvalidParams, err)
		return
	}

	query := validate.ErrorsQuery{}
	err = query.BindingValidParams(c)
	if err != nil {
		response.Fail(c, errno.InvalidParams, err)
		return
	}

	findErrors, total, e := dao.FindErrors(query, p.PageSize, p.Page, p.Order)
	if e != nil {
		response.Fail(c, errno.ErrorQueryData, e)
		return
	}

	response.Success(c, response.PageData{
		Page:     p.Page,
		PageSize: p.PageSize,
		Total:    total,
		Rows:     findErrors,
	})
	return
}

// @Summary 根据id查找错误
// @Produce json
// @Router /api/stat/error/:id [get]
func (s *Monitor) FindErrorById(c *gin.Context) {
}

// @Summary 根据id查找错误
// @Produce json
// @Router /api/stat/device/:uid [get]
func (s *Monitor) FindDeviceByUid(c *gin.Context) {
	val := c.Query("uid")
	if val != "" {
		fmt.Println("", val)
	}

	res, err := dao.FindByUid(val)
	if err != nil {
		response.Fail(c, errno.ErrorNotFound, err)
		return
	}
	response.Success(c, res)
}

func (s *Monitor) FindAllDevice(c *gin.Context) {
	p, err := validate.CheckPagination(c)
	if err != nil {
		response.Fail(c, errno.InvalidParams, err)
		return
	}

	device, total, e := dao.FindAllDevice(p.PageSize, p.Page, p.Order)
	if e != nil {
		response.Fail(c, errno.ErrorQueryData, e)
		return
	}

	response.Success(c, response.PageData{
		Page:     p.Page,
		PageSize: p.PageSize,
		Total:    total,
		Rows:     device,
	})
}

// @Summary 上报打点
// @Produce json
// @Router /api/stat/behavior [post]
func (s *Monitor) CreateBehavior(c *gin.Context) {
}

// @Summary 查找所有打点
// @Produce json
// @Router /api/stat/behaviors [get]
func (s *Monitor) FindAllBehaviors(c *gin.Context) {
}

// @Summary 根据id查找打点
// @Produce json
// @Router /api/stat/behavior/:id [get]
func (s *Monitor) FindBehaviorById(c *gin.Context) {
}

func (s *Monitor) TestFail(c *gin.Context) {
	response.Fail(c, errno.ErrorQueryData, errors.New("query failed"))
}

func (s *Monitor) TestError(c *gin.Context) {
	response.Error(c, http.StatusInternalServerError, errors.New("server error"))
}

func (s *Monitor) TestTimeOut(c *gin.Context) {
	time.Sleep(time.Second * 11)
	response.Success(c, nil, "ok")
}
