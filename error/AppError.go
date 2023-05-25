package error

import "fmt"



type AppError struct{
	Module string
	Message string
	Status bool
}

func (error *AppError) Error(module string,message string) AppError{
	error.Module = module
	error.Message = message
	error.Status = true
	return *error	
}

func (error *AppError) IsError() bool {
	return error.Status
}

func (error *AppError) Print(){
	fmt.Print("Модуль: ",error.Module,", ")
	fmt.Println(error.Message)
}

