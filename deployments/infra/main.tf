terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 4.0.0"
    }
    random = {
      source  = "hashicorp/random"
      version = "~> 3.1.0"
    }
    archive = {
      source  = "hashicorp/archive"
      version = "~> 2.2.0"
    }

    backend "s3" {
      bucket = "engineeringsignalcollector"
      key    = "engineeringsignalcollector/terraform.tfstate"
      region = "us-west-1"
    }
  }

  required_version = "~> 1.0"
}

provider "aws" {
  region = var.aws_region
}

locals {
  service_name = "engineering_signal_collector_${var.environment}"
}

resource "aws_ecr_repository" "registry" {
  name                 = local.service_name
  image_tag_mutability = "IMMUTABLE"

  image_scanning_configuration {
    scan_on_push = true
  }
}

resource "aws_lambda_function" "health" {
  function_name = "${local.service_name}_health"

  role = aws_iam_role.lambda_exec.arn

  image_uri    = "132367819851.dkr.ecr.eu-west-1.amazonaws.com/complete-cow:1.0" #"${aws_ecr_repository.registry.repository_url}/health-api:latest"
  package_type = "Image"
  
}

resource "aws_cloudwatch_log_group" "health" {
  name = "/aws/lambda/${aws_lambda_function.health.function_name}"

  retention_in_days = 30
}

resource "aws_iam_role" "lambda_exec" {
  name = local.service_name

  assume_role_policy = jsonencode({
    Version = "2012-10-17"
    Statement = [{
      Action = "sts:AssumeRole"
      Effect = "Allow"
      Sid    = ""
      Principal = {
        Service = "lambda.amazonaws.com"
      }
      }
    ]
  })
}

resource "aws_iam_role_policy_attachment" "lambda_policy" {
  name       = local.service_name
  role       = aws_iam_role.lambda_exec.name
  policy_arn = "arn:aws:iam::aws:policy/service-role/AWSLambdaBasicExecutionRole"
}

resource "aws_apigatewayv2_api" "lambda" {
  name          = local.service_name
  protocol_type = "HTTP"
}

resource "aws_apigatewayv2_stage" "lambda" {
  name = local.service_name
  api_id = aws_apigatewayv2_api.lambda.id

  auto_deploy = true

  access_log_settings {
    destination_arn = aws_cloudwatch_log_group.api_gw.arn

    format = jsonencode({
      requestId               = "$context.requestId"
      sourceIp                = "$context.identity.sourceIp"
      requestTime             = "$context.requestTime"
      protocol                = "$context.protocol"
      httpMethod              = "$context.httpMethod"
      resourcePath            = "$context.resourcePath"
      routeKey                = "$context.routeKey"
      status                  = "$context.status"
      responseLength          = "$context.responseLength"
      integrationErrorMessage = "$context.integrationErrorMessage"
      }
    )
  }
}

resource "aws_apigatewayv2_integration" "health" {
  name = "${local.service_name}_health"
  api_id = aws_apigatewayv2_api.lambda.id

  integration_uri    = aws_lambda_function.health.invoke_arn
  integration_type   = "AWS_PROXY"
  integration_method = "POST"
}

resource "aws_apigatewayv2_route" "health" {
  name = "${local.service_name}_health"
  api_id = aws_apigatewayv2_api.lambda.id

  route_key = "GET /health"
  target    = "integrations/${aws_apigatewayv2_integration.health.id}"
}

resource "aws_cloudwatch_log_group" "api_gw" {
  name = "/aws/api_gw/${aws_apigatewayv2_api.lambda.name}"

  retention_in_days = 30
}

resource "aws_lambda_permission" "api_gw" {
  name = "${local.service_name}_health"

  statement_id  = "AllowExecutionFromAPIGateway"
  action        = "lambda:InvokeFunction"
  function_name = aws_lambda_function.health.function_name
  principal     = "apigateway.amazonaws.com"

  source_arn = "${aws_apigatewayv2_api.lambda.execution_arn}/*/*"
}