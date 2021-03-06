package category

import "ecommerce/entities"

type CategoryInterface interface {
	GetAll() ([]entities.Category, error)
	Get(categoryId int) (entities.Category, error)
	Create(newCategory entities.Category) (entities.Category, error)
	Update(newCategory entities.Category, categoryId int) (entities.Category, error)
	Delete(categoryId int) (entities.Category, error)
}
