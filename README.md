# Inventory_Management
Inventory management app with shop, update, retrieve orders 

##API Endpoints:
GET   '/api/products'       : Get all products

POST  '/api/product'        : Post new product
      JSON data:
      {
          "unique_id": ,     
          "quantity": 7,      
          "price": 0,
          "description": "",
          "name" : ""
      }
      
POST   '/api/product/:uid'   : update product details

POST   '/api/order/'         : Place new order
        JSON Data:
        {
          "product_id": ,
          "username": ,
          "quantity": ,
        }
      
GET   'api/order/:username'   : get all user orders

GET   'api/orders'            : get all orders

POST  'api/user'              : create new user
      JSON DATA:
      {
        "name": "",
        "user_name": "",        
        "email": "example@gmail.com"
      }
       
GET   'api/user'              : get all users 
       
