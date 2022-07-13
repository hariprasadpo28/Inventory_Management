package Models

import (
	"Inventory_Management/Config"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

//AuthUser - Authenticates the user/customer
func AuthUser(username string, password string) error{
	var user User
	if err:= Config.DB.Where("user_name = ?", username).Find(&user).Error; err != nil{
		return fmt.Errorf("user not found")
	}
	if password == user.Password{
		return nil
	} else {
		return fmt.Errorf("password is incorrect")
	}
}

//AuthRetailer - Authenticates the retailer
func AuthRetailer(username string, password string) error{
	var retailer Retailer
	if err:= Config.DB.Where("user_name = ?", username).Find(&retailer).Error; err != nil{
		return fmt.Errorf("user not found")
	}
	if password == retailer.Password{
		return nil
	} else {
		return fmt.Errorf("password is incorrect")
	}
}

//AuthProductRetailer - Checks if the product the retailer trying to update is added by him or not
func AuthProductRetailer(product *Product, uid string, username string) error {
	var retailer Retailer
	Config.DB.Where("user_name = ?", username).Find(&retailer)
	if err := Config.DB.Where("unique_id = ?", uid).Find(product).Error; err != nil{
		return fmt.Errorf("product not found")
	}
	if retailer.RetailerID != product.RetailerID {
		return fmt.Errorf("you cannot update this product")
	}
	return nil
}

//GetProducts - get all available products from the database
func GetProducts(products *[]Product) (err error) {
	if err := Config.DB.Where("quantity > ?", "0").Find(products).Error; err != nil {
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

func PlaceOrder(order *Order, username string) (err error) {
	//var user User
	var prod Product
	var prevOrder Order

	Config.DB.Where("user_name = ?", username).Last(&prevOrder)

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
		Config.DB.Model(order).Updates(Order{Status: "Failed", TotalAmount: prod.Price * float32(order.Quantity), OrderTime: time.Now().Unix(), RetailerID: prod.RetailerID, UserName: username})
		return fmt.Errorf("out of stock")
	}
	Config.DB.Model(prod).Update("Quantity", prod.Quantity-order.Quantity)
	Config.DB.Model(order).Update(Order{Status: "Placed", TotalAmount: prod.Price * float32(order.Quantity), OrderTime: time.Now().Unix(),RetailerID: prod.RetailerID, UserName: username})

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

func CreateRetailer(user *Retailer) (err error) {
	if err := Config.DB.Create(user).Error; err != nil {
		return err
	}
	return nil
}

func GetRetailerOrders(order *[]Order, id string) (err error) {

	if err := Config.DB.Where( "retailer_id = ?", id).Find(&order).Error; err != nil {
		return err
	}
	return nil
}

func GetRetailerID(username string) string{
	var retailer Retailer
	Config.DB.Where("user_name = ?", username).Find(&retailer)
	return retailer.RetailerID
}
