package vagrantfile

import (
	"fmt"
	"bytes"
	"log"
	"text/template"
	"github.com/jgensler8/kother/pkg/spec"
	"github.com/Masterminds/sprig"
	"github.com/jgensler8/kother/pkg/vagrantfile/nginx"
	"github.com/jgensler8/kother/pkg/configurationsystem"
)

type Vagrantfile struct {
	Contents *string
}

func SpecToVagrantfile(s *spec.Spec) (v *Vagrantfile, err error){
	t, err := template.New("vagrantfile_template").
		Funcs(sprig.TxtFuncMap()).
		Funcs(template.FuncMap{
			"loadBalancerBlock": LoadBalancerBlock,
			"componenetBlock": ComponentBlock,
		}).
		Parse(vagrantfile_template)
	if err != nil {
		log.Printf("Couldn't create the Vagrantfile template")
		return nil, err
	}
	var doc bytes.Buffer
	err = t.Execute(&doc, s)
	if err != nil {
		log.Printf("Couldn't execute Vagrantfile template")
		return nil, err
	}
	r := doc.String()
	return &Vagrantfile {
		Contents: &r,
	}, nil
}

func LoadBalancerBlock(c *spec.Component, s *spec.Spec) (_ string, err error) {
	t, err := template.New("vagrantfile_loadbalancer_template").
		Funcs(sprig.TxtFuncMap()).
		Funcs(template.FuncMap{
			"getPortString": GetPortString,
			"getPod": nginx.GetPod,
			"getNGINXConfig": nginx.GetNGINXConfig,
		}).
		Parse(vagrantfile_loadbalancer_template)
	if err != nil {
		log.Printf("Couldn't create the LoadBalancer template")
		return "", err
	}
	data := struct {
		Component *spec.Component
		Spec  *spec.Spec
	} {
		c,
		s,
	}
	var doc bytes.Buffer
	err = t.Execute(&doc, data)
	if err != nil {
		log.Printf("Couldn't execute LoacBalancer template")
		return "", err
	}
	return doc.String(), nil
}

func ComponentBlock(c *spec.Component, s *spec.Spec) (_ string, err error) {
	t, err := template.New("vagrantfile_component_template").
		Funcs(sprig.TxtFuncMap()).
		Funcs(template.FuncMap{
			"getUserData": configurationsystem.GetUserData,
			"getPod": nginx.GetPod,
		}).
		Parse(vagrantfile_component_template)
	if err != nil {
		log.Printf("Couldn't create the Component template")
		return "", err
	}
	data := struct {
		Component *spec.Component
		Spec  *spec.Spec
	} {
		c,
		s,
	}
	var doc bytes.Buffer
	err = t.Execute(&doc, data)
	if err != nil {
		log.Printf("Couldn't execute Component template")
		return "", err
	}
	return doc.String(), nil
}

func GetPortString(c *spec.Component) (s string) {
	for _, con := range c.Pod.Spec.Containers {
		for _, p := range con.Ports {
			s = fmt.Sprintf(" %s -p %d:%d ", s, p.ContainerPort, p.ContainerPort)
		}
	}
	return
}