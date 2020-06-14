#!/bin/bash

# Update yum
yum update -y

# Install docker and git
yum install -y docker git

# Don't require sudo to run docker
usermod -aG docker ec2-user

# Start the docker service
service docker start

# Checkout our code
git clone https://github.com/dduleone/penn-interactive-challenge.git && cd /penn-interactive-challenge

# Pull down dataset, so we can ignore it from our repo
mkdir data
curl https://datasets.imdbws.com/title.basics.tsv.gz | gunzip -c > data/title.basics.tsv

# Build & execute docker
docker build -t penn-interactive-challenge . > docker_build.log
screen -d -m -S docker-run \
    docker run -it --rm -p 8080:80 --name penn-interactive-challenge penn-interactive-challenge

# Used to confirm the script has completed.
echo "Fin." > fin