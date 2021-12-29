// +build unit

package app_test

import (
	"fmt"
	"testing"

	"github.com/sirupsen/logrus"

	"github.com/alextonkonogov/gb_go_best_practices/homework6/internal/app"
	"github.com/alextonkonogov/gb_go_best_practices/homework6/internal/pkg/config"
	"github.com/alextonkonogov/gb_go_best_practices/homework6/internal/pkg/files"
)

type Stub struct {
	duplicates int
}

func (s Stub) GetDuplicatesCount() int {
	return s.duplicates
}

func (s Stub) Find(searchPath string, workers int) int {
	s.duplicates = 2
	return s.duplicates
}

func (s Stub) Print(searchPath string) {
	logrus.Info(fmt.Sprintf("Found %d unique files and %d duplicates in \"%s\":\n", 5, 2, searchPath))
	var stubFiles = []string{
		"1) file one.txt",
		"2) file two.txt",
		"3) file three.txt",
		"4) file four.txt",
		"5) file five.txt",
	}
	for _, file := range stubFiles {
		logrus.Info(file)
	}
}

func (s Stub) DeleteDuplicates() error {
	logrus.Info("Duplicate files deletion started")
	for i := 0; i < s.GetDuplicatesCount(); i++ {
		logrus.Info(fmt.Sprintf("...deleting three %d.txt", i))
	}

	logrus.Info("Duplicate files deletion ended")
	return nil
}

func TestProgramWithStub(t *testing.T) {
	log := logrus.New()
	cnfg, _ := config.NewAppConfig()

	stab := &Stub{}
	program := app.NewService(
		cnfg,
		stab,
		stab,
		stab,
		stab,
		log,
	)

	err := program.Start()
	if err != nil {
		t.Error("Failed to process")
	}
}

func Example() {
	log := logrus.New()
	cnfg, err := config.NewAppConfig()
	if err != nil {
		log.Fatal(err)
	}
	uniqueFiles := files.NewUniqueFilesMap(log)

	program := app.NewService(cnfg, uniqueFiles, uniqueFiles, uniqueFiles, uniqueFiles, log)
	err = program.Start()
	if err != nil {
		log.Fatal(err)
	}
	// Output:
}

func BenchmarkProgram_Start(b *testing.B) {
	log := logrus.New()

	for j := 0; j < b.N; j++ {
		cnfg, err := config.NewAppConfig()
		if err != nil {
			log.Fatal(err)
		}
		cnfg.PrintResult = false
		uniqueFiles := files.NewUniqueFilesMap(log)

		program := app.NewService(cnfg, uniqueFiles, uniqueFiles, uniqueFiles, uniqueFiles, log)
		err = program.Start()
		if err != nil {
			log.Fatal(err)
		}
	}
}

//goos: darwin
//goarch: amd64
//pkg: github.com/alextonkonogov/gb-golang-level-2/homework8/program
//BenchmarkProgram_Start
//BenchmarkProgram_Start-8   	   23029	     49370 ns/op
//PASS
