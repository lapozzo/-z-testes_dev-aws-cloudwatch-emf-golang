# Introduction

Simple AWS Lambda using [Embedded Metrics Format in Golang](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch_Embedded_Metric_Format_Specification.html)

For EMF with Golang, I used the [aws-embedded-metrics-golang library](https://github.com/prozz/aws-embedded-metrics-golang)

## Requirements
* aws-cli/2.0.56 
* Golang 1.17
* aws sam cli 1.15.0

## Start Environment

### 1 Deploy the lambda:

```bash
cd sum-app
make deploy
```

### 2 Test

2.1. Unit Test
```bash
cd sum-app
make test
```

2.2. Local Test with SAM

```bash
cd sum-app
make invoke
```

2.3 Execute the Lambda via Console

### 3 Validate

1. Browse cloudwatch metrics, custom namespace *LambdaExtra*
