package handler

type serviceConfig struct {
	Environment string
}

type Service struct {
	Endpoint string
}

var Services struct {
	Storage Service
}

func init() {
	Services.Storage = Service{"http://localhost:8010"}
}
