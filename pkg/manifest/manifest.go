package manifest

import (
	"k8s.io/client-go/pkg/api/v1"
	"io/ioutil"
)

var (
	DefaultManifestPath string = "cloud-provider/manifests"
)

func GetPodsFromManifestDirectory(p string) (s []*v1.Pod, err error){
	files, err := ioutil.ReadDir(p)
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		if ! file.IsDir() {
			c, err := ioutil.ReadFile(file.Name())
			if err != nil {
				return nil, err
			}
			p := v1.Pod{}
			p.Unmarshal(c)
			s = append(s, &p)
		}
	}

	return
}

func GetPodsFromDefaultMAnifestDirectory() (s []*v1.Pod, err error) {
	return GetPodsFromManifestDirectory(DefaultManifestPath)
}