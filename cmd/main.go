package main

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server"
)

func main() {

	err := os.Mkdir("logs", 0755)
	// если директория существует, то никак не реагируем
	if err != nil && !errors.Is(err, os.ErrExist) {
		log.Fatal(err)
	}

	flog, err := os.OpenFile("logs/server.logs", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0664)
	if err != nil {
		fmt.Println("Ошибка создания файла для логов: ", err)
		return
	}
	defer flog.Close()

	servLogs := log.New(flog, `serv`, log.LstdFlags|log.Lshortfile)
	servLogs.Println("start serv")

	myServ := server.NewServer(servLogs)

	if err := myServ.Server.ListenAndServe(); err != nil {
		servLogs.Fatal(err)
		return
	}
}
