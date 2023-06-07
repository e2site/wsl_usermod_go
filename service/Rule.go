package service

import (
	"io/fs"
	"os"
	"strconv"
	"wsl_usermod_go/error"
	"wsl_usermod_go/model"
)

type Rule struct{

}

func (rule Rule) SetRule(path string,config model.ConfigModal,isDir bool) error.AppError {
	var err = error.AppError{}
	
	if(config.Skip){
		return err
	}

	errChOwn := os.Chown(path,config.UID,config.GID)
	if(errChOwn!=nil) {
		err.Error("Rule.go",errChOwn.Error())
	}
	var mode string;

	if(isDir) {
		mode = config.DirectoryMod
	} else {
		mode = config.FileMod	
	}

	cm,convErr := strconv.ParseInt(mode, 8, 32)
	if(convErr!= nil) {
		err.Error("Rule.go",convErr.Error())
	} else {
		errChMode := os.Chmod(path,fs.FileMode(cm))
		if(errChMode != nil) {
			err.Error("Rule.go",errChMode.Error())
		}
	}

	return err;
}