package model

type ClusterSnapshot struct {
	ClusterName string       `json:"clusterName"`
	Region      string       `json:"region"`
	EKS         EKSConfig    `json:"eks"`
	VPC         VPCConfig    `json:"vpc"`
	Security    SecurityView `json:"security"`
}

type EKSConfig struct {
	Arn     string `json:"arn"`
	Version string `json:"version"`
	VpcId   string `json:"vpcId"`
}

type VPCConfig struct {
	Id      string   `json:"id"`
	CIDR    string   `json:"cidr"`
	Subnets []Subnet `json:"subnets"`
}

type Subnet struct {
	Id   string `json:"id"`
	CIDR string `json:"cidr"`
	AZ   string `json:"az"`
}

type SecurityView struct {
	SecurityGroups []SecurityGroup `json:"securityGroups"`
}

type SecurityGroup struct {
	Id      string   `json:"id"`
	Name    string   `json:"name"`
	Ingress []SGRule `json:"ingress"`
	Egress  []SGRule `json:"egress"`
}

type SGRule struct {
	Protocol string `json:"protocol"`
	FromPort int32  `json:"fromPort"`
	ToPort   int32  `json:"toPort"`
	Source   string `json:"source"` // CIDR or SG ID
}
