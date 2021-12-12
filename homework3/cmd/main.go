//Доработать тесты в программе по поиску дубликатов файлов
//1 Рефакторим код в соответствие с рассмотренными принципами (SRP, чистые функции, интерфейсы, убираем глобальный стэйт)
//2 Пробуем использовать testify
//3 Делаем стаб/мок (например, для файловой системы) и тестируем свой код без обращений к внешним компонентам (файловой системе)
//4 Делаем отдельно 1-2 интеграционных теста, запускаемых с флагом -integration

package main

import (
	"github.com/sirupsen/logrus"

	"github.com/alextonkonogov/gb_go_best_practices/homework3/pkg/config"
	f "github.com/alextonkonogov/gb_go_best_practices/homework3/pkg/files"
	p "github.com/alextonkonogov/gb_go_best_practices/homework3/pkg/program"
)

func main() {
	log := logrus.New()
	customFormatter := new(logrus.TextFormatter)
	customFormatter.TimestampFormat = "2006-01-02 15:04:05"
	customFormatter.FullTimestamp = true
	log.SetLevel(logrus.DebugLevel)
	log.SetFormatter(customFormatter)

	log.Info("Service started")
	defer log.Info("Service finished")

	log.Info("Config initialization started")
	cnfg, err := config.NewAppConfig()
	if err != nil {
		log.WithError(err).Warn("Invalid config set. Process was stopped")
		return
	}
	log.Info("Config initialization completed")

	uniqueFiles := f.NewUniqueFilesMap(log)

	program := p.NewProgram(
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
