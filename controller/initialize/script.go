package initialize

import (
	"context"
	"github.com/liCells/controller/api/extension"
	"github.com/liCells/controller/global"
	"github.com/robfig/cron"
)

func LoadScheduledTask() {
	extensions := global.Config.Scripts
	for _, ext := range extensions {
		if !checkForActivation(ext.Name) {
			continue
		}

		callScript(ext)
	}

	extensions = global.Config.Services
	for _, ext := range extensions {
		if !checkForActivation(ext.Name) {
			continue
		}

		startupService(ext)
	}
}

func startupService(ext global.Extension) {
	initEsIndexSetting(ext.EsIndexSetting)
	extension.Execute(ext.RelativePath, ext.ManualExecutionParams)
}

func callScript(ext global.Extension) {
	initEsIndexSetting(ext.EsIndexSetting)
	for _, command := range ext.Commands {
		c := cron.New()
		_ = c.AddFunc(command.Cron, func() {
			extension.Execute(ext.RelativePath, command.Params)
		})
		c.Start()
	}
}

func initEsIndexSetting(indexSetting *global.EsIndexSetting) {
	if indexSetting == nil {
		return
	}
	exists, err := global.ESClient.
		IndexExists(indexSetting.Name).
		Do(context.Background())
	if err != nil {
		panic(err)
	}

	if !exists {
		_, err := global.ESClient.
			CreateIndex(indexSetting.Name).
			BodyString(indexSetting.Setting).
			Do(context.Background())
		if err != nil {
			panic(err)
		}
	}
}

func checkForActivation(extensionName string) bool {
	for _, ext := range global.Config.ActivateExtensions {
		if ext == extensionName {
			return true
		}
	}
	return false
}
