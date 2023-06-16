package service

import "time"

type CreateClientRequest struct {
    FirstName string `json:"first_name"`
    LastName string `json:"last_name"`
    Email string `json:"email"`
    Password string `json:"password"`
    PasswordVerify string `json:"password_verify"`
}

type DeleteClientRequest struct {
    Id int `json:"id"`
}

type GetClientRequest struct {
    Email string `json:"email"`
    Password string `json:"password"`
}

type Client struct {
    Id int `json:"id"`
    FirstName string `json:"first_name"`
    LastName string `json:"last_name"`
    Email string `json:"email"`
    Password string 
    CreatedAt time.Time `json:"created_at"`
}

func NewClient(first_name, last_name, email, password string) *Client {
    return &Client{
        FirstName: first_name,
        LastName: last_name,
        Email: email,
        Password: password,
        CreatedAt: time.Now().UTC(),
    }
}
