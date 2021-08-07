# ECS Cluster
data "aws_iam_policy_document" "assume_by_codedeploy" {
  statement {
    sid     = ""
    effect  = "Allow"
    actions = ["sts:AssumeRole"]

    principals {
      type        = "Service"
      identifiers = ["codedeploy.amazonaws.com"]
    }
  }
}

resource "aws_iam_role" "codedeploy" {
  name               = "${var.service_name}-codedeploy"
  assume_role_policy = data.aws_iam_policy_document.assume_by_codedeploy.json
}


resource "aws_ecs_cluster" "ecs-cluster" {
  name = "${var.app}-${var.environment}-cluster"
  tags = {
    Name = "${var.app}-${var.environment}-ecs-cluster"
  }
}

#ECS Service
resource "aws_ecs_service" "ecs-service" {
  name                 = "${var.app}-${var.environment}-ecs-service"
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
      aws_security_group.ecs-service-sg.id,
      aws_security_group.lb-sg.id
    ]
  }

  load_balancer {
    target_group_arn = aws_lb_target_group.lb-target-group.arn
    container_name   = "${var.app}-${var.environment}-container"
    container_port   = var.port
  }

  depends_on = [aws_lb_listener.listener]
}

resource "aws_ecs_task_definition" "task-definition" {
  family                   = "${var.app}-${var.environment}-task-def"
  execution_role_arn       = aws_iam_role.execution_role.arn
  task_role_arn            = aws_iam_role.task_role.arn
  network_mode             = "awsvpc"
  requires_compatibilities = ["FARGATE"]
  container_definitions = jsonencode([
    {
      name      = "first"
      image     = "service-first"
      cpu       = var.task_cpu
      memory    = var.task_memory
      essential = true
      portMappings = [
        {
          containerPort = var.port
          hostPort      = var.port
        }
      ]
    }
  ])
}
