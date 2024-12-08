package registration

import "fmt"

// Функция для аутентификации пользователя
func AuthenticateUser(username, password string) bool {
	users := loadUsers()

	// Проходим по всем пользователям и проверяем хэш пароля
	for _, user := range users.Users {
		if user.Username == username && user.HashedPassword == hashPassword(password) {
			return true
		}
	}

	return false
}

// Проверка прав доступа
func CheckPermissions(username string) {
	users := loadUsers()

	// Ищем пользователя в базе данных
	for _, user := range users.Users {
		if user.Username == username {
			// Если пользователь найден
			fmt.Printf("Пользователь %s существует. Права доступа подтверждены.\n", username)
			return
		}
	}

	// Если пользователя нет
	fmt.Printf("Пользователь %s не найден.\n", username)
}
