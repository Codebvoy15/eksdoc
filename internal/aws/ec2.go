package aws

import (
	"context"

	"eksdoctor/internal/model"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
)

func FetchVPC(ctx context.Context, vpcId, region string) (*model.VPCConfig, error) {
	cfg, _ := config.LoadDefaultConfig(ctx, config.WithRegion(region))
	client := ec2.NewFromConfig(cfg)

	vpcs, err := client.DescribeVpcs(ctx, &ec2.DescribeVpcsInput{
		VpcIds: []string{vpcId},
	})
	if err != nil {
		return nil, err
	}

	vpc := vpcs.Vpcs[0]
	out := &model.VPCConfig{
		Id:   *vpc.VpcId,
		CIDR: *vpc.CidrBlock,
	}

	sn, err := client.DescribeSubnets(ctx, &ec2.DescribeSubnetsInput{
		Filters: []types.Filter{
			{Name: aws.String("vpc-id"), Values: []string{vpcId}},
		},
	})
	if err != nil {
		return nil, err
	}

	for _, s := range sn.Subnets {
		out.Subnets = append(out.Subnets, model.Subnet{
			Id:   *s.SubnetId,
			CIDR: *s.CidrBlock,
			AZ:   *s.AvailabilityZone,
		})
	}

	return out, nil
}

func FetchSecurityGroups(ctx context.Context, vpcId, region string) ([]model.SecurityGroup, error) {
	cfg, _ := config.LoadDefaultConfig(ctx, config.WithRegion(region))
	client := ec2.NewFromConfig(cfg)

	out, err := client.DescribeSecurityGroups(ctx, &ec2.DescribeSecurityGroupsInput{
		Filters: []types.Filter{
			{Name: aws.String("vpc-id"), Values: []string{vpcId}},
		},
	})
	if err != nil {
		return nil, err
	}

	var sgs []model.SecurityGroup
	for _, sg := range out.SecurityGroups {
		entry := model.SecurityGroup{
			Id:   *sg.GroupId,
			Name: *sg.GroupName,
		}

		for _, p := range sg.IpPermissions {
			for _, r := range p.IpRanges {
				entry.Ingress = append(entry.Ingress, model.SGRule{
					Protocol: *p.IpProtocol,
					FromPort: deref(p.FromPort),
					ToPort:   deref(p.ToPort),
					Source:   *r.CidrIp,
				})
			}
		}

		for _, p := range sg.IpPermissionsEgress {
			for _, r := range p.IpRanges {
				entry.Egress = append(entry.Egress, model.SGRule{
					Protocol: *p.IpProtocol,
					FromPort: deref(p.FromPort),
					ToPort:   deref(p.ToPort),
					Source:   *r.CidrIp,
				})
			}
		}

		sgs = append(sgs, entry)
	}

	return sgs, nil
}

func deref(v *int32) int32 {
	if v == nil {
		return -1
	}
	return *v
}
