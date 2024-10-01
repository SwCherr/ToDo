package app

type User struct {
	ID       int    `json:"-" db:"id"`
	Email    string `json:"email" db:"email" binding:"required"`
	Password string `json:"password" db:"password" binding:"required"`
}

type Sesion struct {
	ID           int    `json:"-" db:"id"`
	UserIP       string `json:"-" db:"ip"`
	UserID       int    `json:"id" db:"user_id" binding:"required"`
	GUID         string `json:"guid" db:"guid" binding:"required"` // получать из контекста гина
	RefreshToken string `json:"refreshToken" db:"refresh_token"`
	ExpiresIn    int64  `json:"-" db:"expires_in"`
}
