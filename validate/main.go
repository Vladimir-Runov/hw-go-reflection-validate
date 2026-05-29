package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
	"unicode"
)

//type User struct {
//    Name  string `validate:"min=3"`
//    Age   int    `validate:"min=18;max=65"`
//    Email string `validate:"regexp=^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\\.[a-zA-Z]{2,}$"`
//}

type User struct {
	Name  string
	Age   int
	Email string
}

// Validate(v interface{}) error, которая проверяет структуру на соответствие правилам, заданным через теги.
func Validate(v interface{}) error {
	return nil
}

func shuffleString(s string) string {
	// Преобразуем строку в срез рун
	runes := []rune(s)

	// Инициализируем генератор случайных чисел
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	// Перетасовываем рун
	for i := len(runes) - 1; i > 0; i-- {
		j := r.Intn(i + 1)
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)
}

// generateRandomString создает строку случайной длины из букв и символов.
func generateRandomString(length int) string {
	const charset_с = "abcdefghijklmnopqrstuvwxyz0123456789!@#$%^&*()_+-=[]{}|;:,.<>?/~`"
	var charset = shuffleString(charset_с)
	var sb strings.Builder // Исправлено здесь

	for i := 0; i < length; i++ {
		sb.WriteByte(charset[rand.Intn(len(charset))])
	}
	return sb.String()
}

func generateRandomString2(length int) string {
	const charset_с = "abcdefghijklmnopqrstuvwxyz"
	var charset = shuffleString(charset_с)
	var sb strings.Builder // Исправлено здесь

	for i := 0; i < length; i++ {
		sb.WriteByte(charset[rand.Intn(len(charset))])
	}
	return sb.String()
}

// generateRandomUser создает случайного пользователя.
func generateRandomUser(id int) User {
	var s = generateRandomString2(3 + rand.Intn(15))
	return User{
		Name:  string(unicode.ToUpper(rune(s[0]))) + s[1:],
		Age:   18 + rand.Intn(65-18), // Случайный возраст от 18 до 65
		Email: generateRandomString(10) + "@example.com",
	}
}

func main() {

	users := make([]User, 10) // Создаем массив из 10 пользователей

	users[0] = User{Name: "Jo", Age: 17, Email: "invalid-email"}
	users[1] = User{Name: "Jo_Ok", Age: 18 + rand.Intn(65-18), Email: "invalid@email.su"}
	for i := 2; i < 10; i++ {
		users[i] = generateRandomUser(i + 1) // Заполняем массив случайными пользователями
	}

	// Выводим содержимое массива пользователей в консоль
	for i, user := range users {
		fmt.Printf("Структура %d: %+v\n", i+1, user)
	}

	//
	//
	//    if err := Validate(&user); err != nil {
	//        fmt.Println("Validation error:", err)
	//    } else {
	//        fmt.Println("Validation passed!")
	//    }
}
