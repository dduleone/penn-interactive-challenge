provider "aws" {
    region = "us-east-1"
    shared_credentials_file = "/Users/dduleone/.aws/credentials"
    profile                 = "dduleone"
}

# We need:
#   1 x VPC
#   1 x Internet Gateway
#   1 x SSH Key Pair

#   2 x Security Group
#   2 x Security Group Egress Rules
#   3 x Security Group Ingress Rules

#   1 x EC2 Instance


data "aws_availability_zones" "useast1" {}

data "aws_route_table" "rttble-primary" {
    vpc_id = aws_vpc.dule1.id
}


# 1 x VPC should suffice.
resource "aws_vpc" "dule1" {
    cidr_block           = var.cidr_vpc
    enable_dns_hostnames = true

    tags = var.alltags
}

resource "aws_subnet" "primary" {
    vpc_id                  = aws_vpc.dule1.id
    cidr_block              = "10.0.0.0/24"
    availability_zone       = data.aws_availability_zones.useast1.names[0]
    map_public_ip_on_launch = true
    
    tags = var.alltags
}


# 1 x Internet Gateway so public traffic can get in
resource "aws_internet_gateway" "ig" {
    vpc_id = aws_vpc.dule1.id

    tags = var.alltags
}

# 1 x SSH Key Pair so we can get to our EC2s
resource "aws_key_pair" "ssh-key" {
    key_name   = var.ssh_key
    public_key = "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCGKaDKSORFMf/QVM9Dx1D700Di+AAxUTGx4YL05MT4mb9TVqXxAt4hrh9Pg5kkX/RMVdXGxDARt3my3P0cPj2WwhmQM+b8X1Lp2kne9qL0flpkrTtqTrhh0qWS+PE90I8HKFVoRLfMqH2L+2T3ARSx3NyWRMCxCXpQnlawujjF3gcGWWugoN6KHEnlDH6yzDv0wUxVgCw5r6LXFbD0gbmmaoHeaDRkWfCcDmcqZR6uJvDagfyGdFwWdcU8OlW2T2hQnH/6fl1PM8ZqSB6fHkTXEaVK3H4cYzY71aGAdvF4S5yXSvYaWfFVZwfEx/ugM2dXd0QRGu1PNaMrR9rwNwSj"

    tags = var.alltags
}


# 2 x Security Groups, one for our EC2s (ssh, http, egress) and one for our Load Balancer (http)
resource "aws_security_group" "sg-ec2" {
    name        = "${var.app_name}-ec2"
    description = "[${var.app_name}] Security Group for EC2 Instances in the ASG"
    vpc_id      = aws_vpc.dule1.id

    lifecycle {
        create_before_destroy = true
    }
    
    tags = var.alltags
}
    resource "aws_security_group_rule" "ec2-ingress-ssh" {
        description       = "Provide SSH ingress access to the EC2s."
        type              = "ingress"
        from_port         = 22
        to_port           = 22
        protocol          = "tcp"
        cidr_blocks       = ["0.0.0.0/0"]
        security_group_id = aws_security_group.sg-ec2.id
    }
    resource "aws_security_group_rule" "ec2-ingress-http" {
        description       = "Provide HTTP ingress access to the EC2s."
        type              = "ingress"
        from_port         = 8080
        to_port           = 8080
        protocol          = "tcp"
        cidr_blocks       = ["0.0.0.0/0"]
        security_group_id = aws_security_group.sg-ec2.id
    }
    resource "aws_security_group_rule" "ec2-egress" {
        description       = "Provide world egress access to the EC2s."
        type              = "egress"
        from_port         = 0
        to_port           = 65535
        protocol          = -1
        cidr_blocks       = ["0.0.0.0/0"]
        security_group_id = aws_security_group.sg-ec2.id
    }

resource "aws_instance" "go-server" {
    ami           = var.ec2_ami
    instance_type = var.instance_type

    vpc_security_group_ids = [aws_security_group.sg-ec2.id]
    associate_public_ip_address = true
    subnet_id  = aws_subnet.primary.id

    key_name = var.ssh_key

    user_data = filebase64(var.userdata_script)

    tags = var.alltags
}

resource "aws_route" "internet-route" {
    route_table_id              = data.aws_route_table.rttble-primary.id
    destination_cidr_block    = "0.0.0.0/0"
    gateway_id      = aws_internet_gateway.ig.id
}

resource "aws_route_table_association" "primarysubnet-route-table" {
    subnet_id     = aws_subnet.primary.id
    route_table_id = data.aws_route_table.rttble-primary.id
}


output "test_url" {
    value = "http://${aws_instance.go-server.public_ip}:8080/"
    description = "Use this URL to test the container, once it's all stood up."
}

output "ssh_command" {
    value = "ssh -i ~/.ssh/DuLeoneAWSKey.pem ec2-user@${aws_instance.go-server.public_ip}"
    description = "For convenience, this will let me connect to the EC2 instance over ssh."
}