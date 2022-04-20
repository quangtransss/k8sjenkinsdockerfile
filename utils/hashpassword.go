package utils

import (
	"fmt"


	"golang.org/x/crypto/bcrypt"
)




const (
	MinCost     int = 4  // the minimum allowable cost as passed in to GenerateFromPassword
	MaxCost     int = 31 // the maximum allowable cost as passed in to GenerateFromPassword
	DefaultCost int = 10 // the cost that will actually be set if a cost below MinCost is passed into GenerateFromPassword
)

func HashedPassword(password string) (string ,error)  {

	/// generate password hash

	cost ,err := bcrypt.Cost([]byte(password))
	if err != nil {
		return "", fmt.Errorf("failed to hash password: %w", err)
	}
	hashPassword ,err := bcrypt.GenerateFromPassword([]byte(password),cost)
	if err != nil {
		return "", fmt.Errorf("failed to hash password: %w", err)
	}

	return string(hashPassword),nil
}


func CheckPasswordAndCompare(password string , hashPassword string) error  {
	err := bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(password))
	return err
}

