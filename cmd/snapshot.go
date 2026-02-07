package cmd

import (
	"context"
	"encoding/json"
	"os"

	"eksdoctor/internal/aws"
	"eksdoctor/internal/model"

	"github.com/spf13/cobra"
)

var (
	cluster string
	region  string
	out     string
)

var snapshotCmd = &cobra.Command{
	Use: "snapshot",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		eksCfg, err := aws.FetchEKS(ctx, cluster, region)
		if err != nil {
			return err
		}

		vpc, err := aws.FetchVPC(ctx, eksCfg.VpcId, region)
		if err != nil {
			return err
		}

		sgs, err := aws.FetchSecurityGroups(ctx, eksCfg.VpcId, region)
		if err != nil {
			return err
		}

		snap := model.ClusterSnapshot{
			ClusterName: cluster,
			Region:      region,
			EKS: model.EKSConfig{
				Arn:     eksCfg.Arn,
				Version: eksCfg.Version,
				VpcId:   eksCfg.VpcId,
			},
			VPC: *vpc,
			Security: model.SecurityView{
				SecurityGroups: sgs,
			},
		}

		data, _ := json.MarshalIndent(snap, "", "  ")
		return os.WriteFile(out, data, 0644)
	},
}




func init() {
	snapshotCmd.Flags().StringVar(&cluster, "cluster", "", "EKS cluster name")
	snapshotCmd.Flags().StringVar(&region, "region", "", "AWS region")
	snapshotCmd.Flags().StringVar(&out, "out", "snapshot.json", "Output file")

	snapshotCmd.MarkFlagRequired("cluster")
	snapshotCmd.MarkFlagRequired("region")

	rootCmd.AddCommand(snapshotCmd)
}
