package router

import (
    "github.com/liCells/controller/router/data"
    "github.com/liCells/controller/router/extension"
    "github.com/liCells/controller/router/mapping"
    "github.com/liCells/controller/router/registry"
    "github.com/liCells/controller/router/variables"
)

type RouterGroup struct {
    Registry  registry.RouterGroup
    Variables variables.RouterGroup
    Data      data.RouterGroup
    Mapping   mapping.RouterGroup
    Extension extension.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
