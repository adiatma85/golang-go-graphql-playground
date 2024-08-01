//go:generate go run github.com/99designs/gqlgen generate
package graph

import (
	"github.com/adiatma85/exp-golang-graphql/src/business/usecase"
	"github.com/adiatma85/own-go-sdk/log"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	Uc  *usecase.Usecase
	Log log.Interface
}
