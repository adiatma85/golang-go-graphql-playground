package handler

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/adiatma85/exp-golang-graphql/graph"
	"github.com/adiatma85/exp-golang-graphql/src/business/usecase"
	"github.com/adiatma85/exp-golang-graphql/utils/config"
	"github.com/adiatma85/own-go-sdk/instrument"
	"github.com/adiatma85/own-go-sdk/jwtAuth"
	"github.com/adiatma85/own-go-sdk/log"
	"github.com/adiatma85/own-go-sdk/parser"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

const (
	infoRequest  string = `httpclient Sent Request: uri=%v method=%v`
	infoResponse string = `httpclient Received Response: uri=%v method=%v resp_code=%v`
)

var once = &sync.Once{}

type REST interface {
	// Run()
}

type rest struct {
	http       *gin.Engine
	conf       config.GinConfig
	json       parser.JSONInterface
	log        log.Interface
	uc         *usecase.Usecase
	instrument instrument.Interface
	jwtAuth    jwtAuth.Interface
}

type InitParam struct {
	Http       *gin.Engine
	Conf       config.GinConfig
	Json       parser.JSONInterface
	Log        log.Interface
	Uc         *usecase.Usecase
	Instrument instrument.Interface
	JwtAuth    jwtAuth.Interface
}

func Init(param InitParam) REST {
	r := &rest{}

	once.Do(func() {
		switch param.Conf.Mode {
		case gin.ReleaseMode:
			gin.SetMode(gin.ReleaseMode)
		case gin.DebugMode, gin.TestMode:
			gin.SetMode(gin.TestMode)
		default:
			gin.SetMode("")
		}

		httpServer := gin.New()

		r = &rest{
			conf:       param.Conf,
			log:        param.Log,
			json:       param.Json,
			http:       httpServer,
			uc:         param.Uc,
			instrument: param.Instrument,
			jwtAuth:    param.JwtAuth,
		}

		// Set CORS
		switch r.conf.CORS.Mode {
		case "allowall":
			r.http.Use(cors.New(cors.Config{
				AllowAllOrigins: true,
				AllowHeaders:    []string{"*"},
				AllowMethods: []string{
					http.MethodHead,
					http.MethodGet,
					http.MethodPost,
					http.MethodPut,
					http.MethodPatch,
					http.MethodDelete,
				},
			}))
		default:
			r.http.Use(cors.New(cors.DefaultConfig()))
		}

		// Set Timeout
		// r.http.Use(r.SetTimeout)

		// Set Recovery
		// r.http.Use(r.CustomRecovery)

		r.Register()
	})

	return r
}

func (r *rest) Register() {
	// server health and testing purpose
	r.http.GET("/ping", r.Ping)

	// Server Graphql
	server := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))

	fmt.Println("Server adalah: ", server)
}
