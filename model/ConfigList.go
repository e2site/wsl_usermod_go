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

func (confgiList *ConfigList) GetConfigRootPath()[]string{
	var rootPath [] string;
	var notSkipedPath [] string;
	for _,v := range *confgiList.Configs {
		if(!v.Skip) {
			notSkipedPath = append(notSkipedPath, v.Path)
		}
	}
	for ind,val := range notSkipedPath {
		var isContainted bool = false;
		for i:=0;i<len(notSkipedPath);i++ {
			if(i==ind) {
				continue
			}
			if( helper.PathContain(notSkipedPath[i],val)) {
				if(len(notSkipedPath[i])<len(val)) {
					isContainted = true
					break
				}
			}
		}
		if(!isContainted) {
			rootPath = append(rootPath, val)
		}
	}
	return rootPath
}  

func (confgiList *ConfigList) IsSkipDirectory(dirPath string) bool{
	config,err:=confgiList.GetConfigByPath(dirPath)
	if(err.IsError() || config==nil) {
		return true
	}
	return config.Skip
}


