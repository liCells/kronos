package api

import (
    "github.com/liCells/controller/api/data"
    "github.com/liCells/controller/api/extension"
    "github.com/liCells/controller/api/mapping"
    "github.com/liCells/controller/api/registry"
    "github.com/liCells/controller/api/variables"
)

type Group struct {
    registry.RegistryApi
    variables.VariablesApi
    data.DataApi
    mapping.MappingApi
    extension.ExtensionApi
}

var ApiGroup = new(Group)
