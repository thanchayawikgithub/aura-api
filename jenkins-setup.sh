#!/bin/bash

# Update package list
apt-get update

# Install prerequisites
apt-get install -y \
    apt-transport-https \
    ca-certificates \
    curl \
    software-properties-common \
    wget

# Install Go 1.22.6
wget https://go.dev/dl/go1.22.6.linux-amd64.tar.gz
rm -rf /usr/local/go && tar -C /usr/local -xzf go1.22.6.linux-amd64.tar.gz
export PATH=$PATH:/usr/local/go/bin
echo "export PATH=$PATH:/usr/local/go/bin" >> ~/.bashrc
go version

# Install specific Docker version (27.3.1)
curl -fsSL https://get.docker.com -o get-docker.sh
VERSION=27.3.1 sh get-docker.sh
usermod -aG docker jenkins
systemctl enable docker
docker --version

# Install Docker Compose v2
DOCKER_COMPOSE_VERSION="v2.24.6"  # Latest stable version compatible with Docker 27.3.1
curl -L "https://github.com/docker/compose/releases/download/${DOCKER_COMPOSE_VERSION}/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose
chmod +x /usr/local/bin/docker-compose
docker-compose --version

# Clean up
rm go1.22.6.linux-amd64.tar.gz
rm get-docker.sh 