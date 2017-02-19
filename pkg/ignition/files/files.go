package files

import (
	"path"
	"net/url"
	"github.com/coreos/ignition/config/types"
	"github.com/jgensler8/kother/pkg/spec"
	"github.com/emicklei/go-restful/log"
	"encoding/base64"
	"fmt"
	"k8s.io/client-go/pkg/util/json"
)

var (
	BasePath string = "/var/lib/kubelet/pods"
)

func ManifestToFile(c *spec.Component) (_ *types.File, err error) {
	b, err := json.Marshal(c.Pod)
	if err != nil {
		log.Printf("Couldn't marshal component pod for Unit File")
	}
	u, err := url.Parse( fmt.Sprintf("data:text/plain;base64,%s", base64.StdEncoding.EncodeToString(b)) )
	if err != nil {
		return nil, err
	}
	return &types.File{
		Filesystem: "root",
		Path: types.Path(path.Join(BasePath, c.Pod.Name)),
		Mode: 0755,
		User: types.FileUser{
			Id: 123,
		},
		Group: types.FileGroup{
			Id: 123,
		},
		Contents: types.FileContents{
			Source: types.Url(*u),
			Compression: types.Compression(""),
		},
	}, err
}