package domain

import (
	"github.com/adiatma85/exp-golang-graphql/src/business/domain/role"
	"github.com/adiatma85/exp-golang-graphql/src/business/domain/user"
	"github.com/adiatma85/own-go-sdk/log"
	"github.com/adiatma85/own-go-sdk/parser"
	"github.com/adiatma85/own-go-sdk/sql"
)

type Domain struct {
	User user.Interface
	Role role.Interface
}

type InitParam struct {
	Log  log.Interface
	Db   sql.Interface
	Json parser.JSONInterface
}

func Init(param InitParam) *Domain {
	domain := &Domain{
		User: user.Init(user.InitParam{Log: param.Log, Db: param.Db, Json: param.Json}),
		Role: role.Init(role.InitParam{Log: param.Log, Db: param.Db, Json: param.Json}),
	}

	return domain
}
