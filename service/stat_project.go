package service

import (
	"bastion/internal/datasource"
	"bastion/models"
)

func FindAllProjects(pagesize, page int, order string) (rows []models.StatProject, total int, e error) {
	offset := (page - 1) * pagesize

	var data []models.StatProject
	var count int

	if order == "" {
		order = "id desc"
	}

	var err error
	err = datasource.GormPool.Model(&models.StatProject{}).Order(order).Count(&count).
		Offset(offset).Limit(pagesize).Find(&data).Error

	if err != nil {
		return nil, 0, err
	}

	return data, count, nil
}

func CreateProjects(info models.StatProject) error {
	err := datasource.GormPool.Create(&info).Error
	return err
}

func UpdateProjects(info models.StatProject) error {
	err := datasource.GormPool.Model(&models.StatProject{}).Updates(&info).Error
	return err
}
