package units_test

import (
	"testing"
	"github.com/coreos/go-systemd/unit"
	"github.com/jgensler8/kother/pkg/configurationsystem/units"
	"io/ioutil"
	"strings"
)

var sampleUnitOption = unit.NewUnitOption("Server", "Restart", "never")

func TestUnitFileBuilder_AddUnitOption(t *testing.T) {
	u := units.UnitFileBuilder.
		AddUnitOption(sampleUnitOption).
		Build()
	if len(u.Lines) != 1 {
		t.Fail()
	}
}

func TestUnitFile_Serialize(t *testing.T) {
	u := units.UnitFileBuilder.
		AddUnitOption(sampleUnitOption).
		Build()
	r := u.Serialize()
	doc, err := ioutil.ReadAll(r)
	if err != nil {
		t.Fatalf("%v", err)
	}
	if strings.Count(string(doc), "\n") != 2 {
		t.Logf("%v", string(doc))
		t.Fail()
	}
}