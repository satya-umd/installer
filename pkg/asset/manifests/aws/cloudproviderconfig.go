package aws

import (
	"bytes"
	"fmt"

	"github.com/pkg/errors"
	ini "gopkg.in/ini.v1"

	awstypes "github.com/openshift/installer/pkg/types/aws"
)

// CloudConfig is the aws cloud provider config.
type CloudConfig struct {
	Global global
}

// global struct of CloudConfig which is currently not inialized.
type global struct {
	Zone                        string `ini:"Zone,omitempty"`
	VPC                         string `ini:"VPC,omitempty"`
	SubnetID                    string `ini:"SubnetID,omitempty"`
	RouteTableID                string `ini:"RouteTableID,omitempty"`
	RoleARN                     string `ini:"RoleARN,omitempty"`
	KubernetesClusterTag        string `ini:"KubernetesClusterTag,omitempty"`
	KubernetesClusterID         string `ini:"KubernetesClusterID,omitempty"`
	DisableSecurityGroupIngress bool   `ini:"DisableSecurityGroupIngress,omitempty"`
	ElbSecurityGroup            string `ini:"ElbSecurityGroup,omitempty"`
	DisableStrictZoneCheck      bool   `ini:"DisableStrictZoneCheck,omitempty"`
}

// serviceOverride struct used for AWS service endpoint override.
type serviceOverride struct {
	Service       string `ini:"Service"`
	Region        string `ini:"Region"`
	URL           string `ini:"URL"`
	SigningRegion string `ini:"SigningRegion,omitempty"`
	SigningMethod string `ini:"SigningMethod,omitempty"`
	SigningName   string `ini:"SigningName,omitempty"`
}

// CloudProviderConfig builds the cloud provider config and reflects to an ini file.
func CloudProviderConfig(params *awstypes.Platform) (string, error) {
	file := ini.Empty()
	config := &CloudConfig{
		Global: global{},
	}
	if err := file.ReflectFrom(config); err != nil {
		return "", errors.Wrap(err, "failed to reflect from config")
	}

	index := 1
	for _, t := range params.CustomRegionOverride {
		s, err := file.NewSection(fmt.Sprintf("ServiceOverride \"%d\"", index))
		if err != nil {
			return "", errors.Wrapf(err, "failed to create section for ServiceOverride")
		}
		if err := s.ReflectFrom(
			&serviceOverride{
				Service: t.Service,
				Region:  params.Region,
				URL:     t.URL,
			}); err != nil {
			return "", errors.Wrapf(err, "failed to reflect from  ServiceOverride")
		}
		index++
	}

	buf := &bytes.Buffer{}
	if _, err := file.WriteTo(buf); err != nil {
		return "", errors.Wrap(err, "failed to write out cloud provider config")
	}

	return buf.String(), nil
}
