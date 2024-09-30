package app

type User struct {
	ID           int    `json:"-" db:"id"`
	Guid         int    `json:"guid" db:"guid"`
	UserEmail    string `json:"email" db:"email"` // binding:"required" валидирует наличие полей
	UserIP       string `json:"ip" db:"ip"`       // binding:"required" валидирует наличие полей
	RefreshToken string `json:"token" db:"token"` // являются частью библиотеки гин
	TimeLifeRT   int64  `json:"time" db:"time"`
}
