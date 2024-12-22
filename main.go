package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"registrate/registration"
)

type UserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role,omitempty"` // Role is optional for login
}

type PermissionRequest struct {
	Username string `json:"username"`
}

func main() {
	http.HandleFunc("/register", registerHandler)
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/check_permissions", checkPermissionsHandler)

	fmt.Println("Сервер запущен на порту 8080")
	http.ListenAndServe(":8080", nil)
}

func registerHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
		return
	}

	var userRequest UserRequest
	if err := json.NewDecoder(r.Body).Decode(&userRequest); err != nil {
		http.Error(w, "Неверный запрос", http.StatusBadRequest)
		return
	}

	registration.AddUser(userRequest.Username, userRequest.Password, userRequest.Role)
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "Пользователь %s зарегистрирован", userRequest.Username)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
		return
	}

	var userRequest UserRequest
	if err := json.NewDecoder(r.Body).Decode(&userRequest); err != nil {
		http.Error(w, "Неверный запрос", http.StatusBadRequest)
		return
	}

	if registration.AuthenticateUser(userRequest.Username, userRequest.Password) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Вход успешен для %s", userRequest.Username)
	} else {
		http.Error(w, "Неверный логин или пароль", http.StatusUnauthorized)
	}
}

func checkPermissionsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
		return
	}

	var permissionRequest PermissionRequest
	if err := json.NewDecoder(r.Body).Decode(&permissionRequest); err != nil {
		http.Error(w, "Неверный запрос", http.StatusBadRequest)
		return
	}

	registration.CheckPermissions(permissionRequest.Username)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Проверка прав доступа для %s выполнена", permissionRequest.Username)
}
