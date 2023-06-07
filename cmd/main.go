package main

import (
	"os"
	"wsl_usermod_go/config"
	"wsl_usermod_go/contract"
	"wsl_usermod_go/error"
	"wsl_usermod_go/service"
)

func main() {
	var argv Argv;
	err := argv.ParserArgv();
	appError(err)
	var config = config.JsonConfig{
		Path: argv.ConfigFilePath,
	}
	start(&config,argv.CheckExistFiles)
}

func start(config contract.ConfigContract,checkExistFile bool) {
	configList,err:=config.Parser();
	appError(err)
	
	if(len(*configList.Configs)==0){
		err := error.AppError{}
		err.Error("main.go","Нет загруженных конфигурация для работы")
		appError(err)			
	}
	var watcher service.Watcher;
	var ruleService service.Rule;
	errWatch := watcher.Watch(&configList,ruleService,checkExistFile);
	appError(errWatch)
}

func appError(error error.AppError) {
	if(error.IsError()) {
		error.Print()
		os.Exit(1)
	}
}
