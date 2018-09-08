package discovery

import (
	"fmt"
	"strconv"

	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/selector"
)

type ServiceLocation struct {
	Address string
	Port    string
}

func GetMicroService(registry registry.Registry, service string) (ServiceLocation, error) {
	selector := selector.NewSelector(
		selector.Registry(registry),
	)

	nextNode, err := selector.Select(service)
	if err != nil {
		fmt.Print(err)
		return ServiceLocation{}, err
	}

	node, err := nextNode()
	if err != nil {
		fmt.Print(err)
		return ServiceLocation{}, err
	}

	return ServiceLocation{
		Address: node.Address,
		Port:    strconv.Itoa(node.Port),
	}, nil
}
