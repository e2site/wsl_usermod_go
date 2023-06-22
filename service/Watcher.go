package service

import (
	"os"
	"path/filepath"
	"wsl_usermod_go/contract"
	wslerror "wsl_usermod_go/error"
	"wsl_usermod_go/model"

	fsnotify "github.com/fsnotify/fsnotify"
)

type Watcher struct{
	watch *fsnotify.Watcher
}

func (watcher *Watcher) getWatch() (*fsnotify.Watcher,wslerror.AppError) {
	var appError wslerror.AppError
	if(watcher.watch!=nil) {
		return watcher.watch,appError
	}
	watch,err := fsnotify.NewWatcher()
	if err != nil {
		appError.Error("Watcher.go",err.Error())
		return watch,appError
	}
	watcher.watch = watch
	return watcher.watch,appError
}

func (watcher *Watcher) Close(){
	if(watcher.watch!=nil) {
		watcher.watch.Close()
	}
}

func (watcher *Watcher) Watch(configList *model.ConfigList,  ruleService contract.RuleContract,checkExistFile bool) wslerror.AppError{

	watch, err := watcher.getWatch()
	if( err.IsError() ) {
		return err
	}

	defer watcher.Close()

	go func() {
		for {
			select {
			case event, ok := <-watch.Events:
				if !ok {
					return
				}
				// Обработка только создания новых файлов
				if event.Op&fsnotify.Create == fsnotify.Create {
					config,pErr := configList.GetConfigByPath(event.Name);
					if(pErr.IsError()) {
						pErr.Print();
					}
					if(config!=nil) {
						if(!config.Skip) {
							isDir,dErr := watcher.IsDir(event.Name);
							if(dErr.IsError()) {
								dErr.Print();
							} else {
								ruleService.SetRule(event.Name,*config,isDir)
								if(isDir) {
									errAdd := watcher.AddWatch(event.Name)
									if(errAdd.IsError()) {
										errAdd.Print()
									} 
								}
							}	
						}
					}
				}
				if(event.Op&fsnotify.Remove == fsnotify.Remove) {
					watcher.DelWatch(event.Name)
				}
			case err, ok := <-watch.Errors:
				if !ok {
					return
				}
				var nErr = wslerror.AppError{}
				nErr.Error("Watcher.go", err.Error())
				nErr.Print()
			}
		}
	}()
	watcher.init(configList,ruleService,checkExistFile);		
	
	select{}
}

func (watcher *Watcher) init(configList *model.ConfigList,  ruleService contract.RuleContract,checkExistFile bool) wslerror.AppError{
	var funcErr  wslerror.AppError

	for _,pathItem :=  range configList.GetConfigRootPath() {
		err := filepath.Walk(pathItem, func(path string, info os.FileInfo, err error) error  {
			if err != nil {
				scanError := wslerror.AppError{}
				scanError.Error("Watch.go",err.Error())
				scanError.Print()
				return nil
			}
	
			config,errCfg := configList.GetConfigByPath(path)
			if(errCfg.IsError()) {
				errCfg.Print()
				return nil
			}

			if(checkExistFile) {
				rulErr := ruleService.SetRule(path,*config,info.IsDir())
				if(rulErr.IsError()) {
					rulErr.Print()
				}
			}

			if info.IsDir() {
				watcher.AddWatch(path)
			}
	
			return nil
		})
	
		if err != nil {
			funcErr.Error("Watcher.go",err.Error())
			return funcErr
		}
	}


	//watcher.AddWatch("/home/filipp")
	return funcErr
}

func (watcher *Watcher) AddWatch(path string) wslerror.AppError {
	var addErr = wslerror.AppError{};
	watch,_ := watcher.getWatch()
	err := watch.Add(path)
	if err != nil {
		addErr.Error("Watch.go",err.Error())
	}
	return addErr
}

func (watcher *Watcher) DelWatch(path string) {
	watch,_ := watcher.getWatch()
	for _,pathWatch := range watch.WatchList() {
		if(pathWatch == path) {
			watch.Remove(path)
		}
	}
}



func (watcher *Watcher) IsDir(path string) (bool,wslerror.AppError){
	var errFunc = wslerror.AppError{}
	var isDir bool
	info, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			errFunc.Error("Watch.go IsDir","Путь не существует")
		} else {
			errFunc.Error("Watch.go IsDir",err.Error())
		}
	} else {
		isDir = info.IsDir()
	}
	return isDir,errFunc
}