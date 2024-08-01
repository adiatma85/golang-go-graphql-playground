package handler

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/adiatma85/exp-golang-graphql/src/business/entity"
	"github.com/adiatma85/own-go-sdk/appcontext"
	"github.com/adiatma85/own-go-sdk/codes"
	"github.com/adiatma85/own-go-sdk/errors"
	"github.com/adiatma85/own-go-sdk/header"
	"github.com/gin-gonic/gin"
)

func (r *rest) Ping(ctx *gin.Context) {
	resp := entity.Ping{
		Status:  "OK",
		Version: r.conf.Meta.Version,
	}
	r.httpRespSuccess(ctx, codes.CodeSuccess, resp, nil)
}

// Graphql Handler
func (r *rest) graphqlHandler() gin.HandlerFunc {
	h := handler.NewDefaultServer(r.graphql)

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

// Playground Handler
func (r *rest) playgroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL", "/query")

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func (r *rest) SetTimeout(ctx *gin.Context) {
	// wrap the request context with a timeout
	c, cancel := context.WithTimeout(ctx.Request.Context(), r.conf.Timeout)

	// cancel to clear resources after finished
	defer cancel()

	c = appcontext.SetRequestStartTime(c, time.Now())

	// replace request with context wrapped request
	ctx.Request = ctx.Request.WithContext(c)
	ctx.Next()

}

func (r *rest) httpRespSuccess(ctx *gin.Context, code codes.Code, data interface{}, p *entity.Pagination) {
	successApp := codes.Compile(code, appcontext.GetAcceptLanguage(ctx))
	c := ctx.Request.Context()
	meta := entity.Meta{
		Path:       r.conf.Meta.Host + ctx.Request.URL.String(),
		StatusCode: successApp.StatusCode,
		Status:     http.StatusText(successApp.StatusCode),
		Message:    fmt.Sprintf("%s %s [%d] %s", ctx.Request.Method, ctx.Request.URL.RequestURI(), successApp.StatusCode, http.StatusText(successApp.StatusCode)),
		Timestamp:  time.Now().Format(time.RFC3339),
		RequestID:  appcontext.GetRequestId(c),
	}

	resp := &entity.HTTPResp{
		Message: entity.HTTPMessage{
			Title: successApp.Title,
			Body:  successApp.Body,
		},
		Meta:       meta,
		Data:       data,
		Pagination: p,
	}

	reqstart := appcontext.GetRequestStartTime(c)
	if !time.Time.IsZero(reqstart) {
		resp.Meta.TimeElapsed = fmt.Sprintf("%dms", int64(time.Since(reqstart)/time.Millisecond))
	}

	raw, err := r.json.Marshal(&resp)
	if err != nil {
		r.httpRespError(ctx, errors.NewWithCode(codes.CodeInternalServerError, err.Error()))
		return
	}

	c = appcontext.SetAppResponseCode(c, code)
	c = appcontext.SetResponseHttpCode(c, successApp.StatusCode)
	ctx.Request = ctx.Request.WithContext(c)

	ctx.Header(header.KeyRequestID, appcontext.GetRequestId(c))
	ctx.Data(successApp.StatusCode, header.ContentTypeJSON, raw)
}

func (r *rest) httpRespError(ctx *gin.Context, err error) {
	c := ctx.Request.Context()

	if errors.Is(c.Err(), context.DeadlineExceeded) {
		err = errors.NewWithCode(codes.CodeContextDeadlineExceeded, "Context Deadline Exceeded")
	}

	httpStatus, displayError := errors.Compile(err, appcontext.GetAcceptLanguage(c))
	statusStr := http.StatusText(httpStatus)

	errResp := &entity.HTTPResp{
		Message: entity.HTTPMessage{
			Title: displayError.Title,
			Body:  displayError.Body,
		},
		Meta: entity.Meta{
			Path:       r.conf.Meta.Host + ctx.Request.URL.String(),
			StatusCode: httpStatus,
			Status:     statusStr,
			Message:    fmt.Sprintf("%s %s [%d] %s", ctx.Request.Method, ctx.Request.URL.RequestURI(), httpStatus, statusStr),
			Error: &entity.MetaError{
				Code:    int(displayError.Code),
				Message: err.Error(),
			},
			Timestamp: time.Now().Format(time.RFC3339),
			RequestID: appcontext.GetRequestId(c),
		},
	}

	r.log.Error(c, err)

	c = appcontext.SetAppResponseCode(c, displayError.Code)
	c = appcontext.SetAppErrorMessage(c, fmt.Sprintf("%s - %s", displayError.Title, displayError.Body))
	c = appcontext.SetResponseHttpCode(c, httpStatus)
	ctx.Request = ctx.Request.WithContext(c)

	ctx.Header(header.KeyRequestID, appcontext.GetRequestId(c))
	ctx.AbortWithStatusJSON(httpStatus, errResp)
}
