package aws

type CustomEndpoint struct{
	Service       string `json:"service"`
	URL           string `json:"url"`
}