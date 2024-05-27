package main

import (
	"context"
	"encoding/json"
	"github.com/eegomez/stori-challenge/internal/report"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/eegomez/stori-challenge/cmd/api/configuration"
)

func main() {
	lambda.Start(HandleRequest)
	// Used for local implementation and testing
	//r := gin.Default()
	//r.GET("/ping", func(c *gin.Context) {
	//	c.JSON(200, gin.H{
	//		"message": "pong",
	//	})
	//})
	//cfg := configuration.LoadConfig("erikezgomez@gmail.com")
	//fmt.Println(cfg)
	//rout := router.HandlerFactory(cfg)
	//rout.RegisterRoutes(r)
	//r.Run() // listen and serve on 0.0.0.0:8080
}

type Input struct {
	ID                      int    `json:"id"`
	Message                 string `json:"message"`
	DestinationEmailAddress string `json:"destination_email_address,omitempty"`
}

func HandleRequest(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	requestID := request.RequestContext.RequestID
	log.Printf("Request ID: %s - Request received", requestID)

	var input Input
	if err := json.Unmarshal([]byte(request.Body), &input); err != nil {
		log.Printf("Error unmarshalling request body: %v", err)
		return events.APIGatewayProxyResponse{StatusCode: 400, Body: "Invalid request body"}, nil
	}
	log.Printf("Request ID: %s - Body received: %v", requestID, input)

	cfg := configuration.LoadConfig()
	log.Printf("Request ID: %s - Configuration loaded: %v", requestID, cfg)

	reportUC := report.NewUseCaseFactory(cfg)
	err := reportUC.SendReport(ctx, input.DestinationEmailAddress)
	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: 500, Body: "Internal server error"}, nil
	}

	log.Printf("Request ID: %s - Request finished", requestID)
	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
		Body: "Email has been sent",
	}, nil
}
