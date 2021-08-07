# CloudWatch logs for clusters
resource "aws_cloudwatch_log_group" "log-group" {
  name = "${var.app_name}-${var.app_environment}-logs"
  tags = {
    Name = "${var.app}-${var.environment}-cw-logs"
  }
}

