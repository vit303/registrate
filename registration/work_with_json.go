package registration

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

// Функция для добавления нового пользователя
func AddUser(username, password, role string) {
	users := loadUsers()

	// Хэшируем пароль и генерируем соль
	hashedPassword := hashPassword(password)
	salt := generateSalt()

	// Создаем нового пользователя
	newUser := User{
		Username:       username,
		HashedPassword: hashedPassword,
		Salt:           salt,
		Role:           role,
	}

	// Добавляем пользователя в коллекцию
	users.Users = append(users.Users, newUser)

	// Сохраняем коллекцию пользователей в файл
	saveUsers(users)

	fmt.Println("Пользователь зарегистрирован!")
}

// Функция для загрузки данных пользователей из JSON
func loadUsers() UserCollection {
	file, err := ioutil.ReadFile("users.json")
	if err != nil {
		if os.IsNotExist(err) {
			// Если файл не существует, возвращаем пустую коллекцию
			return UserCollection{}
		}
		log.Fatal(err)
	}

	var users UserCollection
	err = json.Unmarshal(file, &users)
	if err != nil {
		log.Fatal(err)
	}
	return users
}

// Функция для сохранения коллекции пользователей в JSON
func saveUsers(users UserCollection) {
	file, err := json.MarshalIndent(users, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	err = ioutil.WriteFile("users.json", file, 0644)
	if err != nil {
		log.Fatal(err)
	}
}
