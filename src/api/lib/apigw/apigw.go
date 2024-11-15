package apigw

import (
	"net/url"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/apigatewaymanagementapi"
)

func ResolveApiEndpoint(apiClient *apigatewaymanagementapi.Client, domain string, stage string) *apigatewaymanagementapi.Client {
	var endpoint url.URL
	endpoint.Scheme = "https"
	endpoint.Host = domain
	endpoint.Path = stage

	cp := apiClient.Options().Copy()
	cp.BaseEndpoint = aws.String(endpoint.String())
	return apigatewaymanagementapi.New(cp)
}
