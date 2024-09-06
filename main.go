package main

import (
	// graph_handler "github.com/99designs/gqlgen/graphql/handler"

	"github.com/adiatma85/exp-golang-graphql/src/business/domain"
	"github.com/adiatma85/exp-golang-graphql/src/business/graphql"
	"github.com/adiatma85/exp-golang-graphql/src/business/handler"
	"github.com/adiatma85/exp-golang-graphql/src/business/usecase"
	"github.com/adiatma85/exp-golang-graphql/utils/config"
	"github.com/adiatma85/own-go-sdk/configreader"
	"github.com/adiatma85/own-go-sdk/instrument"
	"github.com/adiatma85/own-go-sdk/jwtAuth"
	"github.com/adiatma85/own-go-sdk/log"
	"github.com/adiatma85/own-go-sdk/parser"
	"github.com/adiatma85/own-go-sdk/sql"
)

const (
	configfile   string = "./etc/cfg/conf.json"
	templatefile string = "./etc/tpl/conf.template.json"
)

func main() {
	// Read the Config first
	cfg := config.Init()
	configreader := configreader.Init(configreader.Options{
		ConfigFile: configfile,
	})
	configreader.ReadConfig(&cfg)

	// init logger
	log := log.Init(cfg.Log)

	// init the instrument
	instr := instrument.Init(cfg.Instrument)

	// Init the DB
	db := sql.Init(cfg.SQL, log, instr)

	// init the parser
	parsers := parser.InitParser(log, cfg.Parser)

	// Init the jwt
	jwt := jwtAuth.Init(cfg.JwtAuth)

	// Init the domain
	d := domain.Init(domain.InitParam{Log: log, Db: db, Json: parsers.JSONParser()})

	// Init the usecase
	uc := usecase.Init(usecase.InitParam{Log: log, Dom: d, JwtAuth: jwt})

	// Initialize the Graphql in here
	graphql := graphql.NewExecutableSchema(graphql.Config{Resolvers: &graphql.Resolver{Uc: uc, Log: log}})

	// Init the GIN
	rest := handler.Init(handler.InitParam{Conf: cfg.Gin, Json: parsers.JSONParser(), Log: log, Uc: uc, Instrument: instr, JwtAuth: jwt, Graphql: graphql})

	rest.Run()
}
