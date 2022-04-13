resource "aws_lambda_function" "signal_persistance" {
  function_name = "${local.service_name}_signal_persistance"

  role = aws_iam_role.lambda_exec.arn

  image_uri    = "${aws_ecr_repository.registry.repository_url}:signal_persistance-9"
  package_type = "Image"
  
}

resource "aws_cloudwatch_log_group" "persistance" {
  name = "/aws/lambda/${aws_lambda_function.persistance.function_name}"

  retention_in_days = 30
}

resource "aws_lambda_event_source_mapping" "persistance" {
  event_source_arn = aws_sns_topic.signal_received.arn
  function_name    = aws_lambda_function.persistance.arn
}