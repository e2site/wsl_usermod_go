package main

import (
	flag "github.com/spf13/pflag"
	"wsl_usermod_go/error"
)

type Argv struct{
	ConfigFilePath string
	CheckExistFiles bool

}


func (argv *Argv) ParserArgv() error.AppError{
	var err error.AppError;
	var path = flag.String("config","","Path to config json file");
	var check = flag.Bool("checkExistFile",false,"Check rule for exist files")
	flag.Parse()
	if(len(*path)==0) {
		flag.PrintDefaults()
		err.Error("Argv.go","Необходимо указать обязательный параметр")
	}
	argv.ConfigFilePath = *path;
	argv.CheckExistFiles = *check;
	return err;
}