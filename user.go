package todo

type User struct {
	ID       int    `json:"-"`
	Name     string `json:"name" binding:"required"`     // binding:"required" валидирует наличие полей
	Username string `json:"username" binding:"required"` // являются частью библиотеки гин
	Password string `json:"password" binding:"required"`
}
