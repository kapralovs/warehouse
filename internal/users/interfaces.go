package users

// Структура пользователя
type Profile struct {
	General      *GeneralInfo
	Account      *Account
	Productivity *Productivity
}

// Основные данные о сотрудники, используемые в основном для официальных отчетов, справок, договоров
type GeneralInfo struct {
	Lastname string `json:"lastname"`
	Name     string `json:"name"`
	Age      int    `json:"age"`
	Role     string `json:"role"`
	Salary   int    `json:"Salary"`
}

// Данные учетной записи пользователя
type Account struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	IsAdmin  bool   `json:"is_admin"`
	IsOnline bool   `json:"is_online"`
}

type Productivity struct {
	PositionsCounter int `json:"PositionsCounter"`
}
