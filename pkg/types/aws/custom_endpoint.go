package aws

type CustomEndpoint struct{
	Service       string `json:"service"`
	Region        string `json:"region"`
	URL           string `json:"url"`
}