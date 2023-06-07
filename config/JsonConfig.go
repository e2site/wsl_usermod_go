package config

import (
	"os"
	_ "wsl_usermod_go/contract"
	"wsl_usermod_go/model"
	"wsl_usermod_go/error"
	"encoding/json"
)

type JsonConfig struct{
	Path string
}

func (j *JsonConfig) Parser()  (model.ConfigList, error.AppError) {
	var configs []model.ConfigModal;
	var configList model.ConfigList;
	var moduleError error.AppError;
	data,err := os.ReadFile(j.Path)
	if(err!=nil) {
		moduleError.Error("JsonConfig.go",err.Error())
	} else{
		err := json.Unmarshal(data,&configs)
		if(err!=nil) {
			moduleError.Error("JsonConfig.go",err.Error())
		} else {
			configList.Configs = &configs
		}	
	}
	return configList, moduleError
}