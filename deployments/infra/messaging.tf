resource "aws_sqs_queue" "signal_ingestion" {
  name                  = "${local.service_name}_signal_ingestion"
  fifo_queue            = false
}

resource "aws_sqs_queue_policy" "signal_ingestion_send" {
  queue_url = aws_sqs_queue.signal_ingestion.id

  policy = jsonencode({
    "Version": "2012-10-17",
    "Statement": [
        {
            "Effect": "Allow",
            "Action": [
              "sqs:SendMessage"
            ],
            "Resource": "${aws_sqs_queue.signal_ingestion.arn}",
            "Condition": {
              "ArnEquals": {
                "aws:SourceArn": "${aws_iam_role.lambda_exec.arn}"
              }
            }
        }
    ]
  })
}