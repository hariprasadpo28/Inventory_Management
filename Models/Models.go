package Models

type Product struct {
	Id          int     `json:"id"`
	UniqueId    string  `json:"unique_id" binding:"required" gorm:"unique"`
	Name        string  `json:"name" binding:"required"`
	Price       float32 `json:"price" binding:"required"`
	Quantity    int     `json:"quantity" binding:"required"`
	Description string  `json:"description" binding:"required"`
	RetailerID  string  `json:"retailer_id" binding:"required"`
}

type Retailer struct {
	Id         int    `json:"id"`
	RetailerID string `json:"retailer_id" gorm:"unique"`
	Name       string `json:"name" binding:"required"`
	UserName   string `json:"user_name" binding:"required" gorm:"unique"`
	Password   string `json:"password" binding:"required"`
}

type User struct {
	Id       int    `json:"id"`
	Name     string `json:"name" binding:"required"`
	UserName string `json:"user_name" binding:"required" gorm:"unique"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
	//Orders   []Order `json:"orders"`
}

type Order struct {
	Id          int     `json:"id"`
	ProductId   string  `json:"product_id" binding:"required"`
	Quantity    int     `json:"quantity" binding:"required"`
	TotalAmount float32 `json:"total_amount"`
	UserName    string  `json:"user_name"`
	Status      string  `json:"status"`
	OrderTime   int64   `json:"order_time"`
	RetailerID  string  `json:"retailer_id"`
}

func (b *User) TableName() string {
	return "user"
}

func (b *Product) TableName() string {
	return "product"
}

func (b *Order) TableName() string {
	return "order"
}

func (b *Retailer) TableName() string {
	return "retailer"
}
