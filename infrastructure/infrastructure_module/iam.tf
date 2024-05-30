// allow lambda service to assume (use) the role with such policy
data "aws_iam_policy_document" "assume_lambda_role" {
  statement {
    actions = ["sts:AssumeRole"]

    principals {
      type        = "Service"
      identifiers = ["lambda.amazonaws.com"]
    }
  }
}

resource "aws_iam_role" "api" {
  name = "AssumeRoleForSQS"

  assume_role_policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Action": "sts:AssumeRole",
      "Principal": {
        "Service": "apigateway.amazonaws.com"
      },
      "Effect": "Allow",
      "Sid": ""
    }
  ]
}
EOF
}

// create lambda role, that lambda function can assume (use)
resource "aws_iam_role" "lambda" {
  name               = "AssumeLambdaRolePPTF"
  description        = "Role allowing permission to assume lambda"
  assume_role_policy = data.aws_iam_policy_document.assume_lambda_role.json
  tags = {
    Name        = var.stack_name
    Environment = var.environment
  }
}

data "aws_iam_policy_document" "allow_lambda_logging" {
  statement {
    effect = "Allow"
    actions = [
      "logs:CreateLogStream",
      "logs:PutLogEvents",
    ]

    resources = [
      "arn:aws:logs:*:*:*",
    ]
  }
}

data "aws_iam_policy_document" "allow_sqs_trigger" {
  statement {
    effect = "Allow"
    actions = [
      "sqs:ReceiveMessage",
      "sqs:DeleteMessage",
      "sqs:GetQueueUrl",
      "sqs:ChangeMessageVisibility",
      "sqs:ListDeadLetterSourceQueues",
      "sqs:SendMessageBatch",
      "sqs:PurgeQueue",
      "sqs:ReceiveMessage",
      "sqs:SendMessage",
      "sqs:GetQueueAttributes",
      "sqs:CreateQueue",
      "sqs:ListQueueTags",
      "sqs:ChangeMessageVisibilityBatch",
      "sqs:SetQueueAttributes"
    ]

    resources = [
      aws_sqs_queue.queue.arn,
    ]
  }
}

// create a policy to allow writing into logs and create logs stream
resource "aws_iam_policy" "function_logging_policy" {
  name        = "AllowLambdaLoggingPolicy"
  description = "Policy for lambda cloudwatch logging"
  policy      = data.aws_iam_policy_document.allow_lambda_logging.json
  tags = {
    Name        = var.stack_name
    Environment = var.environment
  }
}

// attach policy to out created lambda role
resource "aws_iam_role_policy_attachment" "lambda_logging_policy_attachment" {
  role       = aws_iam_role.lambda.id
  policy_arn = aws_iam_policy.function_logging_policy.arn
}

resource "aws_iam_role_policy_attachment" "lambda_triggering_policy_attachment" {
  role       = aws_iam_role.lambda.id
  policy_arn = aws_iam_policy.sqs_perms.arn
}

resource "aws_iam_policy" "sqs_perms" {
  name        = "AllowSQSTriggerForLambda"
  description = "Allows events to trigger a lambda from sqs"
  policy      = data.aws_iam_policy_document.allow_sqs_trigger.json
  tags = {
    Name        = var.stack_name
    Environment = var.environment
  }
}

######################## SSM ACCESS #####################
data "aws_iam_policy_document" "allow_ssm_access" {
  statement {
    effect = "Allow"
    actions = [
      "ssm:CreateAssociation",
      "ssm:CreateDocument",
      "ssm:GetDocument",
      "ssm:ListAssociations",
      "ssm:ListDocuments",
      "ssm:SendCommand",
      "ssm:GetCommandInvocation"
    ]
    resources = [
      "*",
    ]
  }
}

resource "aws_iam_policy" "api" {
  name = "ApiLogAndSQSPolicy"

  policy = <<EOF
{
    "Version": "2012-10-17",
    "Statement": [
      {
        "Effect": "Allow",
        "Action": [
          "logs:CreateLogGroup",
          "logs:CreateLogStream",
          "logs:DescribeLogGroups",
          "logs:DescribeLogStreams",
          "logs:PutLogEvents",
          "logs:GetLogEvents",
          "logs:FilterLogEvents"
        ],
        "Resource": "*"
      },
      {
        "Effect": "Allow",
        "Action": [
          "sqs:GetQueueUrl",
          "sqs:ChangeMessageVisibility",
          "sqs:ListDeadLetterSourceQueues",
          "sqs:SendMessageBatch",
          "sqs:PurgeQueue",
          "sqs:ReceiveMessage",
          "sqs:SendMessage",
          "sqs:GetQueueAttributes",
          "sqs:CreateQueue",
          "sqs:ListQueueTags",
          "sqs:ChangeMessageVisibilityBatch",
          "sqs:SetQueueAttributes"
        ],
        "Resource": "${aws_sqs_queue.queue.arn}"
      },
      {
        "Effect": "Allow",
        "Action": "sqs:ListQueues",
        "Resource": "*"
      }      
    ]
}
EOF
}

resource "aws_iam_role_policy_attachment" "api" {
  role       = aws_iam_role.api.name
  policy_arn = aws_iam_policy.api.arn
}
