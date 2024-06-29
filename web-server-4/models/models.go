package models

import "time"

type UserStore interface {
	GetUserByEmail(email string) (*User, error)
	GetUserById(id int) (*User, error)
	CreateUser(User) error
}

type RegisterUserPayload struct {
	FirstName string `json:"first_name" validate:"required"`
	LastName  string `json:"last_name" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required,min=3,max=130"`
}

type User struct {
	ID        int       `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
	CreatedAt time.Time `json:"created_at"`
}

type Products struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Image       string    `json:"image"`
	Quantity    int       `json:"quantity"`
	CreatedAt   time.Time `json:"created_at"`
}

type Orders struct {
	ID        int       `json:"id"`
	UserId    int       `json:"user_id"`
	Total     int       `json:"total"`
	Status    string    `json:"status"`
	Address   string    `json:"address"`
	CreatedAt time.Time `json:"created_at"`
}

type OrdersItems struct {
	ID        int       `json:"id"`
	OrderId   string    `json:"order_id"`
	ProductId string    `json:"product_id"`
	Quantity  int       `json:"quantity"`
	Price     int       `json:"price"`
	CreatedAt time.Time `json:"created_at"`
}
