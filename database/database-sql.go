package database

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
)

func NewProduct(name string, price float64) *Product {
	return &Product{
		ID:    uuid.New().String(),
		Name:  name,
		Price: price,
	}
}

func RunDatabaseSql() {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/goexpert")

	if err != nil {
		panic(err)
	}

	defer db.Close()

	fmt.Println("Inserindo um produto")

	tx, err := db.Begin()

	product := NewProduct("Product", 1899.0)

	if err = insertProduct(tx, *product); err != nil {
		fmt.Println("Rolling back changes because some error occurs.")
		if err := tx.Rollback(); err != nil {
			panic(err)
		}
		panic(err)
	}

	product.Name = "Notebook"
	product.Price = 2089.99

	fmt.Printf("Atualizando um produto %v\n", product.Name)

	if err = updateProduct(tx, *product); err != nil {
		fmt.Println("Rolling back changes because some error occurs.")
		if err := tx.Rollback(); err != nil {
			panic(err)
		}
		panic(err)
	}

	fmt.Printf("Buscando o produto %v\n", product.ID)

	if product, err = getOneProduct(tx, product.ID); err != nil {
		fmt.Println("Rolling back changes because some error occurs.")
		if err := tx.Rollback(); err != nil {
			panic(err)
		}
		panic(err)
	}

	fmt.Printf("Product: %v, possui preço de %.2f\n", product.Name, product.Price)

	product = NewProduct("Mouse", 500.0)

	fmt.Printf("Inserindo o produto %v\n", product.Name)

	if err = insertProduct(tx, *product); err != nil {
		fmt.Println("Rolling back changes because some error occurs.")
		if err := tx.Rollback(); err != nil {
			panic(err)
		}
		panic(err)
	}

	fmt.Println("Excluindo produto")

	if err = deleteProduct(tx, product.ID); err != nil {
		fmt.Println("Rolling back changes because some error occurs.")
		if err := tx.Rollback(); err != nil {
			panic(err)
		}
		panic(err)
	}

	fmt.Println("Buscando todos os produtos")

	var products []Product
	if products, err = getAllProducts(tx); err != nil {
		fmt.Println("Rolling back changes because some error occurs.")
		if err := tx.Rollback(); err != nil {
			panic(err)
		}
		panic(err)
	}

	for _, p := range products {
		fmt.Printf("Product: %v, possui preço de %.2f\n", p.Name, p.Price)
	}

	if err = tx.Commit(); err != nil {
		panic(err)
	}
}

func deleteProduct(tx *sql.Tx, id string) error {
	stmt, err := tx.Prepare("delete from products where id = ?")

	if err != nil {
		return err
	}

	defer stmt.Close()

	if _, err = stmt.Exec(id); err != nil {
		return err
	}

	return nil
}

func getAllProducts(tx *sql.Tx) ([]Product, error) {
	rows, err := tx.Query("select id, name, price from products")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var products []Product

	for rows.Next() {
		var product Product

		if err = rows.Scan(&product.ID, &product.Name, &product.Price); err != nil {
			return nil, err
		}

		products = append(products, product)
	}

	return products, nil
}

func getOneProduct(tx *sql.Tx, id string) (*Product, error) {
	stmt, err := tx.Prepare(`SELECT id, name, price FROM products WHERE id = ?`)

	if err != nil {
		panic(err)
	}

	defer stmt.Close()

	var product Product

	if err = stmt.QueryRow(id).Scan(&product.ID, &product.Name, &product.Price); err != nil {
		return nil, err
	}

	return &product, nil
}

func updateProduct(tx *sql.Tx, product Product) error {
	stmt, err := tx.Prepare(`UPDATE products SET name = ?, price = ? WHERE id = ?`)

	if err != nil {
		panic(err)
	}

	defer stmt.Close()

	if _, err = stmt.Exec(product.Name, product.Price, product.ID); err != nil {
		return err
	}

	return nil
}

func insertProduct(tx *sql.Tx, product Product) error {
	stmt, err := tx.Prepare(`INSERT INTO products(id, name, price) values (?, ?, ?)`)

	if err != nil {
		panic(err)
	}

	defer stmt.Close()

	if _, err = stmt.Exec(product.ID, product.Name, product.Price); err != nil {
		return err
	}

	return nil
}
