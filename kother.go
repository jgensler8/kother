package main

import (
	//cf "github.com/jgensler8/kother/pkg/cloudformation"
	//"github.com/aws/aws-sdk-go/aws/session"
	//"github.com/aws/aws-sdk-go/service/cloudformation"
)

func main() {

	//sess, err := session.NewSession()
	//if err != nil {
	//	return
	//}
	//svc := cloudformation.New(sess)
	//stack := cf.DefaultCloudFormation()
	//cf.CreateOrUpdate(svc, stack)
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