package model
import "wsl_usermod_go/error"

type ConfigList struct{
	Configs *[]ConfigModal
}

func (confgiList *ConfigList) GetConfigByPath(path string) (ConfigModal,error.AppError) {
	var config ConfigModal;
	var appError error.AppError;

	return config,appError
}