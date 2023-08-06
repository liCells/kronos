package extension

import (
	"github.com/gin-gonic/gin"
	"github.com/liCells/controller/api/common/response"
	"github.com/liCells/controller/global"
	"os/exec"
	"strings"
)

type ExtensionApi struct{}

func (r *ExtensionApi) Execute(c *gin.Context) {
	if name, ok := c.GetQuery("name"); ok {
		for _, script := range global.Config.Scripts {
			if script.Name == name {
				Execute(script.RelativePath, script.ManualExecutionParams)
				response.Success(c)
				return
			}
		}
		response.FailWithMessage(c, "script not found")
		return
	}
	response.BadRequest(c)
}

func Execute(path string, params string) {
	cmd := exec.Command(path, strings.Split(params, " ")...)

	err := cmd.Start()
	if err != nil {
		panic(path + params)
	}
}
