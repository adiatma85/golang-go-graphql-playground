package main

import (
	"fmt"

	"github.com/adiatma85/exp-golang-graphql/src/business/domain"
	"github.com/adiatma85/exp-golang-graphql/src/business/usecase"
	"github.com/adiatma85/exp-golang-graphql/utils/config"
	"github.com/adiatma85/own-go-sdk/configreader"
	"github.com/adiatma85/own-go-sdk/instrument"
	"github.com/adiatma85/own-go-sdk/jwtAuth"
	"github.com/adiatma85/own-go-sdk/log"
	"github.com/adiatma85/own-go-sdk/parser"
	"github.com/adiatma85/own-go-sdk/sql"
)

const defaultPort = "8080"

const (
	configfile   string = "./etc/cfg/conf.json"
	templatefile string = "./etc/tpl/conf.template.json"
)

func main() {
	// KODINGAN YANG LAMA

	// // Using chi router
	// router := chi.NewRouter()

	// // Kalau kaya gini berarti semua router kenak dong??
	// // Jawaban iya, dia kenak semuanya
	// router.Use(auth.Middleware())

	// mysql.InitDB()
	// defer mysql.CloseDB()
	// server := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))
	// router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	// router.Handle("/query", server)

	// log.Printf("connect to http://localhost:%s/ for GraphQL playground", cfg.Gin.Port)
	// log.Fatal(http.ListenAndServe(":"+cfg.Gin.Port, router))

	// KODINGAN YANG LAMA

	// Build the config
	// Assume the config is exist in the first place

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

	// Init the GIN
	fmt.Println("uc adalah: ", uc)
}
