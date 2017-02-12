package cloudformation

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/autoscaling"
)

func CreateAutoScalingGroup(svc autoscaling.AutoScaling) (a *autoscaling.CreateAutoScalingGroupInput, err error) {
	params := &autoscaling.CreateAutoScalingGroupInput{
		AutoScalingGroupName: aws.String("XmlStringMaxLen255"), // Required
		MaxSize:              aws.Int64(1),                     // Required
		MinSize:              aws.Int64(1),                     // Required
		AvailabilityZones: []*string{
			aws.String("XmlStringMaxLen255"), // Required
			// More values...
		},
		DefaultCooldown:         aws.Int64(1),
		DesiredCapacity:         aws.Int64(1),
		HealthCheckGracePeriod:  aws.Int64(1),
		HealthCheckType:         aws.String("XmlStringMaxLen32"),
		InstanceId:              aws.String("XmlStringMaxLen19"),
		LaunchConfigurationName: aws.String("ResourceName"),
		LoadBalancerNames: []*string{
			aws.String("XmlStringMaxLen255"), // Required
			// More values...
		},
		NewInstancesProtectedFromScaleIn: aws.Bool(true),
		PlacementGroup:                   aws.String("XmlStringMaxLen255"),
		Tags: []*autoscaling.Tag{
			{ // Required
				Key:               aws.String("TagKey"), // Required
				PropagateAtLaunch: aws.Bool(true),
				ResourceId:        aws.String("XmlString"),
				ResourceType:      aws.String("XmlString"),
				Value:             aws.String("TagValue"),
			},
			// More values...
		},
		TargetGroupARNs: []*string{
			aws.String("XmlStringMaxLen511"), // Required
			// More values...
		},
		TerminationPolicies: []*string{
			aws.String("XmlStringMaxLen1600"), // Required
			// More values...
		},
		VPCZoneIdentifier: aws.String("XmlStringMaxLen2047"),
	}
	//resp, err := svc.CreateAutoScalingGroup(params)
	return params, nil
}