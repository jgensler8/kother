package manifest_test

//import (
//	"testing"
//	"github.com/jgensler8/kother/pkg/manifest"
//	"path"
//	"path/filepath"
//	"os"
//)

// TODO: Apparently, go doesn't copy any other directories during the test suite. Basically, interaction with the (CONT)
// filesystem should be treated like an integration test
//func TestGetPodsFromManifestDirectory(t *testing.T) {
//	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
//	if err != nil {
//		t.Fatalf("%v", err)
//	}
//	p := path.Join(dir, "../../../", "examples", "myproject", manifest.DefaultManifestPath)
//	//p := path.Join(dir, "../../..
//	t.Logf(p)
//
//	pods, err := manifest.GetPodsFromManifestDirectory(p)
//	if err != nil {
//		t.Fatalf("%v", err)
//	}
//	if len(pods) != 6 {
//		t.Errorf("Only got %d", len(pods))
//	}
//}
