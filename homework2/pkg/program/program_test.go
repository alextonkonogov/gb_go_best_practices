package program_test

import (
	"io/ioutil"
	"testing"

	"github.com/sirupsen/logrus"

	"github.com/alextonkonogov/gb_go_best_practices/homework2/pkg/config"
	"github.com/alextonkonogov/gb_go_best_practices/homework2/pkg/files"
	"github.com/alextonkonogov/gb_go_best_practices/homework2/pkg/program"
)

func Example() {
	log := logrus.New()
	log.SetOutput(ioutil.Discard)

	cnfg, err := config.NewAppConfig()
	if err != nil {
		log.Fatal(err)
	}
	uniqueFiles := files.NewUniqueFilesMap(log)

	program := program.NewProgram(cnfg, uniqueFiles, log)
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

		program := program.NewProgram(cnfg, uniqueFiles, log)
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
