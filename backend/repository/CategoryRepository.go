package repository

import (
	"gin_demo/common"
	"gin_demo/model"

	"gorm.io/gorm"
)

type CategoryRepository struct {
	DB *gorm.DB
}

func NewCategoryRepository() CategoryRepository {
	return CategoryRepository{DB: common.GetDB()}
}

// ? 创建
func (c CategoryRepository) Create(name string) (*model.Category, error) {
	category := model.Category{
		Name: name,
	}

	err := c.DB.Create(&category).Error

	if err != nil {
		return nil, err
	}

	return &category, nil
}

// ? 更新
func (c CategoryRepository) Update(category model.Category, name string) (*model.Category, error) {

	err := c.DB.Model(&category).Update("name", name).Error

	if err != nil {
		return nil, err
	}

	return &category, nil
}

// ? 查找
func (c CategoryRepository) SelectByID(id int) (*model.Category, error) {

	category := model.Category{}

	err := c.DB.First(&category, id).Error

	if err != nil {
		return nil, err
	}

	return &category, nil
}

// ? 删除
func (c CategoryRepository) DeleteByID(id int) error {

	err := c.DB.Delete(model.Category{}, id).Error

	if err != nil {
		return err
	}

	return nil
}
