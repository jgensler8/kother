package units_test

import (
	"testing"
	"github.com/jgensler8/kother/pkg/ignition/units"
)

func TestKubeletUnitBuilder_Build(t *testing.T) {
	k := units.KubeletUnitBuilder.
		Name("mykubelet").
		HyperkubeImage("myimage").
		HyperkubeTag("latest").
		APIServerDNS("asdf").
		ClusterDNS("qwer").
		ClusterDomain("qwer").
		ManifestPath("/usr/share/oem").
		Build()

	if len(k.Name) == 0 {
		t.Fail()
	}
	if len(k.HyperkubeImage) == 0 {
		t.Fail()
	}
	if len(k.HyperkubeTag) == 0 {
		t.Fail()
	}
	if len(k.APIServerDNS) == 0 {
		t.Fail()
	}
	if len(k.ClusterDNS) == 0 {
		t.Fail()
	}
	if len(k.ClusterDomain) == 0 {
		t.Fail()
	}
	if len(k.ManifestPath) == 0 {
		t.Fail()
	}
}