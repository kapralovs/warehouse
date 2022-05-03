package users

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
	Login    string
	Password string
}
