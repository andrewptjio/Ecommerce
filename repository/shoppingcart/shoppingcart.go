package shoppingcart

import (
	"ecommerce/entities"

	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

type ShoppingCartRepository struct {
	db *gorm.DB
}

func NewShoppingCartRepo(db *gorm.DB) *ShoppingCartRepository {
	return &ShoppingCartRepository{db: db}
}
func (ur *ShoppingCartRepository) Get(userId int) ([]entities.ShoppingCart, error) {
	cart := []entities.ShoppingCart{}
	if err := ur.db.Where("user_id = ?", userId).Find(&cart).Error; err != nil {
		log.Warn("Found database error", err)
		return cart, err
	}
	return cart, nil
}
func (ur *ShoppingCartRepository) GetById(id, userId int) (entities.ShoppingCart, error) {
	cart := entities.ShoppingCart{}
	if err := ur.db.Where("id=? and user_id = ?", id, userId).First(&cart).Error; err != nil {
		log.Warn("Found database error", err)
		return cart, err
	}
	return cart, nil
}
func (ur *ShoppingCartRepository) Create(newShoppingcart entities.ShoppingCart) (entities.ShoppingCart, error) {
	product := entities.Product{}

	if err := ur.db.Find(&product, "id= ?", newShoppingcart.ProductID).Error; err != nil {
		return newShoppingcart, err
	}
	newShoppingcart.Subtotal = newShoppingcart.Qty * product.Price

	ur.db.Save(&newShoppingcart)
	return newShoppingcart, nil
}
func (ur *ShoppingCartRepository) Update(updateCart entities.ShoppingCart, cartId, userId int) (entities.ShoppingCart, error) {
	cart := entities.ShoppingCart{}
	if err := ur.db.First(&cart, "id=? and user_id=?", cartId, userId).Error; err != nil {
		return cart, err
	}
	if updateCart.Qty != 0 {
		product := entities.Product{}
		if err := ur.db.First(&product, "id= ?", cart.ProductID).Error; err != nil {
			return updateCart, err
		}
		updateCart.Subtotal = updateCart.Qty * product.Price
	}
	ur.db.Model(&cart).Updates(updateCart)

	return cart, nil
}
func (ur *ShoppingCartRepository) Delete(cartId, userId int) (entities.ShoppingCart, error) {
	cart := entities.ShoppingCart{}
	err := ur.db.First(&cart, "id = ? and user_id=?", cartId, userId).Delete(&cart).Error
	if err != nil {
		return cart, err
	}
	return cart, nil
}
