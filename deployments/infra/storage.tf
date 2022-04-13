resource "aws_s3_bucket" "signal_data" {
  bucket = replace("${local.service_name}data","_","")

}

resource "aws_s3_bucket_acl" "signal_data" {
  bucket = aws_s3_bucket.signal_raw.id
  acl    = "private"
}