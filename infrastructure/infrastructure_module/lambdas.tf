locals {
  output_path = "${path.module}/lambdas/zips"
  dist_dir    = "${path.module}/lambdas/bin"
}

data "archive_file" "pptf_sqs_data" {
  type        = "zip"
  source_file = "${local.dist_dir}/${var.sqs_lambda_name}/bootstrap"
  output_path = "${local.output_path}/${var.sqs_lambda_name}.zip"
}


resource "aws_lambda_function" "pptf_sqs_worker" {
  function_name    = var.sqs_lambda_name
  description      = "Lambda to receive data from the sqs queue"
  role             = aws_iam_role.lambda.arn
  handler          = "main.lambda_handler"
  memory_size      = 256
  timeout          = 5
  source_code_hash = data.archive_file.pptf_sqs_data.output_base64sha256
  filename         = data.archive_file.pptf_sqs_data.output_path
  depends_on       = [aws_cloudwatch_log_group.lambda_log_group]
  runtime          = var.go_runtime
  tags = {
    Name        = var.stack_name
    Environment = var.environment
  }
  environment {
    variables = {

    }
  }
}

resource "aws_cloudwatch_log_group" "lambda_log_group" {
  name              = "/aws/lambda/${var.sqs_lambda_name}"
  retention_in_days = 7
  lifecycle {
    prevent_destroy = false
  }
}

resource "aws_lambda_event_source_mapping" "sqs_event" {
  event_source_arn = aws_sqs_queue.queue.arn
  function_name    = aws_lambda_function.pptf_sqs_worker.arn
}

