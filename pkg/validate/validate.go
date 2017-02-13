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
	DefaultAPIServerYAMLFile = "api-server.yaml"
	DefaultControllerManagerYAMLFile = "controller-manager.yaml"
	DefaultEtcdYAMLFile = "etcd.yaml"
	DefaultSchedulerYAMLFile = "scheduler.yaml"
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

func Validate(wd *string) (s *spec.Spec, err error) {
	err = Existence(wd)
	if err != nil{
		fmt.Printf("Some Resources don't exist.")
		return nil, err
	}

	s, err = Correctness(wd)
	if err != nil {
		fmt.Printf("All Resources exist but are invlaid.")
		return nil, err
	}

	return
}

func Existence(wd *string) (err error){
	p := path.Join(*wd, DefaultVariableFile)
	err = ensureVariableFileExists(&p)
	if err != nil {
		fmt.Printf("Default Variable File doesn't exist (%v)\n", p)
		return
	}
	for _, d := range DefaultDirectories {
		p := path.Join(*wd, d)
		err = ensureDirectoryExists(&p)
		if err != nil {
			fmt.Printf("Directory doesn't exists (%v)\n", p)
			return
		}
	}
	for _, f := range DefaultManifestFiles {
		p := path.Join(*wd, f)
		err = ensureYAMLFileExists(&p)
		if err != nil {
			fmt.Printf("Manifest File doesn't exist (%v)\n", p)
			return
		}
	}
	return nil
}

func Correctness(wd *string) (_ *spec.Spec, err error) {
	pods := []*v1.Pod{}

	v, err := ioutil.ReadFile(path.Join(*wd, DefaultVariableFile))
	if err != nil {
		fmt.Printf("Couldn't read the default vars file located at (%s)\n", v)
		return nil, err
	}

	// TODO: make some standard for extra variables
	vars := make(map[string]string)
	err = yaml.Unmarshal(v, &vars)

	vars = AddDefaultVars(vars)

	var c spec.Config
	err = json.Unmarshal(v, &c)
	if err != nil {
		fmt.Printf("Couldn't unmarshal config to spec struct\n")
		return nil, err
	}

	for _, f := range DefaultManifestFiles {
		// Read
		p := v1.Pod{}
		v, err := ioutil.ReadFile(path.Join(*wd, f))
		if err != nil {
			fmt.Printf("Couldn't the pod manifest located at (%s)\n", v)
			return nil, err
		}

		// Template
		tmpl, err := template.New("Test").Parse(string(v))
		if err != nil {
			fmt.Printf("Couldn't create template from file (%f)\n", path.Join(*wd, f))
			fmt.Printf("%v", v)
			return nil, err
		}
		var podTmpl bytes.Buffer
		err = tmpl.Execute(&podTmpl, v)
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
		pods = append(pods, &p)
	}

	s := spec.Spec{
		Pods: pods,
		Config: c,
		ExtraVars: vars,
	}
	return &s, nil
}

func AddDefaultVars(i map[string]string) (_ map[string]string){
	i["EtcdDNS"] = "qwer"
	return i
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