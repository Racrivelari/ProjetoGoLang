package handler

import (
	"deposito/banco"
	"deposito/entity" //pra esse import funcionar, na struct devo declarar a primeira letra do nome dela como maiuscula, ex: Product, invés de product, isso q determina se ela é public ou private
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"net/http"
	"strconv"
)

// insert product to database
func CreateProduct(w http.ResponseWriter, r *http.Request) {
	corpoRequest, erro := io.ReadAll(r.Body)
	if erro != nil {
		w.Write([]byte("Falha ao ler requisição"))
		return
	}
	var product entity.Product

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

	linhas, erro := db.Query("SELECT id_prod, name_prod, price_prod, code_prod FROM Product")
	if erro != nil {
		w.Write([]byte("Erro ao realizar busca"))
		return
	}
	defer linhas.Close()

	var products []entity.Product
	for linhas.Next() {
		var product entity.Product
		if erro := linhas.Scan(&product.ID, &product.Name, &product.Price, &product.Code); erro != nil {
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

// get specific product data on database
func GetById(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)

	ID, erro := strconv.ParseUint(param["id"], 10, 32)
	if erro != nil {
		w.Write([]byte("Erro converter param pra int"))
	}

	db, erro := banco.Conectar()
	if erro != nil {
		w.Write([]byte("Erro ao conectar no bd"))
		return
	}
	defer db.Close()

	linha, erro := db.Query("SELECT id_prod, name_prod, price_prod, code_prod FROM Product where id_prod = ?", ID)
	if erro != nil {
		w.Write([]byte("Erro ao buscar o produto"))
		return
	}

	var product entity.Product
	if linha.Next() {
		if erro := linha.Scan(&product.ID, &product.Name, &product.Price, &product.Code); erro != nil {
			w.Write([]byte("Erro ao escanear produto"))
			return
		}

	}

	if erro := json.NewEncoder(w).Encode(product); erro != nil {
		w.Write([]byte("Erro ao converter produto p json"))
		return
	}

}

// delete specific product from database
func DeleteById(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)

	ID, erro := strconv.ParseUint(param["id"], 10, 32)
	if erro != nil {
		w.Write([]byte("Erro converter param pra int"))
	}

	db, erro := banco.Conectar()
	if erro != nil {
		w.Write([]byte("Erro ao conectar no bd"))
	}
	defer db.Close()

	statement, erro := db.Prepare("DELETE FROM Logs where id_prod = ?")
	if erro != nil {
		w.Write([]byte("Erro ao criar statement"))
	}
	defer statement.Close()

	statement.Exec(ID)

	statement2, erro := db.Prepare("DELETE FROM Product where id_prod = ?")
	if erro != nil {
		w.Write([]byte("Erro ao criar statement"))
	}
	defer statement2.Close()

	if _, erro := statement2.Exec(ID); erro != nil {
		w.Write([]byte(erro.Error()))
		return
	}

	w.Write([]byte(fmt.Sprintf("Produto deletado com sucesso, id: %d", ID)))

}

// update specific product from database
func UpdateById(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)

	ID, erro := strconv.ParseUint(param["id"], 10, 32)
	if erro != nil {
		w.Write([]byte("Erro converter param pra int"))
	}

	corpoRequest, erro := io.ReadAll(r.Body)
	if erro != nil {
		w.Write([]byte("Erro ao ler corpo da requisicao"))
		return
	}

	var product entity.Product
	if erro := json.Unmarshal(corpoRequest, &product); erro != nil {
		w.Write([]byte("Erro ao converter produto em struct"))
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		w.Write([]byte("Erro ao conectar no bd"))
	}
	defer db.Close()

	statement, erro := db.Prepare("UPDATE Product SET name_prod = ?, price_prod = ? where id_prod = ?")
	if erro != nil {
		w.Write([]byte("Erro ao criar statement"))
		return
	}
	defer statement.Close()

	if _, erro := statement.Exec(product.Name, product.Price, ID); erro != nil {
		w.Write([]byte("Erro ao atualizar o produto"))
		return
	}

	w.Write([]byte(fmt.Sprintf("Produto atualizado com sucesso, id: %d", ID)))

}
