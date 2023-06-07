package contract


import(
	"wsl_usermod_go/model"
	"wsl_usermod_go/error"
)

type RuleContract interface{
	SetRule(path string,config model.ConfigModal,isDir bool) error.AppError
}