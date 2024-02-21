package main

import (
	"database/sql"
	"fmt"

	//Esse underline é blank identfy, que significa que ele vai ignorar esse import até o momento de ele ser necessário
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
)

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
	db, err := sql.Open("mysql", "root:Stepflish123@tcp(localhost:3306)/financeiro")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	product := NewProduct("NoteBook", 1899.90)
	err = insertProduct(db, product)
	if err != nil {
		panic(err)
	}

	product.Price = 100.00
	err = updateProducts(db, product)
	if err != nil {
		panic(err)
	}

	/*p, err := selectOneProduct(db, product.ID)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Product: %v, possui o preço de %.2f", p.Name, p.Price)*/

	products, err := selectAllProducts(db)
	if err != nil {
		panic(err)
	}
	for _, p := range products {
		fmt.Printf("Product: %v, possui o preço de %.2f\n", p.Name, p.Price)
	}

	err = deleteProduct(db, product.ID)
	if err != nil {
		panic(err)
	}
}

func insertProduct(db *sql.DB, product *Product) error {
	stmt, err := db.Prepare("insert into products(id, name, price) values(?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(product.ID, product.Name, product.Price)
	if err != nil {
		return err
	}
	return nil
}

func updateProducts(db *sql.DB, product *Product) error {
	stmt, err := db.Prepare("update products set name = ?, price = ? where id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(product.Name, product.Price, product.ID)
	if err != nil {
		return err
	}
	return nil
}

func selectOneProduct( /*ctx context.Context,*/ db *sql.DB, id string) (*Product, error) {
	stmt, err := db.Prepare("select id, name, price from products where id = ?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	var p Product
	/*Dessa forma o ctx entra nesse caso como um contador para por exemplo dar um tempo em uma pesquisa, então se demorar
	muito para fazer a pesquisa ele encerraria a pesquisa*/
	//err = stmt.QueryRowContext(ctx, id).Scan(&p.ID, &p.Name, &p.Price)
	// QueryRow é busca de uma linha apenas
	err = stmt.QueryRow(id).Scan(&p.ID, &p.Name, &p.Price)
	if err != nil {
		return nil, err
	}
	return &p, nil
}

func selectAllProducts(db *sql.DB) ([]Product, error) {
	row, err := db.Query("select id, name, price from products")
	if err != nil {
		return nil, err
	}
	defer row.Close()
	var products []Product
	for row.Next() {
		var p Product
		err = row.Scan(&p.ID, &p.Name, &p.Price)
		if err != nil {
			return nil, err
		}
		products = append(products, p)
	}
	return products, nil
}

func deleteProduct(db *sql.DB, id string) error {
	stmt, err := db.Prepare("delete from products where id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}
	//Esse "return nil" siginifica que ele vai retornar um erro em branco para dizer que deu certo a consulta
	return nil
}
