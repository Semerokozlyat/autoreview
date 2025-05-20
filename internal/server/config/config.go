package config

type Config struct {
	HTTPServer HTTPServer `json:"http_server" yaml:"httpServer"`
}

type HTTPServer struct {
	Address           string `json:"address" yaml:"address"`
	EnableCompression bool   `json:"enable_compression" yaml:"enableCompression"`
}
