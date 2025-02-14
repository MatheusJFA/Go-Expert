package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
)

// _ "github.com/go-sql-driver/mysql" é uma importação anônima, ou seja, não é necessário chamar diretamente

type Product struct {
	ID    string
	Name  string
	Price float64
}

func NewProduct(name string, price float64) *Product {
	return &Product{
		ID:    uuid.New().String(),
		Name:  name,
		Price: price,
	}
}

func main() {
	// The datasource name is the connection string, and it's made up of the username, password, host, port, and database name.
	datasourceName := "aula21:aula21@tcp(localhost:3306)/aula21"
	db, err := sql.Open("mysql", datasourceName)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	truncateProducts(db)

	productList := []*Product{
		NewProduct("Product 1", 10.0),
		NewProduct("Product 2", 20.0),
		NewProduct("Product 3", 30.0),
		NewProduct("Product 4", 40.0),
		NewProduct("Product 5", 50.0),
	}

	for _, product := range productList {
		_, err := insertProduct(db, product)
		if err != nil {
			panic(err)
		}

		fmt.Println("Produto inserido com sucesso: " + product.Name)
	}
	separarLinhas()

	productList[0].Price = 100.0
	_, err = updateProduct(db, productList[0])
	if err != nil {
		panic(err)
	}

	fmt.Println("Produto atualizado com sucesso: " + productList[0].Name)

	separarLinhas()

	product, err := getProductByID(db, productList[0].ID)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Foi encontrado o produto de nome = %v e preço = R$ %.2f \n", product.Name, product.Price)
	separarLinhas()

	products, err := getProducts(db)
	if err != nil {
		panic(err)
	}

	fmt.Println("Produtos encontrados:")
	for _, product := range products {
		fmt.Printf("id: %s, nome: %v, preço: R$ %.2f \n", product.ID, product.Name, product.Price)
	}
	separarLinhas()

	_, err = deleteProductByID(db, productList[3].ID)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Produto deletado com sucesso! \n")

	separarLinhas()

	products, err = getProducts(db)
	if err != nil {
		panic(err)
	}

	fmt.Println("Nova lista de produtos encontrados:")
	for _, product := range products {
		fmt.Printf("id: %s, nome: %v, preço: R$ %.2f \n", product.ID, product.Name, product.Price)
	}

}

func insertProduct(db *sql.DB, product *Product) (sql.Result, error) {
	stmt, err := db.Prepare("INSERT INTO products (id, name, price) VALUES (?, ?, ?)")
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	result, err := stmt.Exec(product.ID, product.Name, product.Price)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func updateProduct(db *sql.DB, product *Product) (sql.Result, error) {
	stmt, err := db.Prepare("UPDATE products SET name = ?, price = ? WHERE id = ?")
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	result, err := stmt.Exec(product.Name, product.Price, product.ID)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func getProductByID(db *sql.DB, id string) (*Product, error) {
	row, err := db.Prepare("SELECT id, name, price FROM products WHERE id = ?")

	if err != nil {
		return nil, err
	}

	product := &Product{}
	err = row.QueryRow(id).Scan(&product.ID, &product.Name, &product.Price)

	if err != nil {
		return nil, err
	}

	return product, nil
}

func getProducts(db *sql.DB) ([]*Product, error) {
	rows, err := db.Query("SELECT id, name, price FROM products")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	products := []*Product{}

	for rows.Next() {
		product := &Product{}
		err := rows.Scan(&product.ID, &product.Name, &product.Price)
		if err != nil {
			return nil, err
		}

		products = append(products, product)
	}

	return products, nil
}

func deleteProductByID(db *sql.DB, id string) (sql.Result, error) {
	stmt, err := db.Prepare("DELETE FROM products WHERE id = ?")
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	result, err := stmt.Exec(id)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func truncateProducts(db *sql.DB) (sql.Result, error) {
	stmt, err := db.Prepare("TRUNCATE TABLE products")
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	result, err := stmt.Exec()

	if err != nil {
		return nil, err
	}

	return result, nil
}

func separarLinhas() {
	fmt.Println("--------------------------------------------------")
}
