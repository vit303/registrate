package main

import (
	"fmt"

	"registrate/registration"
)

func main() {
	for {
		// Меню
		fmt.Println("1. Регистрация")
		fmt.Println("2. Войти")
		fmt.Println("3. Проверить права доступа")
		fmt.Println("4. Выход")
		fmt.Print("Выберите опцию: ")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			// Регистрация нового пользователя
			var username, password, role string
			fmt.Print("Введите логин: ")
			fmt.Scanln(&username)
			fmt.Print("Введите пароль: ")
			fmt.Scanln(&password)
			fmt.Print("Введите роль (user или admin): ")
			fmt.Scanln(&role)
			registration.AddUser(username, password, role)

		case 2:
			// Аутентификация пользователя
			var username, password string
			fmt.Print("Введите логин: ")
			fmt.Scanln(&username)
			fmt.Print("Введите пароль: ")
			fmt.Scanln(&password)

			if registration.AuthenticateUser(username, password) {
				fmt.Println("Вход успешен!")
			} else {
				fmt.Println("Неверный логин или пароль.")
			}

		case 3:
			// Проверка прав доступа
			var username string
			fmt.Print("Введите логин для проверки прав доступа: ")
			fmt.Scanln(&username)
			registration.CheckPermissions(username)

		case 4:
			// Выход из программы
			fmt.Println("Выход...")
			return

		default:
			// Если введен неверный выбор
			fmt.Println("Неверный выбор, попробуйте снова.")
		}
	}
}
