package service

import (
	"github.com/google/wire"
	"xiuyiPro/internal/service/greeter"
	"xiuyiPro/internal/service/idiom"
	"xiuyiPro/internal/service/turtle"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(greeter.NewGreeterService, turtle.NewTurtleService, idiom.NewIdiomService)
