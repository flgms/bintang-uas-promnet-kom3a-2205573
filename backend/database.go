package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

type InventoryItem struct {
	ID           int    `json:"id"`
	NamaBarang   string `json:"nama_barang"`
	Jumlah	 	 int    `json:"jumlah"`
	Harga_satuan int 	`json:"harga_satuan"`
	Lokasi 	 	 string `json:"lokasi"`
	Deskripsi 	 string `json:"deskripsi"`
}

func GetInventoryItems(w http.ResponseWriter, r *http.Request){
	rows, err := db.Query("SELECT * FROM inventory_items")
	if err != nil{
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
}
defer rows.Close()

items := []InventoryItem{}
for rows.Next() {
	var item InventoryItem
	err := rows.Scan(&item.ID, &item.NamaBarang, &item.Jumlah, &item.Harga_satuan, &item.Lokasi, &item.Deskripsi)
	if err != nil{
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	items = append(items, item)
}
json.NewEncoder(w).Encode(items)
}

func PostInventoryItems(w http.ResponseWriter, r http.Request){
	var item InventoryItem
	err := json.NewDecoder(r.Body).Decode(&item)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	stmt, err := db.Prepare("INSERT INT0 inventory_items (nama_barang, jumlah, harga_satuan, lokasi, deskripsi) VALUES (? ? ? ?)")
	if err != nil{
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer stmt.Close()

	result, err := stmt.Exec(item.NamaBarang, item.Jumlah, item.Harga_satuan, item.Lokasi, item.Deskripsi)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(struct {
		ID int `json:"id"`
	}{
		ID: int(rowsAffected),
	})
}

func PutInventoryItems(w http.ResponseWriter, r *http.Request){
	var item InventoryItem
	err := json.NewDecoder(r.Body).Decode(&item)
	if err != nil{
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil{
		log.Println(err)
		http.Error(w, "Failed to get id from request", http.StatusBadRequest)
		return
	}

	stmt, err := db.Prepare("UPDATE inventory_items SET nama_barang=? jumlah=?, harga_satuan=?, lokasi=?, deskripsi=? WHERE id=?")
	if err != nil{
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return 
	}
	defer stmt.Close()

	result, err := stmt.Exec(item.NamaBarang, item.Jumlah, item.Harga_satuan, item.Lokasi, item.Deskripsi, id)
	if err != nil{
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil{
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if rowsAffected == 0 {
		http.Error(w, http.StatusNotFound, nill)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func DeleteinventoryItems(w http.ResponseWriter, r *http.Request){
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	stmt, err := db.Prepare("DELETE FROM inventory_item WHERE id=?")
	if err != nil{
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer stmt.Close()

	result, err := stmt.Exec(id)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil{
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if rowsAffected == 0 {
		http.Error(w, http.StatusNotFound, nil)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func main(){
	db, err := sql.Open("mysql", "root:@tcp(localhost3306)/db_2205573_bintang_uas")
	if err != nil{
		fmt.Println(err)
		return
	}
	defer db.Close()

	router := mux.NewRouter()
	router.HandleFunc("/inventory_items", GetInventoryItems).Methods("GET")
	router.HandleFunc("/inventory_items", PostInventoryItems).Methods("POST")
	router.HandleFunc("/inventory_items", PutInventoryItems).Methods("PUT")
	router.HandleFunc("/inventory_items", DeleteinventoryItems).Methods("DELETE")


	log.Fatal(http.ListenAndServe(":8080", router))
}

