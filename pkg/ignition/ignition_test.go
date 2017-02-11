package ignition_test

import (
	"testing"
	"github.com/jgensler8/kother/pkg/ignition"
	"github.com/jgensler8/kother/pkg/ignition/units"
	"github.com/jgensler8/kother/pkg/ignition/files"
)

func TestDefaultIgnition(t *testing.T) {
	k := units.KubeletUnitBuilder.
		Name("mykubelet").
		HyperkubeImage("myimage").
		HyperkubeTag("latest").
		APIServerDNS("api.mydomain.com").
		ClusterDNS("10.2.0.10").
		ClusterDomain("mydomain.com").
		ManifestPath("/usr/share/oem").
		Build()
	x, err := ignition.DefaultIgnition(k, files.Pod{Name: "qwer"})
	if err != nil {
		t.Fatalf("%v", err)
	}
	t.Logf("%v", x)
}
