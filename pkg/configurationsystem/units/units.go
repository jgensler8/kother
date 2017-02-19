package units

import (
	"io"
	"io/ioutil"
	"github.com/lann/builder"
	"github.com/coreos/go-systemd/unit"
)

type UnitFile struct {
	Lines []*unit.UnitOption
}

type unitFileBuilder builder.Builder

func (b unitFileBuilder) AddUnitOption(o *unit.UnitOption) unitFileBuilder {
	return builder.Append(b, "Lines", o).(unitFileBuilder)
}

func (b unitFileBuilder) Build() UnitFile {
	return builder.GetStruct(b).(UnitFile)
}

var UnitFileBuilder = builder.Register(unitFileBuilder{}, UnitFile{}).(unitFileBuilder)

func (f UnitFile) Serialize() (io.Reader) {
	return unit.Serialize(f.Lines)
}

func (f UnitFile) String() (s string, err error) {
	b, err := ioutil.ReadAll(f.Serialize())
	return string(b), err
}