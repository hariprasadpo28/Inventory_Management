package Models

import (
	"Inventory_Management/Config"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
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
	var prevOrder Order

	Config.DB.Where("user_name = ?", order.UserName).Last(&prevOrder)

	if prevOrder.Id != 0{
		currentTime := time.Now().Unix()
		if (currentTime - prevOrder.OrderTime) < 300 {
			err := fmt.Errorf("please try again after %d minutes", 5 - (currentTime - prevOrder.OrderTime)/60)
			//fmt.Println(err)
			return err
		}
	}

	if err := Config.DB.Create(order).Error; err != nil {
		return err
	}
	//fmt.Println(order)
	if err = Config.DB.Where("unique_id = ?", order.ProductId).First(&prod).Error; err != nil {
		return err
	}
	if prod.Quantity < order.Quantity {
		Config.DB.Model(order).Updates(Order{Status: "Failed (Out of stock)", TotalAmount: prod.Price * float32(order.Quantity), OrderTime: time.Now().Unix(), RetailerID: prod.RetailerID})
		return nil
	}
	Config.DB.Model(prod).Update("Quantity", prod.Quantity-order.Quantity)
	Config.DB.Model(order).Update(Order{Status: "Placed", TotalAmount: prod.Price * float32(order.Quantity), OrderTime: time.Now().Unix(),RetailerID: prod.RetailerID})

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

func GetUserOrders(orders *[]Order, username string) (err error) {
	if err := Config.DB.Find(&orders, "user_name = ?", username).Error; err != nil {
		return err
	}
	return nil
}

func GetAllOrders(order *[]Order) (err error) {
	if err := Config.DB.Find(&order).Error; err != nil {
		return err
	}
	return nil
}
