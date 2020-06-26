package service

import (
	"bastion/internal/datasource"
	"bastion/models"
)

func CreateAdminUser(info models.StatAdmin) error {
	err := datasource.GormPool.Create(&info).Error
	return err
}

func FindAllAdminUsers(pagesize, page int, order string) (rows []*models.StatAdmin, total int, e error) {
	offset := (page - 1) * pagesize

	var users []*models.StatAdmin
	var count int

	if order == "" {
		order = "id desc"
	}

	var err error
	err = datasource.GormPool.Order(order).Count(&count).
		Offset(offset).Limit(pagesize).Find(&users).Error

	if err != nil {
		return nil, 0, err
	}

	return users, count, nil
}

func FindAdminUserByUserName(username string) (*models.StatAdmin, error) {
	user := &models.StatAdmin{}
	err := datasource.GormPool.Where("username = ?", username).First(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func FindAdminUserById(id int) (*models.StatAdmin, error) {
	user := &models.StatAdmin{}
	err := datasource.GormPool.Where("id = ?", id).First(&user).Error

	if err != nil {
		return nil, err
	}

	return user, nil
}

func UpdateAdminUser(info models.StatAdmin) error {
	err := datasource.GormPool.Model(&models.StatAdmin{}).Updates(&info).Error
	return err
}
