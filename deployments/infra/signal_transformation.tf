resource "aws_lambda_function" "signal_transformation" {
  function_name = "${local.service_name}_signal_transformation"

  role = aws_iam_role.lambda_exec.arn

  image_uri    = "${aws_ecr_repository.registry.repository_url}:signal_transformation-${var.container_image_id}"
  package_type = "Image"
  
}

resource "aws_cloudwatch_log_group" "signal_transformation" {
  name = "/aws/lambda/${aws_lambda_function.signal_transformation.function_name}"

  retention_in_days = 30
}

resource "aws_sns_topic_subscription" "signal_transformation" {
  topic_arn = aws_sns_topic.signal_received.arn
  protocol  = "lambda"
  endpoint  = aws_lambda_function.signal_transformation.arn
}

resource "aws_lambda_permission" "signal_transformation" {
  action        = "lambda:InvokeFunction"
  function_name = aws_lambda_function.signal_transformation.arn
  principal     = "sns.amazonaws.com"
  statement_id  = "AllowSubscriptionToSNS"
  source_arn    = aws_sns_topic.signal_received.arn
}

