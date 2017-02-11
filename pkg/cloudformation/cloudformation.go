package cloudformation

import (
	"github.com/aws/aws-sdk-go/aws"
	cf "github.com/aws/aws-sdk-go/service/cloudformation"
)

/*

CloudformationInput{
	[]Resource
	}

Resource {
	LoadBalancer
	DNS
	ASG
	LC
}

ASG {
	Userdata Ignition
}
 */

func DefaultCloudFormation() (*cf.Stack) {
	return &cf.Stack{

	}
}

func Wait() (err error) {
	return
}

func CreateOrUpdate(svc cf.CloudFormation, stack *cf.Stack) (err error){
	var e bool
	e, err = Exists(svc, stack)
	if e == false {
		err = Create(svc, stack)
		if err != nil {
			return
		} else {
			Wait()
		}
	} else {
		err = Update(svc, stack)
		if err != nil {
			return
		} else {
			Wait()
		}
	}
	return
}

func Exists(svc cf.CloudFormation, stack *cf.Stack) (bool, error) {
	e := Describe(svc, stack)
	if e != nil {
		return false, e
	} else {
		return true, e
	}
}

func Describe(svc cf.CloudFormation, stack *cf.Stack) (err error) {
	params := &cf.DescribeStacksInput{
		NextToken: aws.String("NextToken"),
		StackName: stack.StackName,
	}
	_, err = svc.DescribeStacks(params)
	return
}

func Create(svc cf.CloudFormation, stack *cf.Stack) (err error){
	//params := stack.(cloudformation.CreateStackInput)
	params := &cf.CreateStackInput{
		StackName: aws.String("StackName"), // Required
		Capabilities: []*string{
			aws.String("Capability"), // Required
			// More values...
		},
		DisableRollback: aws.Bool(true),
		NotificationARNs: []*string{
			aws.String("NotificationARN"), // Required
			// More values...
		},
		OnFailure: aws.String("OnFailure"),
		Parameters: []*cf.Parameter{
			{ // Required
				ParameterKey:     aws.String("ParameterKey"),
				ParameterValue:   aws.String("ParameterValue"),
				UsePreviousValue: aws.Bool(true),
			},
			// More values...
		},
		ResourceTypes: []*string{
			aws.String("ResourceType"), // Required
			// More values...
		},
		RoleARN:         aws.String("RoleARN"),
		StackPolicyBody: aws.String("StackPolicyBody"),
		StackPolicyURL:  aws.String("StackPolicyURL"),
		Tags: []*cf.Tag{
			{ // Required
				Key:   aws.String("TagKey"),
				Value: aws.String("TagValue"),
			},
			// More values...
		},
		TemplateBody:     aws.String("TemplateBody"),
		TemplateURL:      aws.String("TemplateURL"),
		TimeoutInMinutes: aws.Int64(1),
	}
	_, err = svc.CreateStack(params)
	return
}

func Update(svc cf.CloudFormation, stack *cf.Stack) (err error){
	//params := stack.(cloudformation.UpdateStackInput)
	params := &cf.UpdateStackInput{
		StackName: aws.String("StackName"), // Required
		Capabilities: []*string{
			aws.String("Capability"), // Required
			// More values...
		},
		NotificationARNs: []*string{
			aws.String("NotificationARN"), // Required
			// More values...
		},
		Parameters: []*cf.Parameter{
			{ // Required
				ParameterKey:     aws.String("ParameterKey"),
				ParameterValue:   aws.String("ParameterValue"),
				UsePreviousValue: aws.Bool(true),
			},
			// More values...
		},
		ResourceTypes: []*string{
			aws.String("ResourceType"), // Required
			// More values...
		},
		RoleARN:                     aws.String("RoleARN"),
		StackPolicyBody:             aws.String("StackPolicyBody"),
		StackPolicyDuringUpdateBody: aws.String("StackPolicyDuringUpdateBody"),
		StackPolicyDuringUpdateURL:  aws.String("StackPolicyDuringUpdateURL"),
		StackPolicyURL:              aws.String("StackPolicyURL"),
		Tags: []*cf.Tag{
			{ // Required
				Key:   aws.String("TagKey"),
				Value: aws.String("TagValue"),
			},
			// More values...
		},
		TemplateBody:        aws.String("TemplateBody"),
		TemplateURL:         aws.String("TemplateURL"),
		UsePreviousTemplate: aws.Bool(true),
	}
	_, err = svc.UpdateStack(params)
	return
}