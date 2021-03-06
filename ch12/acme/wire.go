//+build wireinject

package main

import (
	"github.com/PacktPublishing/Hands-On-Dependency-Injection-in-Go/ch12/acme/internal/modules/exchange"
	"github.com/PacktPublishing/Hands-On-Dependency-Injection-in-Go/ch12/acme/internal/modules/get"
	"github.com/PacktPublishing/Hands-On-Dependency-Injection-in-Go/ch12/acme/internal/modules/list"
	"github.com/PacktPublishing/Hands-On-Dependency-Injection-in-Go/ch12/acme/internal/modules/register"
	"github.com/PacktPublishing/Hands-On-Dependency-Injection-in-Go/ch12/acme/internal/rest"
	"github.com/google/wire"
)

// The build tag makes sure the stub is not built in the final build.

func initializeServer() (*rest.Server, error) {
	wire.Build(wireSet)
	return nil, nil
}

func initializeServerCustomConfig(_ exchange.Config, _ get.Config, _ list.Config, _ register.Config, _ rest.Config) *rest.Server {
	wire.Build(wireSetWithoutConfig)
	return nil
}
