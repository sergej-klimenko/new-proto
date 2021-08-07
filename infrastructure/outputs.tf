output "application_url" {
  value = aws_lb.default.dns_name
}