package Models

import "Inventory_Management/Config"

func GetProducts(products *[]Product) (err error) {
	if err := Config.DB.Find(products).Error; err != nil {
		return err
	}
	return nil
}

func AddProduct(product *Product) (err error) {
	if err := Config.DB.Create(product).Error; err != nil {
		return err
	}
	return nil
}

func PlaceOrder(order *Order, id string, prodID string) (err error) {
	var user User
	var prod Product
	if err := Config.DB.Create(order).Error; err != nil {
		return err
	}
	if err = Config.DB.Where("id = ?", id).First(user).Error; err != nil {
		return err
	}
	//return nil

	if err = Config.DB.Where("unique_id = ?", prodID).First(prod).Error; err != nil {
		return err
	}
	//return nil
	user.Orders = append(user.Orders, *order)
	Config.DB.Model(&order).Updates(Order{UserId: user, Status: "Placed", TotalAmount: prod.Price * float32(order.Quantity)})
	Config.DB.Model(&user).Updates(User{Orders: user.Orders})

	return nil

}
