package dao

import (
	"bastion/internal/datasource"
	"bastion/models"
)

func FindAllBehaviors(pagesize, page int, order string) (rows []models.StatBehavior, total int, e error) {
	offset := (page - 1) * pagesize

	var data []models.StatBehavior
	var count int

	if order == "" {
		order = "id desc"
	}

	var err error
	err = datasource.GormPool.Model(&models.StatBehavior{}).Order(order).Count(&count).
		Offset(offset).Limit(pagesize).Find(&data).Error

	if err != nil {
		return nil, 0, err
	}

	return data, count, nil
}

func CreateBehaviors(info models.StatBehavior) error {
	err := datasource.GormPool.Create(&info).Error
	return err
}

func UpdateBehaviors(info models.StatBehavior) error {
	err := datasource.GormPool.Model(&models.StatBehavior{}).Updates(&info).Error
	return err
}

