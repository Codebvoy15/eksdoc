package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/eks"
)

type EKSResult struct {
	Arn     string
	Version string
	VpcId   string
}

func FetchEKS(ctx context.Context, cluster, region string) (*EKSResult, error) {
	cfg, err := config.LoadDefaultConfig(ctx, config.WithRegion(region))
	if err != nil {
		return nil, err
	}

	client := eks.NewFromConfig(cfg)
	out, err := client.DescribeCluster(ctx, &eks.DescribeClusterInput{
		Name: &cluster,
	})
	if err != nil {
		return nil, err
	}

	return &EKSResult{
		Arn:     *out.Cluster.Arn,
		Version: *out.Cluster.Version,
		VpcId:   *out.Cluster.ResourcesVpcConfig.VpcId,
	}, nil
}
