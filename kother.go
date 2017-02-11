package main

import (
	"github.com/hashicorp/terraform/builtin/providers/ignition"
	"k8s.io/client-go/kubernetes"
)

func main() {
/*
	template = cloudformation.basic()

	for manifest in manifests:
		ignition = ignition.basic()
		unit = create_unit_from_manifest(manifest)
		ignition.AddUnit(unit)
		for unit in drops_ins:
		  igniition.AddUnit(unit)
		for file in files:
		  igition.addFile(unit)

		cloudformation.addResourceFromIgnition(ignition.build())

	cloudformation.build()

	cfn.commit(cloudformation)
	cfn.wait

	http poll

	kubernetes.OutClusterConfig(kubeconfig)
	for resource in resource:
	  kubectl apply -f resource
*/
}