package aws

import (
	"bytes"
	"encoding/json"
	"strconv"
	"github.com/pkg/errors"

	awstypes "github.com/openshift/installer/pkg/types/aws"
	
)
type CloudConfig struct {
	Global global
	ServiceOverride map[string]serviceOverride
}

type global struct{
	Zone string
	VPC string
	SubnetID string
	RouteTableID string
	RoleARN string
	KubernetesClusterTag string
	KubernetesClusterID string
	DisableSecurityGroupIngress bool
	ElbSecurityGroup string
	DisableStrictZoneCheck bool
}

type serviceOverride struct {
	Service       string
	Region        string
	URL           string
	SigningRegion string
	SigningMethod string
	SigningName   string
}

func CloudProviderConfig(params *awstypes.Platform) (string, error) {
	serviceOverrideObject, err := convertListToMap(params)
	if err != nil {
		return "", errors.Wrap(err, "could not make serviceOverride map")
	}
	config := &CloudConfig{
		Global: global{

		},
		ServiceOverride: serviceOverrideObject,
	}

	buff := &bytes.Buffer{}
	encoder := json.NewEncoder(buff)
	encoder.SetIndent("", "\t")
	if err := encoder.Encode(config); err != nil {
		return "", err
	}
	return buff.String(), nil
}

func convertListToMap(params *awstypes.Platform) (map[string]serviceOverride, error){
	index := 1
	mapObject := make(map[string]serviceOverride, len(params.CustomRegionOverride))
	for _, t := range params.CustomRegionOverride {
		mapObject[strconv.Itoa(index)] =  serviceOverride{
			Service       : t.Service,
			Region        : params.Region,
			URL           : t.URL,
		}
		index++
	}
	return mapObject,nil
}