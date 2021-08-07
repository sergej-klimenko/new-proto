aws_region        = "us-east-2"
aws_access_key    = "your aws access key"
aws_secret_key    = "your aws secret key"

availability_zones = ["us-east-1c", "us-east-1a"]
public_subnets     = ["10.10.100.0/24", "10.10.101.0/24"]
private_subnets    = ["10.10.0.0/24", "10.10.1.0/24"]

app_name        = "golang-api"
app_environment = "staging"