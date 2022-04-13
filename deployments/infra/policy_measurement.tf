resource "aws_lambda_function" "policy_measurement" {
  function_name = "${local.service_name}_policy_measurement"

  role = aws_iam_role.lambda_exec.arn

  image_uri    = "${aws_ecr_repository.registry.repository_url}:policy_measurement-19"
  package_type = "Image"
  
}

resource "aws_cloudwatch_log_group" "policy_measurement" {
  name = "/aws/lambda/${aws_lambda_function.policy_measurement.function_name}"

  retention_in_days = 30
}

resource "aws_sns_topic_subscription" "policy_measurement" {
  topic_arn = aws_sns_topic.signal_standardized.arn
  protocol  = "lambda"
  endpoint  = aws_lambda_function.policy_measurement.arn
}

resource "aws_lambda_permission" "policy_measurement" {
  action        = "lambda:InvokeFunction"
  function_name = aws_lambda_function.policy_measurement.arn
  principal     = "sns.amazonaws.com"
  statement_id  = "AllowSubscriptionToSNS"
  source_arn    = aws_sns_topic.signal_standardized.arn
}

