resource "aws_lambda_function" "agent_changelog" {
  function_name = "${local.service_name}_agent_changelog"

  role = aws_iam_role.lambda_exec.arn

  image_uri    = "${aws_ecr_repository.registry.repository_url}:agent_changelog-${var.container_image_id}"
  package_type = "Image"
  
}

resource "aws_cloudwatch_log_group" "agent_changelog" {
  name = "/aws/lambda/${aws_lambda_function.agent_changelog.function_name}"

  retention_in_days = 30
}

resource "aws_cloudwatch_event_rule" "every_hour" {
    name = "every-hour"
    description = "Fires every hour"
    schedule_expression = "rate(60 minutes)"
}

resource "aws_cloudwatch_event_target" "check_foo_every_hour" {
    rule = "${aws_cloudwatch_event_rule.every_hour.name}"
    target_id = "check_foo"
    arn = "${aws_lambda_function.agent_changelog.arn}"
}

resource "aws_lambda_permission" "agent_changelog" {
    statement_id = "AllowExecutionFromCloudWatch"
    action = "lambda:InvokeFunction"
    function_name = "${aws_lambda_function.agent_changelog.function_name}"
    principal = "events.amazonaws.com"
    source_arn = "${aws_cloudwatch_event_rule.every_hour.arn}"
}
