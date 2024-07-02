package server

import (
	"crud/db"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type user struct {
	ID    uint32 `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	requestBody, err := io.ReadAll(r.Body)
	if err != nil {
		w.Write([]byte("Failed to read the request body"))
		return
	}

	var user user
	if err = json.Unmarshal(requestBody, &user); err != nil {
		w.Write([]byte("Failed to parse request body to user"))
		return
	}

	db, err := db.Connect()
	if err != nil {
		w.Write([]byte("Failed to connect to database"))
		return
	}
	defer db.Close()

	statement, err := db.Prepare("insert into users (name, email) values (?, ?)")
	if err != nil {
		w.Write([]byte("Failed to prepare insert statement"))
		return
	}
	defer statement.Close()

	insertResult, err := statement.Exec(user.Name, user.Email)
	if err != nil {
		w.Write([]byte("Failed to execute insert statement"))
		return
	}

	id, err := insertResult.LastInsertId()
	if err != nil {
		w.Write([]byte("Failed to get inserted id"))
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("User inserted successfully! Id: %d", id)))
}

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	db, err := db.Connect()
	if err != nil {
		w.Write([]byte("Failed to connect to database"))
		return
	}
	defer db.Close()

	lines, err := db.Query("select * from users")
	if err != nil {
		w.Write([]byte("Failed to query users"))
		return
	}
	defer lines.Close()

	var users []user
	for lines.Next() {
		var user user
		if err := lines.Scan(&user.ID, &user.Name, &user.Email); err != nil {
			w.Write([]byte("Failed to scan user data"))
			return
		}

		users = append(users, user)
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(users); err != nil {
		w.Write([]byte("Failed to parse users to JSON"))
		return
	}
}

func GetUserById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	ID, err := strconv.ParseUint(params["id"], 10, 32)
	if err != nil {
		w.Write([]byte("Failed to get ID path param"))
		return
	}

	db, err := db.Connect()
	if err != nil {
		w.Write([]byte("Failed to connect to database"))
		return
	}
	defer db.Close()

	line, err := db.Query("select * from users where id = ?", ID)
	if err != nil {
		w.Write([]byte("Failed to query user"))
		return
	}

	var user user
	if line.Next() {
		if err := line.Scan(&user.ID, &user.Name, &user.Email); err != nil {
			w.Write([]byte("Failed to scan user data"))
			return
		}
	}

	if user.ID == 0 {
		w.WriteHeader(http.StatusNotFound)

	} else {

		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(user); err != nil {
			w.Write([]byte("Failed to parse users to JSON"))
			return
		}
	}
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	// Read path variable
	params := mux.Vars(r)

	ID, err := strconv.ParseUint(params["id"], 10, 32)
	if err != nil {
		w.Write([]byte("Failed to get ID path param"))
		return
	}

	// Read request body
	requestBody, err := io.ReadAll(r.Body)
	if err != nil {
		w.Write([]byte("Failed to read the request body"))
		return
	}

	var user user
	if err = json.Unmarshal(requestBody, &user); err != nil {
		w.Write([]byte("Failed to parse request body to user"))
		return
	}

	// Open DB connection
	db, err := db.Connect()
	if err != nil {
		w.Write([]byte("Failed to connect to database"))
		return
	}
	defer db.Close()

	// Prepare and execute update statement
	statement, err := db.Prepare("update users set name = ?, email = ? where id = ?")
	if err != nil {
		w.Write([]byte("Failed to prepare update statement"))
		return
	}
	defer statement.Close()

	if _, err := statement.Exec(user.Name, user.Email, ID); err != nil {
		w.Write([]byte("Failed to execute update statement"))
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	ID, err := strconv.ParseUint(params["id"], 10, 32)
	if err != nil {
		w.Write([]byte("Failed to get ID path param"))
		return
	}

	db, err := db.Connect()
	if err != nil {
		w.Write([]byte("Failed to connect to database"))
		return
	}
	defer db.Close()

	// Prepare and execute update statement
	statement, err := db.Prepare("delete from users where id = ?")
	if err != nil {
		w.Write([]byte("Failed to prepare delete statement"))
		return
	}
	defer statement.Close()

	if _, err := statement.Exec(ID); err != nil {
		w.Write([]byte("Failed to execute delete statement"))
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
