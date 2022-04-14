resource "aws_lambda_function" "signal_persistance" {
  function_name = "${local.service_name}_signal_persistance"

  role = aws_iam_role.lambda_exec.arn

  image_uri    = "${aws_ecr_repository.registry.repository_url}:signal_persistance-${var.container_image_id}"
  package_type = "Image"
  
}

resource "aws_cloudwatch_log_group" "signal_persistance" {
  name = "/aws/lambda/${aws_lambda_function.signal_persistance.function_name}"

  retention_in_days = 30
}

resource "aws_sns_topic_subscription" "signal_persistance" {
  topic_arn = aws_sns_topic.signal_received.arn
  protocol  = "lambda"
  endpoint  = aws_lambda_function.signal_persistance.arn
}

resource "aws_lambda_permission" "signal_persistance" {
  action        = "lambda:InvokeFunction"
  function_name = aws_lambda_function.signal_persistance.arn
  principal     = "sns.amazonaws.com"
  statement_id  = "AllowSubscriptionToSNS"
  source_arn    = aws_sns_topic.signal_received.arn
}

