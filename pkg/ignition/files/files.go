package files

import (
	"path"
	"net/url"
	"github.com/coreos/ignition/config/types"
	"k8s.io/client-go/pkg/api/v1"
)

var (
	BasePath string = "/usr/share/oem"
)

func ManifestToFile(p v1.Pod) (_ *types.File, err error) {
	u, err := url.Parse("data:text/plain;base64,b64filecontentswouldgohere")
	if err != nil {
		return nil, err
	}
	return &types.File{
		Node: types.Node{
			Filesystem: "root",
			Path: types.Path(path.Join(BasePath, p.Name)),
			Mode: 0755,
			User: types.NodeUser{
				Id: 123,
			},
			Group: types.NodeGroup{
				Id: 123,
			},
		},
		Contents: types.FileContents{
			Source: types.Url(*u),
			Compression: types.Compression("gzip+base64"),
		},
	}, err
}