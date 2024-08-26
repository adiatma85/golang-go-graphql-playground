package usecase

import (
	"github.com/adiatma85/exp-golang-graphql/src/business/domain"
	"github.com/adiatma85/exp-golang-graphql/src/business/usecase/user"
	"github.com/adiatma85/own-go-sdk/jwtAuth"
	"github.com/adiatma85/own-go-sdk/log"
)

type Usecase struct {
	User user.Interface
}

type InitParam struct {
	Log     log.Interface
	Dom     *domain.Domain
	JwtAuth jwtAuth.Interface
}

func Init(param InitParam) *Usecase {
	usecase := &Usecase{
		User: user.Init(user.InitParam{Log: param.Log, User: param.Dom.User, JwtAuth: param.JwtAuth}),
		// Category: category.Init(category.InitParam{Log: param.Log, Category: param.Dom.Category, JwtAuth: param.JwtAuth}),
		// Task:     task.Init(task.InitParam{Log: param.Log, Task: param.Dom.Task, JwtAuth: param.JwtAuth}),
		// Role:     role.Init(role.InitParam{Log: param.Log, Role: param.Dom.Role, JwtAuth: param.JwtAuth}),
	}

	return usecase
}
