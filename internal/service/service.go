package service

import (
	"context"
	"database/sql"
	"golang/internal/model/domain"
	log "github.com/sirupsen/logrus"
	_ "github.com/lib/pq"

)


type ServiceStore struct {
	*domain.Queries
	db *sql.DB
}
type Service interface {
	domain.Querier
	OrderByIdBiz(ctx context.Context,id int64) (domain.Order,error)
	GetRoleByIdBiz(ctx context.Context,id int64) (domain.Role,error)
}
func NewServiceStore(db *sql.DB) Service {
	return &ServiceStore{
		db : db,
		Queries:  domain.New(db),
	}
}
func (s *ServiceStore) OrderByIdBiz(ctx context.Context,id int64) (domain.Order,error)  {
	a := domain.NewSQLstore(s.db)
	order , err := a.GetOrdersById(ctx,id)
	if err != nil {
		log.Warn("order not exist  ")
		return order,err
	}

	log.WithFields(log.Fields{
		"orderID": order.ID,
		"createAt":   order.CreatedAt,
	  }).Info("order info ")
	return order,nil
}

func (s *ServiceStore) GetRoleByIdBiz(ctx context.Context,id int64) (domain.Role,error)  {
	a := NewServiceStore(s.db)
	role , err := a.GetRole(ctx,id)
	if err != nil {
		log.Warn("order not exist  ")
		return role,err
	}
	log.WithFields(log.Fields{
		"orderID": role.ID,
		"createAt":   role.CreatedAt,
	  }).Info("role info ")
	return role,nil
}