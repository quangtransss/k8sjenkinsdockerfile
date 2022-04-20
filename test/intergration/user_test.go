package intergration

import (
	"context"
	"fmt"

	"golang/internal/model/domain"
	"golang/utils"
	"testing"

	"github.com/stretchr/testify/require"
)

func CreateRandomUser() domain.User {
	role := GetRandomRole()
	arg := domain.CreateUserParams{
		FullName:       utils.RandomString(10),
		Username:       utils.RandomString(10),
		HashedPassword: utils.RandomString(10),
		Email:          utils.RandomEmail(8),
		Mobile:         utils.RandomInt(8, 16),
		Roleid:         role.ID,
	}

	user,err := OpenConnection.CreateUser(context.Background(),arg)
	if err != nil {
		fmt.Errorf("Cannot create user :%w" ,err)
	}

	return user
}
func TestCreateUser(t *testing.T)  {
	for i := 0; i < 10; i++ {
	user := CreateRandomUser()
	require.NotEmpty(t,user)
	}
}


func TestListUser(t *testing.T)  {
	var user domain.User
	for i := 0; i < 10; i++ {
		user = CreateRandomUser()
	}

	arg := domain.ListUsersParams{
		Username: user.Username,
		Limit: 4,
		Offset: 0,
	}

	users ,err := OpenConnection.ListUsers(context.Background(),arg)




	require.NoError(t,err)
	require.NotEmpty(t,users)

}