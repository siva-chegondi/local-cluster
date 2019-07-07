#! /bin/sh

# need following tools to add repo
apt-get update
apt-get install -y curl apt-transport-https ca-certificates gnupg-agent software-properties-common

# adding gpg key to apt
curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo apt-key add -
sudo apt-key fingerprint 0EBFCD88
sudo add-apt-repository "deb [arch=amd64] https://download.docker.com/linux/ubuntu $(lsb_release -cs) stable"

# installing docker
apt-get update
apt-get install -y docker-ce
 
# add user to docker group
# sudo groupadd docker
sudo usermod -aG docker $USER
