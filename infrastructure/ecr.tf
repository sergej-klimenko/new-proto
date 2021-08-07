resource "aws_ecr_repository" "ecr" {
  name = "${var.app}-${var.environment}-ecr"
  tags = {
    Name = "${var.app}-${var.environment}-ecr"
  }
}
