// Для выполнения практического задания воспользуйтесь кодом из курсового задания предыдущего модуля
// или кодом web-crawler’а из урока 1.
// Выберете понравившуюся методику структурирования из описанных в разделе про многослойные архитектуры.
// Разбейте код на пакеты в соответствии с выбранной методикой
// linter: commentFormatting: put a space between `//` and comment text (gocritic)
package main

import (
	a "github.com/alextonkonogov/gb_go_best_practices/homework4/internal/app"
	c "github.com/alextonkonogov/gb_go_best_practices/homework4/internal/pkg/config"
	f "github.com/alextonkonogov/gb_go_best_practices/homework4/internal/pkg/files"
	l "github.com/alextonkonogov/gb_go_best_practices/homework4/internal/pkg/log"
)

func main() {
	log := l.NewLogWithConfuguration()
	log.Info("Service started")
	defer log.Info("Service finished")

	log.Info("Config initialization started")
	cnfg, err := c.NewAppConfig()
	if err != nil {
		log.WithError(err).Warn("Invalid config set. Process was stopped")
		return
	}
	log.Info("Config initialization completed")

	uniqueFiles := f.NewUniqueFilesMap(log)

	program := a.NewService(
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
