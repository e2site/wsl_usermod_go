package model

import (
	"os"
	"wsl_usermod_go/error"
	"wsl_usermod_go/helper"
)

type ConfigList struct{
	Configs *[]ConfigModal
}

func (confgiList *ConfigList) GetConfigByPath(path string) (*ConfigModal,error.AppError) {
	config := &ConfigModal{};
	var retConfig *ConfigModal;
	
	var appError error.AppError;
	var cntSize int = 99999

	for _,cfg := range *confgiList.Configs {
		if(len(cfg.Path)>len(path)) {
			continue 
		}
		var listCfgPath = helper.Explode(cfg.Path,string(os.PathSeparator))
		var listPath = helper.Explode(path,string(os.PathSeparator))
		
		var isContainted bool = true

		for ind,cPath := range listCfgPath {
			if(!isContainted) {
				break
			}
			if(len(cPath) != len(listPath[ind])) {
				isContainted = false
				break
			}
			for i:=0; i < len(cPath);i++{
				if(cPath[i] != listPath[ind][i]) {
					isContainted = false
					break
				} 
			}
		}
		
	
		var cntTmp = len(listPath) - len(listCfgPath)
		if(isContainted  && cntSize>cntTmp) {
			cntSize = cntTmp
			*config = cfg	
			retConfig = config;
		}
	}
	return retConfig,appError
}

func (confgiList *ConfigList) IsSkipDirectory(dirPath string) bool{
	config,err:=confgiList.GetConfigByPath(dirPath)
	if(err.IsError() || config==nil) {
		return true
	}
	return config.Skip
}


