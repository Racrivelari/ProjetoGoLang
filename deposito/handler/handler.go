package handler

import (
	// "deposito/entity"
	"deposito/banco"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	_ "github.com/gorilla/mux"
)

type product struct { //NOME DE VARIAVEL TEM Q SER MAISUCULA POR CAUSA DO JSON
	ID        uint32 `json:"id"`
	Name      string `json:"name_prod"`
	Price     string `json:"price_prod"`
	Code      string `json:"code_prod"`
	CreatedAt string `json:"created_at,omitempty"`
	UpdatedAt string `json:"updated_at,omitempty"`
}

// insert product to database
func CreateProduct(w http.ResponseWriter, r *http.Request) {
	corpoRequest, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		w.Write([]byte("Falha ao ler requisição"))
		return
	}

	var product product

	if erro = json.Unmarshal(corpoRequest, &product); erro != nil {
		w.Write([]byte("Erro ao converter produto em struct"))
		return
	}

	fmt.Println(product)

	db, erro := banco.Conectar()
	if erro != nil {
		w.Write([]byte("Erro ao conectar no bd"))
		return
	}

	defer db.Close()

	statement, erro := db.Prepare("INSERT INTO Product (name_prod, price_prod, code_prod) values (?, ?, ?)")
	if erro != nil {
		w.Write([]byte("Erro ao criar statement"))
		return
	}
	defer statement.Close()

	insercao, erro := statement.Exec(product.Name, product.Price, product.Code)
	if erro != nil {
		w.Write([]byte("Erro ao executar statement"))
		return
	}

	idInserido, erro := insercao.LastInsertId()
	if erro != nil {
		w.Write([]byte("Erro ao buscar o ultimo id inserido "))
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("Produto inserido com sucesso, id: %d", idInserido)))

}

// look for all the products on database
func GetAllProducts(w http.ResponseWriter, r *http.Request) {

	db, erro := banco.Conectar()
	if erro != nil {
		w.Write([]byte("Erro ao conectar no bd"))
		return
	}
	defer db.Close()

	linhas, erro := db.Query("SELECT * FROM PRODUCT")
	if erro != nil {
		w.Write([]byte("Erro ao realizar busca"))
		return
	}
	defer linhas.Close()

	var products []product
	for linhas.Next() {
		var product product

		if erro := linhas.Scan(&product.ID, &product.Name, &product.Price, &product.Code, &product.CreatedAt, &product.UpdatedAt); erro != nil {
			w.Write([]byte("Erro escanear o produto"))
			return
		}

		products = append(products, product)
	}

	w.WriteHeader(http.StatusOK)
	if erro := json.NewEncoder(w).Encode(products); erro != nil {
		w.Write([]byte("Erro ao converter produtos p json"))
		return
	}
}
