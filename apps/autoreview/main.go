package main

import (
	"flag"
	"log"

	"github.com/Semerokozlyat/autoreview/internal/server"
	"github.com/Semerokozlyat/autoreview/internal/server/config"
)

var (
	addr     = flag.String("addr", ":3000", "Server TCP address")
	compress = flag.Bool("compress", false, "Enable transparent response compression")
)

func main() {
	flag.Parse()

	cfg := &config.Config{
		HTTPServer: config.HTTPServer{
			Address:           *addr,
			EnableCompression: *compress,
		},
	}

	srv := server.NewServer(cfg)

	err := srv.ListenAndServe(cfg.HTTPServer.Address)
	if err != nil {
		log.Fatal(err)
	}
}
