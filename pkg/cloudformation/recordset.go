package cloudformation

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/route53"
)

func ChangeRecordSet(svc route53.Route53) (r *route53.ChangeResourceRecordSetsInput, err error) {
	params := &route53.ChangeResourceRecordSetsInput{
		ChangeBatch: &route53.ChangeBatch{ // Required
			Changes: []*route53.Change{ // Required
				{ // Required
					Action: aws.String("ChangeAction"), // Required
					ResourceRecordSet: &route53.ResourceRecordSet{ // Required
						Name: aws.String("DNSName"), // Required
						Type: aws.String("RRType"),  // Required
						AliasTarget: &route53.AliasTarget{
							DNSName:              aws.String("DNSName"),    // Required
							EvaluateTargetHealth: aws.Bool(true),           // Required
							HostedZoneId:         aws.String("ResourceId"), // Required
						},
						Failover: aws.String("ResourceRecordSetFailover"),
						GeoLocation: &route53.GeoLocation{
							ContinentCode:   aws.String("GeoLocationContinentCode"),
							CountryCode:     aws.String("GeoLocationCountryCode"),
							SubdivisionCode: aws.String("GeoLocationSubdivisionCode"),
						},
						HealthCheckId: aws.String("HealthCheckId"),
						Region:        aws.String("ResourceRecordSetRegion"),
						ResourceRecords: []*route53.ResourceRecord{
							{ // Required
								Value: aws.String("RData"), // Required
							},
							// More values...
						},
						SetIdentifier: aws.String("ResourceRecordSetIdentifier"),
						TTL:           aws.Int64(1),
						TrafficPolicyInstanceId: aws.String("TrafficPolicyInstanceId"),
						Weight:                  aws.Int64(1),
					},
				},
				// More values...
			},
			Comment: aws.String("ResourceDescription"),
		},
		HostedZoneId: aws.String("ResourceId"), // Required
	}
	//resp, err := svc.ChangeResourceRecordSets(params)
	return params, nil
}