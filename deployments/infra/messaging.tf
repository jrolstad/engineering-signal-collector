resource "aws_sqs_queue" "signal_ingestion" {
  name                  = "${local.service_name}_signal_ingestion"
  fifo_queue            = false
}

resource "aws_sns_topic" "signal_received" {
  name = "${local.service_name}_signal_received"
}

resource "aws_sns_topic" "signal_standardized" {
  name = "${local.service_name}_signal_standardized"
}

resource "aws_sns_topic" "policy_measured" {
  name = "${local.service_name}_policy_measured"
}