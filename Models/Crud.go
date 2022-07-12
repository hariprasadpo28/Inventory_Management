package Models

import (
	"Inventory_Management/Config"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func GetProducts(products *[]Product) (err error) {
	if err := Config.DB.Find(products).Error; err != nil {
		return err
	}
	return nil
}

func GetProductByID(product *Product, uid string) (err error) {
	if err = Config.DB.Where("unique_id = ?", uid).First(product).Error; err != nil {
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

func UpdateProduct(product *Product, id string) (err error) {
	//fmt.Println(user)
	Config.DB.Save(product)
	return nil
}

func PlaceOrder(order *Order) (err error) {
	//var user User
	var prod Product

	if err := Config.DB.Create(order).Error; err != nil {
		return err
	}
	fmt.Println(order)
	Config.DB.Where("unique_id = ?", order.ProductId).First(&prod)
	fmt.Println(prod)
	if err = Config.DB.Where("unique_id = ?", order.ProductId).First(&prod).Error; err != nil {
		return err
	}
	fmt.Println(prod)
	if prod.Quantity < order.Quantity {
		Config.DB.Model(order).Updates(Order{Status: "Failed (Product is not available)", TotalAmount: prod.Price * float32(order.Quantity)})
		return nil
	}
	Config.DB.Model(prod).Update("Quantity", prod.Quantity-order.Quantity)
	Config.DB.Model(order).Update(Order{Status: "Placed", TotalAmount: prod.Price * float32(order.Quantity)})

	return nil

}

func CreateUser(user *User) (err error) {
	if err := Config.DB.Create(user).Error; err != nil {
		return err
	}
	return nil
}

func GetAllUsers(users *[]User) (err error) {
	if err := Config.DB.Find(users).Error; err != nil {
		return err
	}
	return nil
}

func GetUserOrders(orders *[]Order, username string) (err error){
	if err := Config.DB.Find(&orders,"user_name = ?", username).Error; err != nil{
		return err
	}
	return nil
}

func GetAllOrders (order *[]Order) (err error){
	if err := Config.DB.Find(&order).Error; err != nil{
		return err
	}
	return nil
}
