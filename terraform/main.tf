provider "aws" {
    region = "us-east-1"
    shared_credentials_file = "/Users/dduleone/.aws/credentials"
    profile                 = "dduleone"
}

# We need:
#   1 x VPC
#   1 x Subnet
#   1 x Internet Gateway
#   1 x SSH Key Pair

#   1 x Security Group
#   2 x Security Group Ingress Rules
#   1 x Security Group Egress Rules
#   1 x Route Table Route
#   1 x Route Table Association

#   1 x EC2 Instance


data "aws_availability_zones" "useast1" {}

data "aws_route_table" "rttble-primary" {
    vpc_id = aws_vpc.dule1.id
}


# 1 x VPC
resource "aws_vpc" "dule1" {
    cidr_block           = var.cidr_vpc
    enable_dns_hostnames = true

    tags = var.alltags
}

# 1 x Subnet
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

# 1 x SSH Key Pair so I can get to the EC2
resource "aws_key_pair" "ssh-key" {
    key_name   = var.ssh_key
    public_key = "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCFlfn2J/RyZWFEFya19HanON9peIDV/SctbrvvOZCKdOt0YStwLsbZqr+DMdXBdJf+7Gjk/wHrRZ3QHNCL8PLMEHt1Hu4zlX82L9YpGNISapoeq+s7L/0Pj/Wb0zOyFPLHn8GU71JRVdzdfFaCJhPwIIZ6JgoPF9z7p3zrlCqncBN1X9SA7yz23o91F+Cs2bAdxVZ/nmtRK1gnVEVtbgHbF8I2iuOjA0tNoFJRjFVQuR9qOxOoCGxUatXno8HFKt2QN4NC+5pLkDfb8ms1ajWitSDzvS0TyyJz/uioozndgiAkWWBOf32HYpjDa4Xe85p0m+qPfZ+IX0W/ZhIv8YYb"

    tags = var.alltags
}

# 1 x Security Groups for the EC2: ssh, http, egress
resource "aws_security_group" "sg-ec2" {
    name        = "${var.app_name}-ec2"
    description = "[${var.app_name}] Security Group for EC2 Instances in the ASG"
    vpc_id      = aws_vpc.dule1.id

    lifecycle {
        create_before_destroy = true
    }
    
    tags = var.alltags
}
    # SSH Ingress
    resource "aws_security_group_rule" "ec2-ingress-ssh" {
        description       = "Provide SSH ingress access to an EC2."
        type              = "ingress"
        from_port         = 22
        to_port           = 22
        protocol          = "tcp"
        cidr_blocks       = ["0.0.0.0/0"]
        security_group_id = aws_security_group.sg-ec2.id
    }
    # HTTP Ingress (port 8080)
    resource "aws_security_group_rule" "ec2-ingress-http" {
        description       = "Provide HTTP ingress access to an EC2."
        type              = "ingress"
        from_port         = 8080
        to_port           = 8080
        protocol          = "tcp"
        cidr_blocks       = ["0.0.0.0/0"]
        security_group_id = aws_security_group.sg-ec2.id
    }
    # World Egress
    resource "aws_security_group_rule" "ec2-egress" {
        description       = "Provide world egress access to an EC2."
        type              = "egress"
        from_port         = 0
        to_port           = 65535
        protocol          = -1
        cidr_blocks       = ["0.0.0.0/0"]
        security_group_id = aws_security_group.sg-ec2.id
    }


# 1 x Route Table Route to connect the Internet Gateway
resource "aws_route" "internet-route" {
    route_table_id              = data.aws_route_table.rttble-primary.id
    destination_cidr_block    = "0.0.0.0/0"
    gateway_id      = aws_internet_gateway.ig.id
}

# 1 x Route Table Association to associate the Subnet with the Route Table
resource "aws_route_table_association" "primarysubnet-route-table" {
    subnet_id     = aws_subnet.primary.id
    route_table_id = data.aws_route_table.rttble-primary.id
}

# 1 x EC2 Instance
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


# Output the URL we can use to test the service, once it's online.
output "test_url" {
    value = "http://${aws_instance.go-server.public_ip}:8080/"
    description = "Use this URL to test the container, once it's all stood up."
}

# Output an ssh command I can use to connect to my EC2 instance, for debugging.
output "ssh_command" {
    value = "ssh -i ./penn_interactive_challenge.pem ec2-user@${aws_instance.go-server.public_ip}"
    description = "For convenience, this will let me connect to the EC2 instance over ssh."
}