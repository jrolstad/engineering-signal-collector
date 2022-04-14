resource "aws_lambda_function" "health" {
  function_name = "${local.service_name}_health"

  role = aws_iam_role.lambda_exec.arn

  image_uri    = "${aws_ecr_repository.registry.repository_url}:health_api-${var.container_image_id}"
  package_type = "Image"
  
}

resource "aws_cloudwatch_log_group" "health" {
  name = "/aws/lambda/${aws_lambda_function.health.function_name}"

  retention_in_days = 30
}

resource "aws_apigatewayv2_integration" "health" {
  api_id = aws_apigatewayv2_api.lambda.id

  integration_uri    = aws_lambda_function.health.invoke_arn
  integration_type   = "AWS_PROXY"
  integration_method = "POST"
}

resource "aws_apigatewayv2_route" "health" {
  api_id = aws_apigatewayv2_api.lambda.id

  route_key = "GET /health"
  target    = "integrations/${aws_apigatewayv2_integration.health.id}"
}

resource "aws_lambda_permission" "api_gw_health" {

  statement_id  = "AllowExecutionFromAPIGateway"
  action        = "lambda:InvokeFunction"
  function_name = aws_lambda_function.health.function_name
  principal     = "apigateway.amazonaws.com"

  source_arn = "${aws_apigatewayv2_api.lambda.execution_arn}/*/*"
}