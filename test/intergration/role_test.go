package intergration

import (
	"context"
	"database/sql"
	"fmt"
	// "time"

	"math/rand"

	"golang/internal/model/domain"
	"golang/utils"
	"log"
	"testing"

	"github.com/stretchr/testify/require"
)
func RandomIntExceptEqualZero(n int64) int64  {
	id := rand.Int63n(n)

	if id == 0 {
		a := id + 1
		return a
	}
	return id
}

func CreateRandomRole() domain.Role {
	arg := domain.CreateRoleParams{
		Title:       utils.RandomString(8),
		Slug:        utils.RandomString(8),
		Active:      "disable",
		Description: sql.NullString{utils.RandomString(20), true},
	}

	role , err := OpenConnection.CreateRole(context.Background(), arg)
	if err != nil {
		log.Fatal("cannot create role")
	}
	return role
}


func GetRandomRole() domain.Role {
	id := RandomIntExceptEqualZero(4)
	role , err := OpenConnection.GetRole(context.Background(),id)

	if err != nil {
		fmt.Errorf("Cannot exist role %w",err)
	}
	return role
}
func TestGet(t *testing.T)  {
	GetRandomRole()
}
func TestCreateRole(t *testing.T)  {
	for i := 0; i < 10; i++ {
		arg := domain.CreateRoleParams{
			Title: utils.RandomString(8),
			Slug: utils.RandomString(8),
			Active: "active" ,
			Description: sql.NullString{utils.RandomString(20),true || false} ,
		}
		role , err := OpenConnection.CreateRole(context.Background(), arg)
		require.NoError(t, err)
		require.NotEmpty(t,role)
		// time.Sleep(time.Duration(time.Second*1))
	}

}


func TestGetRole(t *testing.T)  {
	id := RandomIntExceptEqualZero(4)
	role , err := OpenConnection.GetRole(context.Background(),id)

	require.NoError(t,err)
	require.NotEmpty(t,role)

}


