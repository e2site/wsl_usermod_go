package service

import (
	"os"
	"wsl_usermod_go/contract"
	"wsl_usermod_go/error"
	"wsl_usermod_go/model"

	fsnotify "github.com/fsnotify/fsnotify"
)

type Watcher struct{

}

func (watcher *Watcher) Watch(configList *model.ConfigList,  ruleService contract.RuleContract,checkExistFile bool) error.AppError{
	var appError  error.AppError;
	watch, err := fsnotify.NewWatcher()
	if err != nil {
		appError.Error("Watcher.go",err.Error())
		return appError
	}

	defer watch.Close()

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
								pErr.Print();
							} else {
								ruleService.SetRule(event.Name,*config,isDir)
								if(isDir) {
									errAdd := watcher.AddWatch(event.Name,watch)
									if(errAdd.IsError()) {
										errAdd.Print()
									} 
								}
							}	
						}
					}
				}
				if(event.Op&fsnotify.Remove == fsnotify.Remove) {
					watcher.DelWatch(event.Name,watch)
				}
			case err, ok := <-watch.Errors:
				if !ok {
					return
				}
				var nErr = error.AppError{}
				nErr.Error("Watcher.go", err.Error())
				nErr.Print()
			}
		}
	}()
	watcher.AddWatch("/home/filipp",watch)
	select{}
}

func (watcher *Watcher) AddWatch(path string, watch *fsnotify.Watcher) error.AppError {
	var addErr = error.AppError{};
	err := watch.Add(path)
	if err != nil {
		addErr.Error("Watch.go",err.Error())
	}
	return addErr
}

func (watcher *Watcher) DelWatch(path string, wach *fsnotify.Watcher) {
	for _,pathWatch := range wach.WatchList() {
		if(pathWatch == path) {
			wach.Remove(path)
		}
	}
}



func (watcher *Watcher) IsDir(path string) (bool,error.AppError){
	var errFunc = error.AppError{}
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