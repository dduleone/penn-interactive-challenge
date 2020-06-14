variable "domain_name" {
    default = "dule1.com"
}

variable "app_name" {
    default = "dule1"
}

variable "alltags" {
    default = {
        "Application" = "penn-interactive-challenge"
        "ManagedBy" = "terraform"
    }
}

variable "userdata_script" {
    default = "userdata.sh"
}

variable "ssh_key" {
    default = "penn-interactive-challenge"
}

variable "ec2_ami" {
    # [us-east-1]: Amazon Linux 2 AMI
    default = "ami-0a887e401f7654935"
}

variable "instance_type" {
    default = "t2.micro"
}

variable "cidr_vpc" {
    default = "10.0.0.0/16"
}
