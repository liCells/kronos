package global

import (
	"github.com/olivere/elastic/v7"
)

var (
	Config Configuration
	// ESClient is the global ES client
	ESClient *elastic.Client
)

type Configuration struct {
	Port               uint        `json:"port"`
	Register           control     `json:"register"`
	Es                 es          `json:"es"`
	ActivateExtensions []string    `json:"activate_extensions"`
	Services           []Extension `json:"services"`
	Scripts            []Extension `json:"scripts"`
}

type control struct {
	Renew                                uint   `json:"renew"`
	AllowedMaximumNumberOfDisconnections string `json:"allowed_maximum_number_of_disconnections"`
}

type es struct {
	Host     string `json:"host"`
	Port     uint   `json:"port"`
	Analyzer string `json:"analyzer"`
}

type Extension struct {
	Name                  string         `json:"name"`
	Description           string         `json:"description"`
	Version               string         `json:"version"`
	Author                string         `json:"author"`
	RelativePath          string         `json:"relative_path"`
	InterfaceType         string         `json:"interface_type"`
	EsIndexSetting        EsIndexSetting `json:"es_index_setting"`
	Commands              []Command      `json:"commands"`
	ManualExecutionParams string         `json:"manual_execution_params"`
}

type EsIndexSetting struct {
	Name    string  `json:"name"`
	Setting string  `json:"setting"`
	Mapping Mapping `json:"mapping"`
}

type Mapping struct {
	Id      string `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Source  string `json:"source"`
	Tag     string `json:"tag"`
}

type Command struct {
	Params string `json:"params"`
	Cron   string `json:"cron"`
}
