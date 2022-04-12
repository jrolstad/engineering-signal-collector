package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	runtime "github.com/aws/aws-lambda-go/lambda"
	"github.com/jrolstad/engineering-signal-collector/internal/pkg/core"
	"github.com/jrolstad/engineering-signal-collector/internal/pkg/orchestration"
)

func init() {
	fmt.Println("Initializing")
}

func main() {
	runtime.Start(handleRequest)
}

func handleRequest(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	data := orchestration.GetApplicationHealth()
	response := core.MapToJson(data)
	return events.APIGatewayProxyResponse{StatusCode: 200, Body: response}, nil
}
