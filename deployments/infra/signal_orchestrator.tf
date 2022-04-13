resource "aws_lambda_function" "signal_orchestrator" {
  function_name = "${local.service_name}_signal_orchestrator"

  role = aws_iam_role.lambda_exec.arn

  image_uri    = "${aws_ecr_repository.registry.repository_url}:signal_orchestrator-7"
  package_type = "Image"
  
}

resource "aws_cloudwatch_log_group" "signal_orchestrator" {
  name = "/aws/lambda/${aws_lambda_function.signal_orchestrator.function_name}"

  retention_in_days = 30
}

resource "aws_lambda_event_source_mapping" "signal_orchestrator" {
  event_source_arn = aws_sqs_queue.signal_ingestion.arn
  function_name    = aws_lambda_function.signal_orchestrator.arn
}