#!/bin/bash

echo "
#########################################################################
 ___ _   _ ___ _____ _   _ ____   ___ _   _ ____ _____  _    _     _     
|_ _| \ | |_ _|_   _| | | | __ ) |_ _| \ | / ___|_   _|/ \  | |   | |    
 | ||  \| || |  | | | | | |  _ \  | ||  \| \___ \ | | / _ \ | |   | |    
 | || |\  || |  | |_| |_| | |_) | | || |\  |___) || |/ ___ \| |___| |___ 
|___|_| \_|___| |_(_)\___/|____(_)___|_| \_|____/ |_/_/   \_\_____|_____|
#########################################################################
"

if [[ "${UID}" -ne 0 ]]
then
 echo 'Must execute with sudo or root' >&2
 exit 1
fi

# Ensure system is up to date
sudo apt update -y

# Upgrade the system
sudo apt upgrade -y


# Enable Firewall
sudo ufw enable


# configure the firewall
sudo ufw default deny incoming
sudo ufw default allow outgoing
sudo ufw allow OpenSSH
sudo ufw allow ssh
sudo ufw allow 22/tcp
sudo ufw allow https
sudo ufw allow 443/tcp
sudo ufw allow 8080/tcp
sudo ufw allow 5050/tcp
sudo ufw allow http


# Disabling root login 
echo "PermitRootLogin no" >> /etc/ssh/sshd_config 


# Fail2Ban install - BLOCKS many ssh attempts
sudo apt install fail2ban -y
sudo systemctl start fail2ban
sudo systemctl enable fail2ban

echo "
[sshd]
enabled = true
port = 22
filter = sshd
logpath = /var/log/auth.log
maxretry = 4
" >> /etc/fail2ban/jail.local


# Install essential packs
sudo apt install netdata p7zip-full net-tools speedtest-cli build-essential checkinstall libssl-dev nginx -y


# Docker option install 
echo "
######################################################################################################
Do you want to install docker? If so type y / If you dont want to install enter n
######################################################################################################
"
read $docker

if [[ $docker -eq "y" ]] || [[ $docker -eq "yes" ]]; then
    sudo apt install apt-transport-https ca-certificates curl software-properties-common -y
    curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo apt-key add -
    sudo add-apt-repository "deb [arch=amd64] https://download.docker.com/linux/ubuntu bionic stable"
    sudo apt update -y
    apt-cache policy docker-ce
    sudo apt install docker-ce -y
    sudo apt install docker-compose -y 


    echo "
#####################################################################################################    
                            Congrats Docker has been installed
######################################################################################################
"
    docker -v

else 
    echo "Docker was not installed"
 
fi



# Cleanup
sudo apt autoremove
sudo apt clean 


# add user to docker group to use
sudo usermod -aG docker ubuntu


exit 0
