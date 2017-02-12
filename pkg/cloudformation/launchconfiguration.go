package cloudformation

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/autoscaling"
)

func CreateLaunchConfiguration(svc autoscaling.AutoScaling) (lc *autoscaling.CreateLaunchConfigurationInput, err error) {
	params := &autoscaling.CreateLaunchConfigurationInput{
		LaunchConfigurationName:  aws.String("XmlStringMaxLen255"), // Required
		AssociatePublicIpAddress: aws.Bool(true),
		BlockDeviceMappings: []*autoscaling.BlockDeviceMapping{
			{ // Required
				DeviceName: aws.String("XmlStringMaxLen255"), // Required
				Ebs: &autoscaling.Ebs{
					DeleteOnTermination: aws.Bool(true),
					Encrypted:           aws.Bool(true),
					Iops:                aws.Int64(1),
					SnapshotId:          aws.String("XmlStringMaxLen255"),
					VolumeSize:          aws.Int64(1),
					VolumeType:          aws.String("BlockDeviceEbsVolumeType"),
				},
				NoDevice:    aws.Bool(true),
				VirtualName: aws.String("XmlStringMaxLen255"),
			},
			// More values...
		},
		ClassicLinkVPCId: aws.String("XmlStringMaxLen255"),
		ClassicLinkVPCSecurityGroups: []*string{
			aws.String("XmlStringMaxLen255"), // Required
			// More values...
		},
		EbsOptimized:       aws.Bool(true),
		IamInstanceProfile: aws.String("XmlStringMaxLen1600"),
		ImageId:            aws.String("XmlStringMaxLen255"),
		InstanceId:         aws.String("XmlStringMaxLen19"),
		InstanceMonitoring: &autoscaling.InstanceMonitoring{
			Enabled: aws.Bool(true),
		},
		InstanceType:     aws.String("XmlStringMaxLen255"),
		KernelId:         aws.String("XmlStringMaxLen255"),
		KeyName:          aws.String("XmlStringMaxLen255"),
		PlacementTenancy: aws.String("XmlStringMaxLen64"),
		RamdiskId:        aws.String("XmlStringMaxLen255"),
		SecurityGroups: []*string{
			aws.String("XmlString"), // Required
			// More values...
		},
		SpotPrice: aws.String("SpotPrice"),
		UserData:  aws.String("XmlStringUserData"),
	}
	//resp, err := svc.CreateLaunchConfiguration(params)
	return params, nil
}