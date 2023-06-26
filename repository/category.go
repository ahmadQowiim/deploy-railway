package repository

import (
	"a21hc3NpZ25tZW50/model"
	//"fmt"

	"gorm.io/gorm"
)

type CategoryRepository interface {
	Store(Category *model.Category) error
	Update(id int, category model.Category) error
	Delete(id int) error
	GetByID(id int) (*model.Category, error)
	GetList() ([]model.Category, error)
}

type categoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepo(db *gorm.DB) *categoryRepository {
	return &categoryRepository{db}
}

func (c *categoryRepository) Store(Category *model.Category) error {
	err := c.db.Create(Category).Error
	if err != nil {
		return err
	}

	return nil
}

func (c *categoryRepository) Update(id int, category model.Category) error {
	//return nil // TODO: replace this
	err := c.db.
		Model(model.Category{}).
		Where("id=?", id).
		UpdateColumns(map[string]interface{}{
		"name": category.Name,
	}).Error

	if err != nil {
		return err
	}
	return nil
}

func (c *categoryRepository) Delete(id int) error {
	//return fmt.Errorf("not implement") // TODO: replace this
	var Category model.Category
	err := c.db.Delete(&Category, id).Error
	if err != nil {
		return err
	}
	return nil
}

func (c *categoryRepository) GetByID(id int) (*model.Category, error) {
	var Category model.Category
	err := c.db.Where("id = ?", id).First(&Category).Error
	if err != nil {
		return nil, err
	}

	return &Category, nil
}

func (c *categoryRepository) GetList() ([]model.Category, error) {
	//return nil, nil // TODO: replace this
	result := make([]model.Category, 0)
	rows, err := c.db.Table("categories").Select("*").Rows()
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		err := c.db.ScanRows(rows, &result)
		if err != nil {
			return nil, err
		}
	}
	return result, nil
}
