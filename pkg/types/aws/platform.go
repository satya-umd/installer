package aws

// Platform stores all the global configuration that all machinesets
// use.
type Platform struct {
	// Region specifies the AWS region where the cluster will be created.
	Region string `json:"region"`

	// UserTags specifies additional tags for AWS resources created for the cluster.
	// +optional
	UserTags map[string]string `json:"userTags,omitempty"`

	// DefaultMachinePlatform is the default configuration used when
	// installing on AWS for machine pools which do not define their own
	// platform configuration.
	// +optional
	DefaultMachinePlatform *MachinePool `json:"defaultMachinePlatform,omitempty"`

	//Custom AWS Secret Region Struct
	CustomRegionOverride []CustomEndpoint  `json:"customRegionOverride,omitempty"` 
}

func (p *Platform) FetchCustomEndpointsMap() (map[string]string)  {
	endpointMap := make(map[string]string)	
	if(len(p.CustomRegionOverride) == 0){
		return nil
	}
	for _, endpoint := range p.CustomRegionOverride {
		endpointMap[endpoint.Service] = endpoint.URL
	}
	return endpointMap
}