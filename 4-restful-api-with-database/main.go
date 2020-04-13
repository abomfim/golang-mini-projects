package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"

	"github.com/julienschmidt/httprouter"
	uuid "github.com/satori/go.uuid"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Username  string    `json:"userName"`
}

func dbConnect() (db *sql.DB) {
	db, err := sql.Open("sqlite3", "./devdb.db")
	checkErr(err)
	return db
}

func getUsers(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	db := dbConnect()
	rows, err := db.Query("SELECT * FROM users")
	checkErr(err)

	user := User{}
	res := []User{}
	for rows.Next() {
		var id uuid.UUID
		var username, firstName, lastName string
		err = rows.Scan(&id, &username, &firstName, &lastName)
		checkErr(err)

		user.ID = id
		user.Username = username
		user.FirstName = firstName
		user.LastName = lastName
		res = append(res, user)
	}
	defer db.Close()

	responseJSON(w, http.StatusOK, res)
}

func createUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var user User
	_ = json.NewDecoder(r.Body).Decode(&user)
	user.ID = uuid.NewV4()
	db := dbConnect()
	stmt, _ := db.Prepare("INSERT INTO users(id, username, firstname, lastname) VALUES(?,?,?,?)")

	_, err := stmt.Exec(user.ID, user.Username, user.FirstName, user.LastName)
	checkErr(err)
	defer db.Close()
	responseJSON(w, http.StatusCreated, user)
}

func getUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	db := dbConnect()
	paramID := params.ByName("id")
	rows, err := db.Query("SELECT * FROM users WHERE id = ?", paramID)
	checkErr(err)

	user := User{}
	res := []User{}
	for rows.Next() {
		var id uuid.UUID
		var username, firstName, lastName string
		err = rows.Scan(&id, &username, &firstName, &lastName)
		checkErr(err)

		user.ID = id
		user.Username = username
		user.FirstName = firstName
		user.LastName = lastName
		res = append(res, user)
	}
	defer db.Close()

	responseJSON(w, http.StatusOK, res)
}

func updateUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	var user User
	paramID, _ := uuid.FromString(params.ByName("id"))
	user.ID = paramID
	_ = json.NewDecoder(r.Body).Decode(&user)
	db := dbConnect()
	stmt, _ := db.Prepare("UPDATE users SET username = ?, firstname =?, lastname = ? WHERE id = ?")

	_, err := stmt.Exec(user.Username, user.FirstName, user.LastName, paramID)
	checkErr(err)
	defer db.Close()
	responseJSON(w, http.StatusOK, user)
}

func deleteUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	paramID, _ := uuid.FromString(params.ByName("id"))
	db := dbConnect()
	stmt, _ := db.Prepare("DELETE FROM users WHERE id = ?")
	_, err := stmt.Exec(paramID)
	checkErr(err)
	defer db.Close()
	responseJSON(w, http.StatusNoContent, nil)
}

func responseJSON(w http.ResponseWriter, status int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write([]byte(response))
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	router := httprouter.New()
	router.GET("/api/users", getUsers)
	router.POST("/api/user", createUser)
	router.GET("/api/user/:id", getUser)
	router.PUT("/api/user/:id", updateUser)
	router.DELETE("/api/user/:id", deleteUser)

	log.Fatal(http.ListenAndServe(":5555", router))
}
