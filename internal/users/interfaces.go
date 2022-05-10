package users

// Структура пользователя
type User struct {
	Account *Account
	General *GeneralInfo
}

// Основные данные о сотрудники, используемые в основном для официальных отчетов, справок, договоров
type GeneralInfo struct {
	Lastname string
	Name     string
	Age      int
	Position string
	Salary   int
}

// Данные учетной записи пользователя
type Account struct {
	ID       string
	Login    string
	Password string
	IsAdmin  bool
}
