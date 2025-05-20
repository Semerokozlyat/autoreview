package server

import (
	fasthttprouter "github.com/fasthttp/router"
	"github.com/valyala/fasthttp"

	"github.com/Semerokozlyat/autoreview/internal/server/config"
	"github.com/Semerokozlyat/autoreview/internal/server/handlers"
)

type Server struct {
	*fasthttp.Server
}

func NewServer(cfg *config.Config) *Server {
	r := initRouter(cfg)
	server := &Server{
		Server: &fasthttp.Server{
			Handler: r.Handler,
			Name:    "",
		},
	}
	return server
}

func initRouter(cfg *config.Config) *fasthttprouter.Router {
	r := fasthttprouter.New()
	rootHandler := r.Handler
	if cfg.HTTPServer.EnableCompression {
		rootHandler = fasthttp.CompressHandler(rootHandler)
	}

	r.GET("/", handlers.IndexHandler)
	r.GET("/css/output.css", handlers.NewCSSFSHandler())

	r.GET("/company/add", handlers.CompanyAddHandler)
	r.POST("/company", handlers.CompanyCreateHandler)
	r.DELETE("/company/{id}", handlers.CompanyDeleteHandler)

	r.GET("/custom", handlers.NewCustomHandler("message1").HandleFastHTTP)

	return r
}
