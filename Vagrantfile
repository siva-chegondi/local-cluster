# -*- mode: ruby -*-
# vi: set ft=ruby :

# All Vagrant configuration is done below. The "2" in Vagrant.configure
# configures the configuration version (we support older styles for
# backwards compatibility). Please don't change it unless you know what
# you're doing.
Vagrant.configure("2") do |config|
  # The most common configuration options are documented and commented below.
  # For a complete reference, please see the online documentation at
  # https://docs.vagrantup.com.

  # Every Vagrant development environment requires a box. You can search for
  # boxes at https://vagrantcloud.com/search.
  #
  # DOCKER MASTER
  config.vm.define "docker-manager" do |docmaster|
    docmaster.vm.box = "ubuntu-base"
    docmaster.vm.network "forwarded_port", guest: 80, host: 8080, host_ip: "127.0.0.1"
    docmaster.vm.synced_folder "/data/docker-swarm/master", "/docker-master"
    docmaster.vm.hostname = "docmaster"
    docmaster.vm.provider "virtual-box" do |vb|
      vb.name = "docker-master"
      vb.memory 1024
      vb.cpu 1
    end
    docmaster.vm.provision "shell", path: "provisioners/install-docker.sh"
    docmaster.vm.provision "file", source: "provisioners/services", destination: "/docker-master/services"
    docmaster.vm.provision "shell", path: "provisioners/master/start-swarm.sh"
  end

  # Configuration for docker node to join docker swarm
  #
  # DOCKER NODE
  config.vm.define "docker-node" do |docnode|
    docnode.vm.box = "ubuntu-base"
    docnode.vm.network "forwarded_port", guest: 80, host: 1080, host_ip: "127.0.0.1"
    docnode.vm.synced_folder "/data/docker-swarm/node", "/docker-node"
    docnode.vm.hostname = "docnode"
    docnode.vm.provider "virtual-box" do |vb|
      vb.name = "docker-node"
      vb.memory 1024
      vb.cpu 1
    end
    docnode.vm.provision "shell", path: "provisioners/install-docker.sh"
  end
  

  # Disable automatic box update checking. If you disable this, then
  # boxes will only be checked for updates when the user runs
  # `vagrant box outdated`. This is not recommended.
  # config.vm.box_check_update = false

  # Create a forwarded port mapping which allows access to a specific port
  # within the machine from a port on the host machine. In the example below,
  # accessing "localhost:8080" will access port 80 on the guest machine.
  # NOTE: This will enable public access to the opened port
  # config.vm.network "forwarded_port", guest: 80, host: 8080

  # Create a forwarded port mapping which allows access to a specific port
  # within the machine from a port on the host machine and only allow access
  # via 127.0.0.1 to disable public access
  # config.vm.network "forwarded_port", guest: 80, host: 8080, host_ip: "127.0.0.1"

  # Create a private network, which allows host-only access to the machine
  # using a specific IP.
  # config.vm.network "private_network", ip: "192.168.33.10"

  # Create a public network, which generally matched to bridged network.
  # Bridged networks make the machine appear as another physical device on
  # your network.
  # config.vm.network "public_network"

  # Share an additional folder to the guest VM. The first argument is
  # the path on the host to the actual folder. The second argument is
  # the path on the guest to mount the folder. And the optional third
  # argument is a set of non-required options.
  # config.vm.synced_folder "/data/docker-swarm", "/docker-master"

  # Configuring host name for docker master for easy access
  # config.vm.hostname = "docmaster"


  # Provider-specific configuration so you can fine-tune various
  # backing providers for Vagrant. These expose provider-specific options.
  # Example for VirtualBox:
  #
  # config.vm.provider "virtualbox" do |vb|
  #   # Display the VirtualBox GUI when booting the machine
  #   vb.gui = true
  #
  #   # Customize the amount of memory on the VM:
  #   vb.memory = "1024"
  #   vb.name = "docker-master"
  # end
  #
  # View the documentation for the provider you are using for more
  # information on available options.

  # Enable provisioning with a shell script. Additional provisioners such as
  # Puppet, Chef, Ansible, Salt, and Docker are also available. Please see the
  # documentation for more information about their specific syntax and use.
  # config.vm.provision "shell", inline: <<-SHELL
  #  apt-get update
  #  apt-get install -y nginx
  # SHELL
end
