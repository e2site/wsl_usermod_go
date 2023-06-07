package contract

import(
	"wsl_usermod_go/model"
	"wsl_usermod_go/error"
)
package contract

import(
	"wsl_usermod_go/model"
	"wsl_usermod_go/error"
)

type ConfigContract interface{
	Parser() (model.ConfigList, error.AppError)
}
type ConfigContract interface{
	Parser() (model.ConfigList, error.AppError)
}