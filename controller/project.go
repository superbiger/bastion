package controller

import (
	"bastion/common/errno"
	"bastion/controller/validate"
	"bastion/dao"
	"bastion/internal/response"
	"github.com/gin-gonic/gin"
)

// @Summary 创建项目
// @Produce json
// @Router /api/stat/project [post]
func (s *Monitor) CreateProject(c *gin.Context) {
}

// @Summary 查找所有项目
// @Produce json
// @Router /api/stat/projects [get]
func (s *Monitor) FindAllProjects(c *gin.Context) {
	p, err := validate.CheckPagination(c)
	if err != nil {
		response.Fail(c, errno.InvalidParams, err)
		return
	}

	rows, total, e := dao.FindAllProjects(p.PageSize, p.Page, p.Order)
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

// @Summary 查找项目
// @Produce json
// @Router /api/stat/project/:id [get]
func (s *Monitor) FindProjectById(c *gin.Context) {
}
