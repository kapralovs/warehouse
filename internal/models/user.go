package models

// Структура пользователя
type User struct {
	// Идентификатор
	ID string `json:"id"`
	// Фамилия
	Lastname string `json:"lastname"`
	// Имя
	Name string `json:"name"`
	// Возраст
	Age int `json:"age"`
	// Роль
	Role string `json:"role"`
	// Зарплата
	Salary int `json:"Salary"`
	// Имя учетной записи
	Username string `json:"username"`
	// Пароль
	Password string `json:"password"`
	// Является ли админом
	IsAdmin bool `json:"is_admin"`
	// Статус пользователя
	IsOnline bool `json:"is_online"`
}
