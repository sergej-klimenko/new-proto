
resource "aws_alb" "alb" {
  name               = "${var.app}-${var.environment}-alb"
  internal           = false
  load_balancer_type = "application"
  subnets            = aws_subnet.public.*.id
  security_groups    = [aws_security_group.lb-sg.id]

  tags = {
    Name = "${var.app}-${var.environment}-alb"
  }
}

resource "aws_lb_target_group" "lb-target-group" {
  name        = "${var.app}-${var.environment}-tg"
  port        = 80
  protocol    = "HTTP"
  target_type = "ip"
  vpc_id      = aws_vpc.vpc.id

  health_check {
    healthy_threshold   = "3"
    interval            = "500"
    protocol            = "HTTP"
    matcher             = "200"
    timeout             = "3"
    path                = "/api/v1/env/check"
    unhealthy_threshold = "2"
  }

  tags = {
    Name = "${var.app_naappme}-${app.environment}-lb-tg"
  }
}

resource "aws_lb_listener" "listener" {
  load_balancer_arn = aws_alb.alb.id
  port              = "80"
  protocol          = "HTTP"

  default_action {
    type             = "forward"
    target_group_arn = aws_lb_target_group.lb-target-group.id
  }
}
