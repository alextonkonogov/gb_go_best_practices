// +build integration

package program_test

import (
	"github.com/sirupsen/logrus"
	"testing"

	"github.com/alextonkonogov/gb_go_best_practices/homework3/pkg/config"
	f "github.com/alextonkonogov/gb_go_best_practices/homework3/pkg/files"
	p "github.com/alextonkonogov/gb_go_best_practices/homework3/pkg/program"
)

func TestIntegration(t *testing.T) {
	log := logrus.New()
	cnfg, err := config.NewAppConfig()
	if err != nil {
		t.Fail()
	}

	cnfg.DeleteDublicates = false

	type tCase struct {
		Path string
		Want int
	}
	tCases := []tCase{
		{
			".",
			3,
		},
		{
			"./..",
			9,
		},
		{
			"./../files",
			4,
		},
		{
			"./../..",
			13,
		},
	}

	for i, tc := range tCases {
		cnfg.Path = tc.Path
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
			t.Error(err)
		}

		if got := len(uniqueFiles.Map); got != tc.Want {
			t.Errorf("on index %v got %v want %v", i, got, tc.Want)
		}
	}
}
