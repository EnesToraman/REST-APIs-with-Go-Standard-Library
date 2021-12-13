package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"main/data"
	"main/models"
)

// UserHandler handles all the requests are made with containing "/users/" as an URL path.
// Allowed HTTP methods are GET, POST, PUT, DELETE.
func UserHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case http.MethodGet:
		GetUsers(w, r)
	case http.MethodPost:
		CreateUser(w, r)
	case http.MethodPut:
		UpdateUser(w, r)
	case http.MethodDelete:
		DeleteUser(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// GetUsers gets all the user information from Users array and prints them. Function is called via HTTP GET method.
func GetUsers(w http.ResponseWriter, r *http.Request) {
	query := r.URL.RawQuery
	if query == "" {
		res, err := json.Marshal(data.Users)
		CheckError(err)
		fmt.Fprint(w, string(res))
	} else {
		queryParam := query[:strings.IndexByte(query, '=')]
		switch queryParam {
		case "id":
			GetUserByID(w, r)
		case "email":
			GetUserByEmail(w, r)
		case "username":
			GetUserByUsername(w, r)
		case "isActive":
			GetUsersByActivity(w, r)
		default:
			http.Error(w, "Query parameter not allowed", http.StatusNotAcceptable)
		}
	}
}

// CreateUser creates a user with the given information. Function is called via HTTP POST method.
func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	CheckError(err)
	data.Users = append(data.Users, user)
}

// UpdateUser updates the user information with the given ID. Function is called via HTTP PUT method.
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	CheckError(err)
	for i := range data.Users {
		if data.Users[i].ID == user.ID {
			data.Users[i].Email = user.Email
			data.Users[i].Username = user.Username
			data.Users[i].FirstName = user.FirstName
			data.Users[i].LastName = user.LastName
			data.Users[i].IsActive = user.IsActive
		}
	}
}

// DeleteUser deletes the user information with the given ID. Function is called via HTTP DELETE method.
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	CheckError(err)

	// Find the index of given ID, then re-slice to delete.
	index := -1
	for i := range data.Users {
		if data.Users[i].ID == user.ID {
			index = i
		}
	}
	data.Users = append(data.Users[:index], data.Users[index+1:]...)
}

// GetUserByID gets the user by ID field provided in URL query.
// URL should be built with the query "?id=" and the corresponding value in order to get the related user.
func GetUserByID(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	index := -1
	if id != "" {
		userID, err := strconv.Atoi(id)
		CheckError(err)
		for i := range data.Users {
			if data.Users[i].ID == userID {
				index = i
			}
		}
		res, err := json.Marshal(data.Users[index])
		CheckError(err)
		fmt.Fprint(w, string(res))
	}
}

// GetUserByEmail gets the user by Email field provided in URL query.
// URL should be built with the query "?email=" and the corresponding value in order to get the related user.
func GetUserByEmail(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	index := -1
	if email != "" {
		for i := range data.Users {
			if data.Users[i].Email == email {
				index = i
			}
		}
		res, err := json.Marshal(data.Users[index])
		CheckError(err)
		fmt.Fprint(w, string(res))
	}
}

// GetUserByUsername gets the user by Username field provided in URL query.
// URL should be built with the query "?username=" and the corresponding value in order to get the related user.
func GetUserByUsername(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	index := -1
	if username != "" {
		for i := range data.Users {
			if data.Users[i].Username == username {
				index = i
			}
		}
		res, err := json.Marshal(data.Users[index])
		CheckError(err)
		fmt.Fprint(w, string(res))
	}
}

// GetUsersByActivity gets the user(s) by isActive field provided in URL query.
// URL should be built with the query "?isActive=" and the corresponding value in order to get the related user(s).
func GetUsersByActivity(w http.ResponseWriter, r *http.Request) {
	var users []models.User
	isActive := r.FormValue("isActive")
	active, err := strconv.ParseBool(isActive)
	CheckError(err)
	if active {
		for i := range data.Users {
			if data.Users[i].IsActive == active {
				users = append(users, data.Users[i])
			}
		}
	} else {
		for i := range data.Users {
			if data.Users[i].IsActive == active {
				users = append(users, data.Users[i])
			}
		}
	}
	res, err := json.Marshal(users)
	CheckError(err)
	fmt.Fprint(w, string(res))
}
