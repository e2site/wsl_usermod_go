package main

import (
	"fmt"
	"os"
	"wsl_usermod_go/config"
	"wsl_usermod_go/contract"
	"wsl_usermod_go/error"
)

func main() {
	var a Argv;
	err := a.ParserArgv();
	if(err.IsError()) {
		appError(err)
	}
	var config = config.JsonConfig{
		Path: a.ConfigFilePath,
	}
	start(&config)
}

func start(config contract.ConfigContract) {
	configList,err:=config.Parser();
	if(err.IsError()) {
		appError(err)
	}
	fmt.Println(configList)

}

func appError(error error.AppError) {
	error.Print()
	os.Exit(1)
}
