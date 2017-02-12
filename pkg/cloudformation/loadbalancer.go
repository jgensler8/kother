package cloudformation

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/elb"
)

func CreateLoadBalancer(svc elb.ELB) (lb *elb.CreateLoadBalancerInput, err error) {
	params := &elb.CreateLoadBalancerInput{
		Listeners: []*elb.Listener{ // Required
			{ // Required
				InstancePort:     aws.Int64(1),           // Required
				LoadBalancerPort: aws.Int64(1),           // Required
				Protocol:         aws.String("Protocol"), // Required
				InstanceProtocol: aws.String("Protocol"),
				SSLCertificateId: aws.String("SSLCertificateId"),
			},
			// More values...
		},
		LoadBalancerName: aws.String("AccessPointName"), // Required
		AvailabilityZones: []*string{
			aws.String("AvailabilityZone"), // Required
			// More values...
		},
		Scheme: aws.String("LoadBalancerScheme"),
		SecurityGroups: []*string{
			aws.String("SecurityGroupId"), // Required
			// More values...
		},
		Subnets: []*string{
			aws.String("SubnetId"), // Required
			// More values...
		},
		Tags: []*elb.Tag{
			{ // Required
				Key:   aws.String("TagKey"), // Required
				Value: aws.String("TagValue"),
			},
			// More values...
		},
	}
	//resp, err := svc.CreateLoadBalancer(params)
	return params, nil
}