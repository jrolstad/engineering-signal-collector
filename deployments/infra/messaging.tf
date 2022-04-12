resource "aws_sqs_queue" "signal_ingestion" {
  name                  = "${local.service_name}_signal_ingestion"
  fifo_queue            = true
  deduplication_scope   = "messageGroup"
  fifo_throughput_limit = "perMessageGroupId"
}