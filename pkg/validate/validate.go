package validate

import (
	"path"
	"text/template"
	"encoding/json"
	"io/ioutil"
	"bytes"
	"fmt"
	"os"
	"github.com/ghodss/yaml"
	"github.com/jgensler8/kother/pkg/spec"
	"k8s.io/client-go/pkg/api/v1"
)

var (
	DefaultManfiestDirectory = "cloud-provider/manifests"
	DefaultAPIServerYAMLFile = "kube-apiserver.yaml"
	DefaultControllerManagerYAMLFile = "kube-controller-manager.yaml"
	DefaultEtcdYAMLFile = "etcd.yaml"
	DefaultSchedulerYAMLFile = "kube-scheduler.yaml"
	DefaultVaultYAMLFile = "vault.yaml"
	DefaultVariableFile = "cluster.yaml"
	DefaultDirectories = []string{DefaultManfiestDirectory}
	DefaultManifestFiles = []string{
		path.Join(DefaultManfiestDirectory, DefaultAPIServerYAMLFile),
		path.Join(DefaultManfiestDirectory, DefaultControllerManagerYAMLFile),
		path.Join(DefaultManfiestDirectory, DefaultEtcdYAMLFile),
		path.Join(DefaultManfiestDirectory, DefaultSchedulerYAMLFile),
		path.Join(DefaultManfiestDirectory, DefaultVaultYAMLFile)}
)

func Validate(cc *spec.CLIContext) (s *spec.Spec, err error) {
	err = Existence(cc)
	if err != nil{
		fmt.Printf("Some Resources don't exist.")
		return nil, err
	}

	s, err = Correctness(cc)
	if err != nil {
		fmt.Printf("All Resources exist but are invlaid.")
		return nil, err
	}

	return
}

func Existence(cc *spec.CLIContext) (err error){
	p := path.Join(cc.WorkDir, DefaultVariableFile)
	err = ensureVariableFileExists(&p)
	if err != nil {
		fmt.Printf("Default Variable File doesn't exist (%v)\n", p)
		return
	}
	for _, d := range DefaultDirectories {
		p := path.Join(cc.WorkDir, d)
		err = ensureDirectoryExists(&p)
		if err != nil {
			fmt.Printf("Directory doesn't exists (%v)\n", p)
			return
		}
	}
	for _, f := range DefaultManifestFiles {
		p := path.Join(cc.WorkDir, f)
		err = ensureYAMLFileExists(&p)
		if err != nil {
			fmt.Printf("Manifest File doesn't exist (%v)\n", p)
			return
		}
	}
	return nil
}

func Correctness(cc *spec.CLIContext) (_ *spec.Spec, err error) {
	v, err := ioutil.ReadFile(path.Join(cc.WorkDir, DefaultVariableFile))
	if err != nil {
		fmt.Printf("Couldn't read the default vars file located at (%s)\n", v)
		return nil, err
	}

	s := spec.Spec{
		Context: spec.CLIContext(*cc),
	}
	err = yaml.Unmarshal(v, &s)
	if err != nil {
		fmt.Printf("Couldn't unmarshal cluster.yaml to struct ()\n")
	}

	AddDefaultVars(&s)

	for _, f := range DefaultManifestFiles {
		// Read
		p := v1.Pod{}
		v, err := ioutil.ReadFile(path.Join(s.Context.WorkDir, f))
		if err != nil {
			fmt.Printf("Couldn't the pod manifest located at (%s)\n", v)
			return nil, err
		}

		// Template
		tmpl, err := template.New("Test").Parse(string(v))
		if err != nil {
			fmt.Printf("Couldn't create template from file (%f)\n", path.Join(s.Context.WorkDir, f))
			fmt.Printf("%v", v)
			return nil, err
		}
		var podTmpl bytes.Buffer
		err = tmpl.Execute(&podTmpl, &s)
		if err != nil {
			fmt.Printf("Couldn't execute template (%v)\n", podTmpl.String())
			return nil, err
		}

		// Validate YAML to JSON
		j, err := yaml.YAMLToJSON(podTmpl.Bytes())
		if err != nil {
			fmt.Printf("Couldn't convert YAML to JSON\n")
			fmt.Printf("%v", podTmpl.String())
			return nil, err
		}

		// Unmarshal
		err = json.Unmarshal(j, &p)
		if err != nil {
			fmt.Printf("Couldn't unmarshal to v1.Pod.\n")
			fmt.Printf("%v\n", podTmpl.String())
			return nil, err
		}

		addPod(&s, &p)
	}

	return &s, nil
}

func AddDefaultReplicas(c *spec.Component) () {
	if c.Replicas == 0 {
		c.Replicas = 3
	}
}

func AddDefaultVars(s *spec.Spec) (_ *spec.Spec){
	AddDefaultReplicas(&s.Etcd.Component)
	AddDefaultReplicas(&s.Vault.Component)
	AddDefaultReplicas(&s.APIServer.Component)
	AddDefaultReplicas(&s.ControllerManager.Component)
	AddDefaultReplicas(&s.Scheduler.Component)
	s.Config.DNS.EtcdDNS = fmt.Sprintf("etcd.%s", s.Config.DNS.RootDomain)
	s.Config.DNS.VaultDNS = fmt.Sprintf("vault.%s", s.Config.DNS.RootDomain)
	s.Config.DNS.APIServerDNS = fmt.Sprintf("apiserver.%s", s.Config.DNS.RootDomain)
	s.Config.DNS.ControllerManagerDNS = fmt.Sprintf("controller-manager.%s", s.Config.DNS.RootDomain)
	s.Config.DNS.SchedulerDNS = fmt.Sprintf("scheduler.%s", s.Config.DNS.RootDomain)
	return s
}

func ensureVariableFileExists(f *string) (err error) {
	_, err = os.Stat(*f)
	if err != nil {
		return err
	}
	return nil
}

func ensureYAMLFileExists(f *string) (err error) {
	i, err := os.Stat(*f)
	if err != nil || i.IsDir() {
		return err
	}
	return nil
}

func ensureDirectoryExists(d *string) (err error) {
	i, err := os.Stat(*d)
	if err != nil || ( ! i.IsDir() ){
		return err
	}
	return nil
}

func addPod(s *spec.Spec, p *v1.Pod) (err error){
	switch p.Name {
	case "etcd":
		s.Etcd.Pod = p
	case "vault":
		s.Vault.Pod = p
	case "kube-apiserver":
		s.APIServer.Pod = p
		defaultHyperkubeComponent(s.APIServer.Component, &s.Hyperkube)
	case "kube-controller-manager":
		s.ControllerManager.Pod = p
		defaultHyperkubeComponent(s.ControllerManager.Component, &s.Hyperkube)
	case "kube-scheduler":
		s.Scheduler.Pod = p
		defaultHyperkubeComponent(s.Scheduler.Component, &s.Hyperkube)
	default:
		return fmt.Errorf("Unknown component in manifest directory")
	}
	return nil
}

func defaultHyperkubeComponent(c interface{}, h *spec.Hyperkube) {
	s := c.(spec.Component)
	if s.Image.Name == "" {
		s.Image.Tag = h.Image.Tag
	}
	if s.Image.Tag == "" {
		s.Image.Tag = h.Image.Tag
	}
}