package main

/*
	Главный модуль для запуска персонального сайта prospero78su
*/
import (
	мСвеб "github.com/prospero78/mySite/пакСерверВеб"
	мВрем "time"
)

func main() {
	go мСвеб.ВебСервер.Старт()
	for {
		мВрем.Sleep(мВрем.Second * 10)
	}
}
