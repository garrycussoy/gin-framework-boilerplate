package esb_adapters

import (
	"gin-framework-boilerplate/internal/config"
	ESBPorts "gin-framework-boilerplate/internal/ports/clients/esb"
	Errors "gin-framework-boilerplate/pkg/errors"

	"context"

	"github.com/go-resty/resty/v2"
)

type esbClient struct {
	httpClient *resty.Client
}

func NewESBClient(httpClient *resty.Client) ESBPorts.ESBClient {
	return &esbClient{
		httpClient: httpClient,
	}
}

// Sample service to illustrate the code structure of an HTTP client
func (esbC *esbClient) Sample(ctx context.Context) (ESBPorts.GeneralResponseDTO, Errors.CustomError) {
	// Define some variables needed
	serviceName := "Sample"
	resp := ESBPorts.GeneralResponseDTO{}

	// Hit sample endpoint
	httpResp, httpErr := esbC.httpClient.
		R().
		SetContext(ctx).
		SetHeaders(map[string]string{
			"Authorization": "APIKey",
		}).
		// SetFormData(map[string]string{
		// 	"username": "username",
		// 	"password": "password",
		// }).
		// SetQueryParams(map[string]string{
		// 	"key1": "value1",
		// 	"key2": "value2",
		// }).
		SetBody(map[string]string{
			"username": "username",
			"password": "password",
		}).
		SetResult(&resp).
		// Get(config.AppConfig.ESBHost + "/sample")
		Post(config.AppConfig.ESBHost + "/sample")

	// Handling the response
	handledErr := HandleResponse(serviceName, httpErr, httpResp)
	if handledErr != nil {
		return resp, handledErr
	}

	return resp, nil
}

// Function to handle the response of created HTTP client request (both error and success)
func HandleResponse(serviceName string, httpErr error, httpResp *resty.Response) Errors.CustomError {
	if httpErr != nil {
		return Errors.ESBClientError(serviceName, httpErr.Error())
	}
	if httpResp.StatusCode() != 200 {
		return Errors.ESBClientError(serviceName, httpResp.String())
	}

	return nil
}
