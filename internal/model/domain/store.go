package domain

import (
	"context"
	"database/sql"
	log "github.com/sirupsen/logrus"
)

// Queries to db and excute query
type SQLstore struct {
	*Queries
	db *sql.DB
}


// provide funtion querier
type Store interface {
	Querier
	RoleIdBiz(ctx context.Context,id int64) (Order,error)
}

//Create new instance of SqlStore
func NewSQLstore(db *sql.DB) Store {
	return &SQLstore{
		db:      db,
		Queries: New(db),
	}
}


/// apply design patterns proxy

func (s *SQLstore) RoleIdBiz(ctx context.Context,id int64) (Order,error)  {
	a := NewSQLstore(s.db)
	order , err := a.GetOrdersById(ctx,id)
	if err != nil {
		log.Infoln("order not exist  ")
		return order,err
	}
	log.WithFields(log.Fields{
		"orderID": order.ID,
		"createAt":   order.CreatedAt,
	  }).Info("order info ")
	return order,nil
}

