package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	Name  string  `gorm:"type:varchar(100)"`
	Price float64 `gorm:"type:decimal(10,2)"`
	gorm.Model
}

func main() {
	databaseConnectionString := "aula22:aula22@tcp(localhost:3306)/aula22?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(databaseConnectionString), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&Product{})

	// Truncate table
	db.Exec("TRUNCATE TABLE products")

	// Create one
	product := Product{Name: "Product 0", Price: 5.00}
	fmt.Println("Criando o produto com ID:", product.ID)
	db.Create(&product)

	// Create many

	products := []Product{
		{Name: "Product 1", Price: 10.00},
		{Name: "Product 2", Price: 20.00},
		{Name: "Product 3", Price: 30.00},
		{Name: "Product 4", Price: 40.00},
		{Name: "Product 5", Price: 50.00},
		{Name: "Product 6", Price: 60.00},
		{Name: "Product 7", Price: 70.00},
		{Name: "Product 8", Price: 80.00},
		{Name: "Product 9", Price: 90.00},
	}

	fmt.Println("Criando uma lista de produtos")
	db.Create(&products)

	// Read one
	var productRead Product
	db.First(&productRead, "id = ?", product.ID)
	fmt.Printf("Lendo o produto com ID: %v \n", product.ID)
	productRead.information()

	// Read One with many conditions
	var productRead2 Product
	db.First(&productRead2, "name = ? AND price = ?", "Product 1", 10.00)
	productRead2.information()

	// Read many
	var productsRead []Product
	db.Find(&productsRead)
	fmt.Println("Lendo todos os produtos cadastrados")
	for _, product := range productsRead {
		product.information()
	}

	// Update
	product.Name = "Product 0 Updated"
	product.Price = 10.00
	db.Save(&product)
	fmt.Println("Atualizando o produto com ID:", product.ID)

	// Read many
	var updatedProducts []Product
	db.Find(&updatedProducts)
	fmt.Println("Lendo todos os produtos após a atualização")
	for _, product := range updatedProducts {
		product.information()
	}

	// Soft Delete
	db.Delete(&product)
	fmt.Println("Deletando o produto com ID:", product.ID)

	// Read 3 Elements in the second page
	var remainingProducts []Product
	db.Limit(3).Offset(2).Find(&remainingProducts)
	fmt.Println("Lendo os 3 primeiros produtos da segunda página")
	for _, product := range remainingProducts {
		product.information()
	}

	// Get Products with price equal or greater than 50

	var expensiveProducts []Product
	db.Where("price >= ?", 50).Find(&expensiveProducts)
	fmt.Println("Lendo os produtos com preço maior ou igual a 50")
	for _, product := range expensiveProducts {
		product.information()
	}
}

func (product Product) information() {
	created_at := product.CreatedAt.Format("02/01/2006")
	updated_at := product.UpdatedAt.Format("02/01/2006")

	var deletedAt string
	if product.DeletedAt.Valid {
		deletedAt = product.DeletedAt.Time.Format("02/01/2006") // Valid é um ponteiro para um struct, por isso é necessário acessar o campo Time
	} else {
		deletedAt = "null"
	}

	fmt.Printf("ID: %v | Name: %s | Price: %.2f | created_at: %s | updated_at: %s | deleted_at: %s \n", product.ID, product.Name, product.Price, created_at, updated_at, deletedAt)
	fmt.Println("--------------------------------")
}
