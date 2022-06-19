package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/prozz/aws-embedded-metrics-golang/emf"

	runtime "github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-lambda-go/lambdacontext"
)

type Event struct {
	N1 float64 `json:"n1"`
	N2 float64 `json:"n2"`
}

func handleRequest(ctx context.Context, event Event) (string, error) {
	// event
	eventJson, _ := json.MarshalIndent(event, "", "  ")
	log.Printf("EVENT: %s", eventJson)
	//result calculation
	result := event.N1 + event.N2
	result_fmt := fmt.Sprintf("%s%g", "Result=", result)
	// send metric to cloudwatch via EMF
	emf.New().Namespace("LambdaExtra").Dimension("FunctionName", lambdacontext.FunctionName).
		MetricFloatAs("ResultsFromCalculations", result, emf.None).Log()
	return result_fmt, nil
}

func main() {
	runtime.Start(handleRequest)
}
