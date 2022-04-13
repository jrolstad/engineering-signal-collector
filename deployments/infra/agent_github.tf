resource "aws_lambda_function" "agent_github" {
  function_name = "${local.service_name}_agent_github"

  role = aws_iam_role.lambda_exec.arn

  image_uri    = "${aws_ecr_repository.registry.repository_url}:agent_github-6"
  package_type = "Image"
  
}

resource "aws_cloudwatch_log_group" "agent_github" {
  name = "/aws/lambda/${aws_lambda_function.agent_github.function_name}"

  retention_in_days = 30
}

resource "aws_apigatewayv2_integration" "agent_github" {
  api_id = aws_apigatewayv2_api.lambda.id

  integration_uri    = aws_lambda_function.agent_github.invoke_arn
  integration_type   = "AWS_PROXY"
  integration_method = "POST"
  
}

resource "aws_apigatewayv2_route" "agent_github" {
  api_id = aws_apigatewayv2_api.lambda.id

  route_key = "POST /agent_github"
  target    = "integrations/${aws_apigatewayv2_integration.agent_github.id}"
  
}

resource "aws_lambda_permission" "api_gw_agent_github" {

  statement_id  = "AllowExecutionFromAPIGateway"
  action        = "lambda:InvokeFunction"
  function_name = aws_lambda_function.agent_github.function_name
  principal     = "apigateway.amazonaws.com"

  source_arn = "${aws_apigatewayv2_api.lambda.execution_arn}/*/*"
}
