package nginx

import (
	"bytes"
	"log"
	"text/template"
	"k8s.io/client-go/pkg/api/v1"
	"github.com/Masterminds/sprig"
	"github.com/jgensler8/kother/pkg/spec"
)

func GetNGINXConfig(c *spec.Component, s *spec.Spec) (string, error) {
	t, err := template.New("nginx").
		Funcs(sprig.TxtFuncMap()).
		Funcs(template.FuncMap{ "getPod": GetPod }).
		Parse(nginx_template)
	if err != nil {
		log.Printf("Couldn't create the Vagrantfile template")
		return "", err
	}
	var doc bytes.Buffer
	data := struct {
		Component *spec.Component
		Spec  *spec.Spec
	} {
		c,
		s,
	}
	err = t.Execute(&doc, data)
	if err != nil {
		log.Printf("Couldn't execute Vagrantfile template")
		return "", err
	}
	return doc.String(), nil
}

func GetPod(c *spec.Component) (p v1.Pod) {
	return *c.Pod
}