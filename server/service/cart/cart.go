package cart

import "github.com/zqhhh/gomall/model"

type service struct {
}

type Service interface {
	Add(cart model.Cart) error
	GetMyCart(userID uint) ([]model.Cart, error)
	Delete(cartID, userID uint) error
	UpdateChecked(cartID, userID uint, checked bool) error
	UpdateNumber(cartID, userID, number uint) error
	Clear(userID uint) error
}

func NewCartService() Service {
	return service{}
}

func (s service) Add(cart model.Cart) error {
	return cart.Insert()
}

func (s service) GetMyCart(userID uint) ([]model.Cart, error) {
	carts, err := model.GetUserCart([]string{"*"}, userID)
	return carts, err
}

func (s service) Delete(cartID, userID uint) error {
	cart := model.Cart{}
	cart.ID, cart.UserID = cartID, userID
	return cart.Delete()
}

func (s service) UpdateNumber(cartID, userID, number uint) error {
	cart := model.Cart{}
	cart.ID, cart.UserID = cartID, userID
	return cart.UpdateFields(map[string]interface{}{
		"number": number,
	})
}

func (s service) UpdateChecked(cartID, userID uint, checked bool) error {
	cart := model.Cart{}
	cart.ID, cart.UserID = cartID, userID
	return cart.UpdateFields(map[string]interface{}{
		"checked": checked,
	})
}

func (s service) Clear(userID uint) error {
	cart := model.Cart{}
	cart.UserID = userID
	return cart.Clear()
}
