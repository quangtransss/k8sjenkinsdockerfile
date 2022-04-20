package intergration

import (
	"context"
	"fmt"
	"golang/internal/model/domain"
	"golang/utils"
	"testing"

	"github.com/stretchr/testify/require"
)


func TestCreateOrderDetail(t *testing.T)  {
	order ,err :=  RamdomCreateOrders()
	if err != nil {
		fmt.Errorf("cannot create random order: %w",err)
	}

	arg := domain.CreateOrderDetailParams{
		OrderID: order.ID,
		ProductID: 1 ,
		Active: "active" ,
		Total: utils.RandomInt(4,8),
	}

	orderDetail , err := OpenConnection.CreateOrderDetail(context.Background(),arg)
	require.NoError(t,err)
	require.NotEmpty(t,orderDetail)
}