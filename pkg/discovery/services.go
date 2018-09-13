package discovery

import (
	"strconv"

	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/selector"
	"github.com/mtbarta/monocorpus/pkg/logging"
)

type ServiceLocation struct {
	Address string
	Port    string
}

// GetMicroService reuses the go-micro registry to fetch an address and port
// of a registered service.
func GetMicroService(registry registry.Registry, service string) (ServiceLocation, error) {
	selector := selector.NewSelector(
		selector.Registry(registry),
	)

	nextNode, err := selector.Select(service)
	if err != nil {
		logging.Logger.Fatalf("failed to select service", "error", err.Error())
		return ServiceLocation{}, err
	}

	node, err := nextNode()
	if err != nil {
		logging.Logger.Fatalf("failed to find next node", "error", err.Error())
		return ServiceLocation{}, err
	}

	return ServiceLocation{
		Address: node.Address,
		Port:    strconv.Itoa(node.Port),
	}, nil
}
