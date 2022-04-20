package intergration

import (
	"context"
	"golang/internal/model/domain"
	"log"
	"testing"

	"github.com/stretchr/testify/require"
)

func RamdomCreateOrders() (domain.Order,error)  {
	user := CreateRandomUser()
	orders , err := OpenConnection.CreateOrders(context.Background(), user.ID)
	if err != nil {
		log.Fatal("no orders to create ")
	}

	return orders,err
}

func TestCreateRandomOrders(t *testing.T)  {
	order ,err  := RamdomCreateOrders()
	require.NoError(t,err)
	require.NotEmpty(t,order)
}

func TestGetOrderById(t *testing.T)  {
	arg ,err  := RamdomCreateOrders()
	if err != nil {
		log.Fatal("err when create order")
	}
	order , err := OpenConnection.GetOrdersById(context.Background(),arg.ID)
	require.NoError(t ,err)
	require.NotEmpty(t,order)
}

func Tes()  {
	
}


