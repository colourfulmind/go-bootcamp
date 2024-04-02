package models

type User struct {
	//jwt.RegisteredClaims
	ID      int64
	Email   string
	PassHas []byte
	IsAdmin bool
}
