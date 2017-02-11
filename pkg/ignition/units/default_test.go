package units_test

import (
	"testing"
	"github.com/jgensler8/kother/pkg/ignition/units"
	"io/ioutil"
	"strings"
)

func TestDefaultUnit(t *testing.T) {
	k := units.KubeletUnitBuilder.
		Name("mykubelet").
		HyperkubeImage("myimage").
		HyperkubeTag("latest").
		APIServerDNS("asdf").
		ClusterDNS("qwer").
		ClusterDomain("qwer").
		ManifestPath("/usr/share/oem").
		Build()

	u := units.DefaultUnit(k)
	if len(u.Lines) != 8 {
		t.Fail()
	}
}

func TestDefaultUnit_Serialize(t *testing.T) {
	k := units.KubeletUnitBuilder.
		Name("mykubelet").
		HyperkubeImage("myimage").
		HyperkubeTag("latest").
		APIServerDNS("api.mydomain.com").
		ClusterDNS("10.2.0.10").
		ClusterDomain("mydomain.com").
		ManifestPath("/usr/share/oem").
		Build()

	u := units.DefaultUnit(k)
	r := u.Serialize()
	doc, err := ioutil.ReadAll(r)
	if err != nil {
		t.Fatalf("%v", err)
	}
	if strings.Count(string(doc), "\n") != 23 {
		t.Logf("Lines: %d", strings.Count(string(doc), "\n"))
		t.Logf("%v", string(doc))
		t.Fail()
	}
}