package infra

import (
	"fmt"
	infrahttp "poc-ddd/application/infra/http"
)

const (
	INFRA_Http = "http"
)

type Infra interface {
	Run()
}

type InfraFactory struct {
}

func NewInfraFactory() *InfraFactory {
	return &InfraFactory{}
}

func (i *InfraFactory) CreateInfra(infraType string, port string) (Infra, error) {
	switch infraType {
	case INFRA_Http:
		return infrahttp.NewRouter(port), nil
	default:
		return nil, fmt.Errorf("infra with type %s doesnt exists", infraType)
	}
}
