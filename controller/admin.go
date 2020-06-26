package controller

import (
	"bastion/common/constant"
	"bastion/common/errno"
	"bastion/controller/validate"
	"bastion/internal/response"
	"bastion/models"
	"bastion/service"
	"errors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// @Summary 管理员登录
// @Produce json
// @Router /api/stat/admin/login [post]
func (s *Monitor) AdminLogin(c *gin.Context) {
	session := sessions.Default(c)

	p := validate.AdminAccount{}
	err := p.BindingValidParams(c)
	if err != nil {
		response.Fail(c, errno.ErrorNotFound, err)
		return
	}

	// 查找
	user, err := service.FindAdminUserByUserName(p.Username)
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			response.Fail(c, errno.ErrorUserNotFound, err)
			return
		}
		response.Fail(c, errno.ErrorNotFound, err)
		return
	}

	// 校验
	if p.Password != user.Password {
		response.Fail(c, errno.ErrorUsePassword, errors.New("密码错误"))
		return
	}

	session.Set(constant.SessionKeyStatAdmin, user)
	err = session.Save()
	if err != nil {
		response.Fail(c, errno.ErrorSession, err)
		return
	}

	response.Success(c, user, "登录成功")
	return
}

// @Summary 创建管理员
// @Produce json
// @Router /api/stat/admin/create [post]
func (s *Monitor) AdminCreate(c *gin.Context) {
	p := validate.AdminAccountRegister{}
	err := p.BindingValidParams(c)
	if err != nil {
		response.Fail(c, errno.InvalidParams, err)
		return
	}

	err = service.CreateAdminUser(models.StatAdmin{
		Username: p.Username,
		Password: p.Password,
		Email:    p.Email,
	})
	if err != nil {
		response.Fail(c, errno.ErrorCreateData, err)
		return
	}

	response.Success(c, nil, "创建成功")
}

// @Summary 管理员信息
// @Produce json
// @Router /api/stat/admin/info [get]
func (s *Monitor) AdminInfo(c *gin.Context) {
	session := sessions.Default(c)
	get := session.Get(constant.SessionKeyStatAdmin)
	response.Success(c, get)
	return

}

// @Summary 所有管理员
// @Produce json
// @Router /api/stat/admin/list [get]
func (s *Monitor) AdminList(c *gin.Context) {
	p, err := validate.CheckPagination(c)
	if err != nil {
		response.Fail(c, errno.InvalidParams, err)
		return
	}

	rows, total, e := service.FindAllAdminUsers(p.PageSize, p.Page, p.Order)
	if e != nil {
		response.Fail(c, errno.ErrorQueryData, e)
		return
	}

	response.Success(c, response.PageData{
		Page:     p.Page,
		PageSize: p.PageSize,
		Total:    total,
		Rows:     rows,
	})
	return
}

// @Summary 管理员信息更新
// @Produce json
// @Router /api/stat/admin/update [post]
func (s *Monitor) AdminUpdate(c *gin.Context) {
}
