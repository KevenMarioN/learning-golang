package server

import (
	"crud-basico/database"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type userType struct {
	ID    uuid.UUID `json:"id"`
	Name  string    `json:"name"`
	Email string    `json:"email"`
}

func CreateUser(write http.ResponseWriter, request *http.Request) {
	body, err := io.ReadAll(request.Body)
	if err != nil {
		write.WriteHeader(http.StatusBadRequest)
		write.Write([]byte("Error on read requestr body"))
		return
	}

	var newUser userType

	if err = json.Unmarshal(body, &newUser); err != nil {
		write.WriteHeader(http.StatusBadRequest)
		write.Write([]byte("Error on parsed struct user"))
		return
	}
	id := uuid.New()
	newUser.ID = id
	db, err := database.Conectar()
	if err != nil {
		write.WriteHeader(http.StatusBadRequest)
		write.Write([]byte("Error on connected in database"))
		return
	}
	defer db.Close()

	statement, err := db.Prepare("insert into users (id,name,email) values ($1,$2,$3)")

	if err != nil {
		write.WriteHeader(http.StatusBadRequest)
		write.Write([]byte("Error on prepare statement for insert database"))
		return
	}
	defer statement.Close()

	insert, err := statement.Exec(newUser.ID, newUser.Name, newUser.Email)

	if err != nil {
		write.WriteHeader(http.StatusBadRequest)
		write.Write([]byte(fmt.Sprintf("Error on execute o insert!\n%s", err.Error())))
		return
	}

	_, err = insert.RowsAffected()
	if err != nil {
		fmt.Println(err.Error())
		write.WriteHeader(http.StatusBadRequest)
		write.Write([]byte("Error on get ID insert"))
		return
	}

	write.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(write).Encode(newUser); err != nil {
		write.WriteHeader(http.StatusBadRequest)
		write.Write([]byte("Error on converted user to JSON"))
		return
	}
}

func GetAllUser(write http.ResponseWriter, request *http.Request) {
	db, err := database.Conectar()
	if err != nil {
		write.WriteHeader(http.StatusBadRequest)
		write.Write([]byte("Error on connected in database"))
		return
	}
	defer db.Close()

	result, err := db.Query("select * from users")

	if err != nil {
		write.WriteHeader(http.StatusBadRequest)
		write.Write([]byte("Error on get users"))
		return
	}
	defer result.Close()

	var users []userType

	for result.Next() {
		var user userType
		if err := result.Scan(&user.ID, &user.Name, &user.Email); err != nil {
			write.WriteHeader(http.StatusBadRequest)
			write.Write([]byte("Error on scanner user"))
			return
		}
		users = append(users, user)
	}
	write.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(write).Encode(users); err != nil {
		write.WriteHeader(http.StatusBadRequest)
		write.Write([]byte("Error on converted slice to JSON"))
		return
	}
}
func GetUserId(write http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)

	ID, err := uuid.Parse(params["id"])
	if err != nil {
		write.Write([]byte("Error ID is not type uuid"))
	}

	db, err := database.Conectar()
	if err != nil {
		write.WriteHeader(http.StatusBadRequest)
		write.Write([]byte("Error on connected in database"))
		return
	}
	defer db.Close()

	row := db.QueryRow("select * from users where id = $1", ID)
	if err != nil {
		write.WriteHeader(http.StatusBadRequest)
		write.Write([]byte("Error on  get user!"))
		return
	}

	var user userType

	if err := row.Scan(&user.ID, &user.Name, &user.Email); err != nil {
		if err == sql.ErrNoRows {
			write.WriteHeader(http.StatusBadRequest)
			write.Write([]byte("Error user not exists!"))
			return
		}
		write.WriteHeader(http.StatusBadRequest)
		write.Write([]byte("Error on scanner the user!"))
		return
	}

	write.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(write).Encode(user); err != nil {
		write.WriteHeader(http.StatusBadRequest)
		write.Write([]byte("Error on converted user for JSON!"))
		return
	}
}

func UpdateUser(write http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)

	ID, err := uuid.Parse(params["id"])
	if err != nil {
		write.WriteHeader(http.StatusBadRequest)
		write.Write([]byte("Error on convert id params to UUID"))
		return
	}

	body, err := io.ReadAll(request.Body)
	if err != nil {
		write.WriteHeader(http.StatusBadRequest)
		write.Write([]byte("Error on convert body"))
		return
	}

	var userUpdated userType

	if err := json.Unmarshal(body, &userUpdated); err != nil {
		write.WriteHeader(http.StatusBadRequest)
		write.Write([]byte("Error on convert body to JSON"))
		return
	}

	db, err := database.Conectar()
	if err != nil {
		write.WriteHeader(http.StatusBadRequest)
		write.Write([]byte("Error on connected database!"))
		return
	}
	defer db.Close()
	stmt, err := db.Prepare("update users set name = $1, email = $2 where id = $3")
	if err != nil {
		write.WriteHeader(http.StatusBadRequest)
		write.Write([]byte("Error on created statement!"))
		return
	}
	defer stmt.Close()
	if _, err := stmt.Exec(userUpdated.Name, userUpdated.Email, ID); err != nil {
		write.WriteHeader(http.StatusNoContent)
		write.Write([]byte("Error on atualization the user"))
		return
	}
	write.WriteHeader(http.StatusNoContent)
}

func DeleteUser(write http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)

	ID, err := uuid.Parse(params["id"])
	if err != nil {
		write.WriteHeader(http.StatusBadRequest)
		write.Write([]byte("Error on convert id params for UUID!"))
		return
	}

	db, err := database.Conectar()
	if err != nil {
		write.WriteHeader(http.StatusBadRequest)
		write.Write([]byte("Error on connected database!"))
	}
	defer db.Close()

	stmt, err := db.Prepare("DELETE from users where id = $1")
	if err != nil {
		write.WriteHeader(http.StatusBadRequest)
		write.Write([]byte("Error on created statement!"))
		return
	}
	defer stmt.Close()
	if _, err = stmt.Exec(ID); err != nil {
		write.WriteHeader(http.StatusBadRequest)
		write.Write([]byte("Error on deleted"))
		return
	}

	write.WriteHeader(http.StatusNoContent)
}
