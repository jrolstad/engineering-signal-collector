resource "aws_lambda_function" "policy_persistence" {
  function_name = "${local.service_name}_policy_persistence"

  role = aws_iam_role.lambda_exec.arn

  image_uri    = "${aws_ecr_repository.registry.repository_url}:policy_persistence-14"
  package_type = "Image"
  
}

resource "aws_cloudwatch_log_group" "policy_persistence" {
  name = "/aws/lambda/${aws_lambda_function.policy_persistence.function_name}"

  retention_in_days = 30
}

resource "aws_sns_topic_subscription" "policy_persistence" {
  topic_arn = aws_sns_topic.policy_measured.arn
  protocol  = "lambda"
  endpoint  = aws_lambda_function.policy_persistence.arn
}

resource "aws_lambda_permission" "policy_persistence" {
  action        = "lambda:InvokeFunction"
  function_name = aws_lambda_function.policy_persistence.arn
  principal     = "sns.amazonaws.com"
  statement_id  = "AllowSubscriptionToSNS"
  source_arn    = aws_sns_topic.policy_measured.arn
}

