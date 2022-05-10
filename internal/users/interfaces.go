package users

// Структура пользователя
type User struct {
	General      *GeneralInfo
	Account      *Account
	Productivity *Productivity
}

// Основные данные о сотрудники, используемые в основном для официальных отчетов, справок, договоров
type GeneralInfo struct {
	Lastname string `json:"lastname"`
	Name     string `json:"name"`
	Age      int    `json:"age"`
	Position string `json:"position"`
	Salary   int    `json:""`
}

// Данные учетной записи пользователя
type Account struct {
	ID       string `json:""`
	Login    string `json:""`
	Password string `json:""`
	IsAdmin  bool   `json:""`
}

type Productivity struct {
	PositionsCounter int `json:""`
}
