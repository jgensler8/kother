package files

import (
	"path"
	"encoding/base64"
	"encoding/json"
	"github.com/coreos/coreos-cloudinit/config"
	"github.com/jgensler8/kother/pkg/spec"
	"github.com/emicklei/go-restful/log"
	"fmt"
)

var (
	BasePath string = "/var/lib/kubelet/pods"
)

func ManifestToFile(c *spec.Component) (_ *config.File, err error) {
	b, err := json.Marshal(c.Pod)
	if err != nil {
		log.Printf("Couldn't marshal component pod for Unit File")
	}
	return &config.File{
		Path: path.Join(BasePath, fmt.Sprintf("%s.json", c.Pod.Name)),
		RawFilePermissions: "0755",
		Owner: "root:root",
		Encoding: "base64",
		Content: base64.StdEncoding.EncodeToString(b),
	}, err
}