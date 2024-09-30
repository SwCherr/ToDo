package app

type User struct {
	ID           int    `json:"-" db:"id"`
	Guid         int    `json:"guid"`
	UserIP       string `json:"ip"`    // binding:"required" валидирует наличие полей
	RefreshToken string `json:"token"` // являются частью библиотеки гин
	TimeLifeRT   int64  `json:"time"`
}
