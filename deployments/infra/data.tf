resource "aws_dynamodb_table" "policy_result" {
    name = "${local.service_name}_policyresult"
    billing_mode = "PAY_PER_REQUEST"
    attribute {
        name = "id"
        type = "S"
    }
    hash_key = "id"
}

resource "aws_dynamodb_table" "standard_data" {
    name = "${local.service_name}_standarddata"
    billing_mode = "PAY_PER_REQUEST"
    attribute {
        name = "id"
        type = "S"
    }
    hash_key = "id"
}
