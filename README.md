# erago

# library
- github.com/gorilla/mux
- github.com/jinzhu/gorm/dialects/mysql

# language
- GO

# How to Run

**go run main.go**

# Port
8080

# structure
**root**
  - controllers
    - productcontrollers.go
  - database
    - client.go
    - config.go
  - entity
    - product.go
   
# Routes
- **localhost:8080/create**
  _Input Product data to database_
- **localhost:8080/get**
  _Get All product data_ 
- **localhost:8080/get/{id}**
  _Get single product data by id product_
- **localhost:8080/getTerbaru**
  _get product by newest product_
- **localhost:8080/getTermurah**
  _get product by cheap price_
- **localhost:8080/getTermahal**
  _get product by most expensive price_
- **localhost:8080/getAllProductAsc**
  _get product data sort by name ascending (A-Z)_
- **localhost:8080/getAllProductDesc**
  _get product data sort by name descending (Z-A)_
