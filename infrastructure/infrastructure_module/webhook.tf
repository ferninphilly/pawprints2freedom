
locals {
  fb_route_key = "ANY /fb/{proxy+}"
}

resource "aws_apigatewayv2_api" "main_api" {
  name          = var.api_gateway_name
  protocol_type = "HTTP"
}

resource "aws_apigatewayv2_route" "fb" {
  api_id    = aws_apigatewayv2_api.main_api.id
  route_key = local.fb_route_key

  target = "integrations/${aws_apigatewayv2_integration.sqs_integration.id}"
}

resource "aws_apigatewayv2_stage" "fb_stage" {
  api_id      = aws_apigatewayv2_api.main_api.id
  name        = var.environment
  description = "Stage to send to sqs for facebook"
  auto_deploy = true
  default_route_settings {
    throttling_burst_limit = 1
    throttling_rate_limit  = 1
  }
  access_log_settings {
    destination_arn = aws_cloudwatch_log_group.api_gateway_name_log_group.arn
    format          = "$context.requestId $context.error.message"
  }

  route_settings {
    route_key                = aws_apigatewayv2_route.fb.route_key
    detailed_metrics_enabled = true
  }

}

resource "aws_apigatewayv2_integration" "sqs_integration" {
  api_id              = aws_apigatewayv2_api.main_api.id
  credentials_arn     = aws_iam_role.api.arn
  description         = "SQS example"
  integration_type    = "AWS_PROXY"
  integration_subtype = "SQS-SendMessage"
  request_parameters = {
    "QueueUrl"    = aws_sqs_queue.queue.id
    "MessageBody" = "$request.body.message"
  }
}

resource "aws_api_gateway_account" "role_acct" {
  cloudwatch_role_arn = aws_iam_role.api.arn
}

resource "aws_cloudwatch_log_group" "api_gateway_name_log_group" {
  name              = "/aws/apigateway/${var.api_gateway_name}"
  retention_in_days = 7
  lifecycle {
    prevent_destroy = false
  }
}


