package todo

type User struct {
	ID           int    `json:"-" db:"id"`
	Username     string `json:"username" binding:"required"` // binding:"required" валидирует наличие полей
	Password     string `json:"password" binding:"required"` // являются частью библиотеки гин
	UserIP       string `json:"user_ip"`                     // binding:"required" валидирует наличие полей
	RefreshToken string `json:"refresh_token"`               // являются частью библиотеки гин
	TimeLifeRT   int64  `json:"time_life_rt"`
}
