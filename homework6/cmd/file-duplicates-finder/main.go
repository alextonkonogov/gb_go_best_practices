// Для одного или нескольких своих Go-репозиториев реализуйте подходы, изученные в этом уроке:
// Добавьте Makefile или Taskfile с часто используемыми командами по запуску линтеров, тестов и сборке приложения.
// Проверьте работу заданной конфигурации, вызвав make или go-task.
// Установите утилиту pre-commit, добавьте конфигурацию хуков и выполните необходимые действия по их установке.
// Задайте одну или несколько конфигураций для Github Actions и проверьте их работу, запушив изменения в коде.
// Для проверки домашнего задания приложите ссылки на созданные файлы в вашем Github-репозитории.

package main

import (
	"github.com/alextonkonogov/gb_go_best_practices/homework6/internal/app"
	"github.com/alextonkonogov/gb_go_best_practices/homework6/internal/pkg/config"
	"github.com/alextonkonogov/gb_go_best_practices/homework6/internal/pkg/files"
	"github.com/alextonkonogov/gb_go_best_practices/homework6/internal/pkg/log"
)

func main() {
	log := log.NewLogWithConfuguration()
	log.Info("Service started")
	defer log.Info("Service finished")

	log.Info("Config initialization started")
	cnfg, err := config.NewAppConfig()
	if err != nil {
		log.WithError(err).Warn("Invalid config set. Process was stopped")
		return
	}
	log.Info("Config initialization completed")

	uniqueFiles := files.NewUniqueFilesMap(log)

	program := app.NewService(
		cnfg,
		uniqueFiles,
		uniqueFiles,
		uniqueFiles,
		uniqueFiles,
		log,
	)

	err = program.Start()
	if err != nil {
		log.WithError(err).Fatal("Failed to process")
	}
}
