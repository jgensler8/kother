package ignition_test

import (
	"testing"
	"github.com/jgensler8/kother/pkg/ignition"
	"github.com/jgensler8/kother/pkg/ignition/units"
	"k8s.io/client-go/pkg/api/v1"
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
	x, err := ignition.DefaultIgnition(k, v1.Pod{})
	if err != nil {
		t.Fatalf("%v", err)
	}
	t.Logf("%v", x)
}
