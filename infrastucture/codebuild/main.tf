

data "aws_caller_identity" "current" {}

data "aws_iam_policy_document" "assume_by_codebuild" {
  statement {
    sid     = "AllowAssumeByCodebuild"
    effect  = "Allow"
    actions = ["sts:AssumeRole"]

    principals {
      type        = "Service"
      identifiers = ["codebuild.amazonaws.com"]
    }
  }
}

resource "aws_s3_bucket" "main" {
  bucket        = "${var.name}-${var.environment}-codepipeline"
  force_destroy = true
}

resource "aws_iam_role" "codebuild" {
  name               = "${var.name}-${var.environment}-codebuild-role"
  assume_role_policy = data.aws_iam_policy_document.assume_by_codebuild.json
}

data "aws_iam_policy_document" "codebuild" {
  statement {
    sid    = "AllowS3"
    effect = "Allow"

    actions = [
      "s3:GetObject",
      "s3:ListBucket",
      "s3:PutObject",
    ]

    resources = [
      "${aws_s3_bucket.main.arn}",
      "${aws_s3_bucket.main.arn}/*",
    ]
  }

  statement {
    sid       = "AllowECRAuth"
    effect    = "Allow"
    actions   = ["ecr:GetAuthorizationToken"]
    resources = ["*"]
  }

  statement {
    sid    = "AllowECRUpload"
    effect = "Allow"
    actions = [
      "ecr:InitiateLayerUpload",
      "ecr:UploadLayerPart",
      "ecr:CompleteLayerUpload",
      "ecr:BatchCheckLayerAvailability",
      "ecr:PutImage",
    ]
    resources = [var.ecr_repository_arn]
  }

  statement {
    sid       = "AllowECSDescribeTaskDefinition"
    effect    = "Allow"
    actions   = ["ecs:DescribeTaskDefinition"]
    resources = ["*"]
  }

  statement {
    sid    = "AllowLogging"
    effect = "Allow"
    actions = [
      "logs:CreateLogGroup",
      "logs:CreateLogStream",
      "logs:PutLogEvents",
    ]
    resources = ["*"]
  }
}

resource "aws_iam_role_policy" "codebuild" {
  role   = aws_iam_role.codebuild.name
  policy = data.aws_iam_policy_document.codebuild.json
}

resource "aws_codebuild_project" "main" {
  name         = "${var.name}-${var.environment}-codebuild"
  description  = "Codebuild for the ECS Green/Blue (${var.name}-${var.environment})"
  service_role = aws_iam_role.codebuild.arn

  artifacts {
    type = "NO_ARTIFACTS"
  }

  environment {
    compute_type    = "BUILD_GENERAL1_SMALL"
    image           = "aws/codebuild/standard:5.0"
    type            = "LINUX_CONTAINER"
    privileged_mode = true

    environment_variable {
      name  = "REPOSITORY_URI"
      value = var.ecr_repository_url
    }

    environment_variable {
      name  = "ECR_NAME"
      value = "${var.name}-${var.environment}"
    }

    environment_variable {
      name  = "REGION"
      value = var.region
    }

    environment_variable {
      name  = "ACCOUNT_ID"
      value = data.aws_caller_identity.current.account_id
    }

    environment_variable {
      name  = "ECS_CLUSTER_NAME"
      value = var.ecs_cluster_name
    }

    environment_variable {
      name  = "ECS_SERVICE_NAME"
      value = var.ecs_service_name
    }
  }

  source {
    type                = "GITHUB"
    location            = "https://github.com/${var.github_owner}/${var.github_repo}.git"
    buildspec           = "pipeline/deploy.yml"
    git_clone_depth     = 1
    report_build_status = true
    git_submodules_config {
      fetch_submodules = true
    }
  }
  source_version = "main"
}

resource "aws_codebuild_source_credential" "main" {
  auth_type   = "PERSONAL_ACCESS_TOKEN"
  server_type = "GITHUB"
  token       = var.github_token
}

resource "aws_codebuild_webhook" "example" {
  project_name = aws_codebuild_project.main.name
  filter_group {
    filter {
      type    = "EVENT"
      pattern = "PUSH"
    }

    filter {
      type    = "HEAD_REF"
      pattern = "main"
    }
  }
}
