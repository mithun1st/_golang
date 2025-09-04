package config

import (
	"gin-basic/pkg/enum"
)

var enviromentE enum.EnviromentE

func SetEnviroment(e enum.EnviromentE) {
	enviromentE = e
}

func CuerrentEnviroment() enum.EnviromentE {
	return enviromentE
}
