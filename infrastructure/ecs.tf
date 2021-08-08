data "aws_iam_policy_document" "assume_role_policy" {
  statement {
    actions = ["sts:AssumeRole"]

    principals {
      type        = "Service"
      identifiers = ["ecs-tasks.amazonaws.com"]
    }
  }
}

resource "aws_iam_role" "ecsTaskExecutionRole" {
  name               = "${var.service}-${var.environment}-execution-task-role"
  assume_role_policy = data.aws_iam_policy_document.assume_role_policy.json
  tags = {
    Name        = "${var.service}-${var.environment}-iam-role"
    Environment = var.environment
  }
}

resource "aws_ecr_repository" "aws-ecr" {
  name = "${var.service}-${var.environment}-ecr"
  tags = {
    Name        = "${var.service}-${var.environment}-ecr"
    Environment = var.environment
  }
}

resource "aws_ecs_cluster" "aws-ecs-cluster" {
  name = "${var.service}-${var.environment}-cluster"
  tags = {
    Name        = "${var.service}-${var.environment}-ecs"
    Environment = var.environment
  }
}

resource "aws_cloudwatch_log_group" "log-group" {
  name = "${var.service}-${var.environment}-logs"
  tags = {
    Application = var.service
    Environment = var.environment
  }
}

resource "aws_ecs_task_definition" "task-definition" {
  family                = "${var.service}-${var.environment}-task"
  container_definitions = <<DEFINITION
  [
    {
      "name": "${var.service}-${var.environment}-container",
      "image": "${aws_ecr_repository.aws-ecr.repository_url}:latest",
      "entryPoint": [],
      "environment": [
        {
          "name": "PORT",
          "value": "${var.port}"
        },
        {
          "name" : "ENVIRONMENT",
          "value": "${var.environment}"
        }
      ],
      "essential": true,
      "logConfiguration": {
        "logDriver": "awslogs",
        "options": {
          "awslogs-group": "${aws_cloudwatch_log_group.log-group.id}",
          "awslogs-region": "${var.aws_region}",
          "awslogs-stream-prefix": "${var.service}-${var.environment}"
        }
      },
      "portMappings": [
        {
          "containerPort": ${var.port},
          "hostPort": ${var.port}
        }
      ],
      "cpu": ${var.task_cpu},
      "memory": ${var.task_memory},
      "networkMode": "awsvpc"
    }
  ]
  DEFINITION

  requires_compatibilities = ["FARGATE"]
  network_mode             = "awsvpc"
  memory                   = var.task_memory
  cpu                      = var.task_cpu
  execution_role_arn       = aws_iam_role.ecsTaskExecutionRole.arn
  task_role_arn            = aws_iam_role.ecsTaskExecutionRole.arn

  tags = {
    Name        = "${var.service}-${var.environment}-ecs-td"
    Environment = var.environment
  }
}

resource "aws_ecs_service" "aws-ecs-service" {
  name                 = "${var.service}-${var.environment}-ecs-service"
  cluster              = aws_ecs_cluster.aws-ecs-cluster.id
  task_definition      = aws_ecs_task_definition.task-definition.id
  launch_type          = "FARGATE"
  scheduling_strategy  = "REPLICA"
  desired_count        = 2
  force_new_deployment = true

  network_configuration {
    subnets          = aws_subnet.private.*.id
    assign_public_ip = false
    security_groups = [
      aws_security_group.service_security_group.id,
      aws_security_group.load_balancer_security_group.id
    ]
  }

  load_balancer {
    target_group_arn = aws_lb_target_group.target_group.arn
    container_name   = "${var.service}-${var.environment}-container"
    container_port   = var.port
  }

  depends_on = [aws_lb_listener.listener]
}

resource "aws_security_group" "service_security_group" {
  vpc_id = aws_vpc.vpc.id

  ingress {
    from_port       = 0
    to_port         = 0
    protocol        = "-1"
    security_groups = [aws_security_group.load_balancer_security_group.id]
  }

  egress {
    from_port        = 0
    to_port          = 0
    protocol         = "-1"
    cidr_blocks      = ["0.0.0.0/0"]
    ipv6_cidr_blocks = ["::/0"]
  }

  tags = {
    Name = "${var.service}-${var.environment}-service-sg"
  }
}

