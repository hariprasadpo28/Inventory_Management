package Models

type Product struct{
	Id	int 	`json:"id"`
	UniqueId	string	`json:"unique_id"`
	Name	string `json:"name"`
	Price	float32	`json:"price"`
	Quantity	int	`json:"quantity"`
	Description	string	`json:"description"`

}

type User struct {
	Id	int	`json:"id"`
	Name	string	`json:"name" binding:"required"`
	UserName	string	`json:"user_name" binding:"required"`
	Email	string	`json:"email" binding:"required"`
	Orders	[]Order `json:"orders"`
}

type Order struct {
	Id	int `json:"id"`
	ProductId	string `json:"product_id"`
	Quantity	int	`json:"quantity"`
	TotalAmount	float32 `json:"total_amount"`
	UserId User `json:"user_id"`
}

