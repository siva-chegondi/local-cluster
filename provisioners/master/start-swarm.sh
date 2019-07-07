#! /bin/sh

# TODO: check and create group 'vboxfs'
# also add user to that group
# sudo groupadd vboxfs
# sudo usermod -aG vboxfs $USER

# start docker swarm mode
docker swarm init | awk '/--token/ {print $5}' > /docker-master/join-token.log

## Loop service files to create services
for file in /docker-master/services/*
do 
    # services (stacks) creation
    docker stack deploy --compose-file $file
done
