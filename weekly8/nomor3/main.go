package main

import (
	"errors"
	"fmt"
)

// User struct
type User struct {
	ID   int
	Nama string
}

// UserManager struct
type UserManager struct {
	users map[int]*User
}

// constructor
func NewUserManager() *UserManager {
	return &UserManager{
		users: make(map[int]*User),
	}
}

// AddUser method
func (um *UserManager) AddUser(user *User) error {
	if _, exists := um.users[user.ID]; exists {
		return errors.New("user id sudah terdaftar")
	}

	um.users[user.ID] = user
	return nil
}

// GetUser method
func (um *UserManager) GetUser(id int) (*User, error) {
	user, exists := um.users[id]
	if !exists {
		return nil, errors.New("user tidak ditemukan")
	}
	return user, nil
}

func main() {
	// inisialisasi sistem
	userManager := NewUserManager()

	// tambah minimal 3 user
	users := []*User{
		{ID: 1, Nama: "Andi"},
		{ID: 2, Nama: "Budi"},
		{ID: 3, Nama: "Citra"},
	}

	// fmt.Println(users)
	for _, u := range users {
		fmt.Println(u.ID, u.Nama)
	}


	for _, u := range users {
		err := userManager.AddUser(u)
		if err != nil {
			fmt.Println("Error:", err)
		}
	}

	// coba tambah user dengan ID yang sama
	err := userManager.AddUser(&User{ID: 1, Nama: "Duplicate"})
	if err != nil {
		fmt.Println("Error:", err)
	}

	// ambil user yang ada
	user, err := userManager.GetUser(2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("User ditemukan:", user)
	}

	// ambil user yang tidak ada
	user, err = userManager.GetUser(99)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("User ditemukan:", user)
	}
}

