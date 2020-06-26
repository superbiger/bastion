package service

import (
	"bastion/controller/validate"
	"bastion/internal/datasource"
	"bastion/models"
)

func CreateErrors(data models.StatError) error {
	err := datasource.GormPool.Create(&data).Error
	return err
}

func FindErrors(query validate.ErrorsQuery, pagesize, page int, order string) (
	rows []models.StatError, total int, e error) {

	offset := (page - 1) * pagesize

	var data []models.StatError
	var count int

	if order == "" {
		order = "id desc"
	}

	var err error
	db := datasource.GormPool.Model(&models.StatError{})

	if query.AppId != "" {
		db = db.Where("appid = ?", query.AppId)
	}
	if query.Tag != "" {
		db = db.Where("tag = ?", query.Tag)
	}
	if query.Uid != "" {
		db = db.Where("uid = ?", query.Uid)
	}
	if query.ErrorMsg != "" {
		db = db.Where("error_msg = ?", query.ErrorMsg)
	}
	if query.CreatedAt != "" {
		db = db.Where("created_at > ?", query.CreatedAt)
	}

	err = db.Order(order).Count(&count).
		Offset(offset).Limit(pagesize).Find(&data).Error

	if err != nil {
		return nil, 0, err
	}

	return data, count, nil
}

func UpdateErrors(info models.StatError) error {
	err := datasource.GormPool.Model(&models.StatError{}).Updates(&info).Error
	return err
}
