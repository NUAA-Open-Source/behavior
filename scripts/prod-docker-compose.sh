#!/bin/bash
# Author:   TripleZ<me@triplez.cn>
# Date:     2019-04-16

echo -e "\n Build, up, down, restart, pull, check logs for Behavior production docker clusters.\n"

if [ "$1" == "up" ]
then
    mkdir -p ../data
    sudo docker-compose -f ../deployments/prod-behavior/docker-compose.yml up -d

elif [ "$1" == "down" ]
then
    sudo docker-compose -f ../deployments/prod-behavior/docker-compose.yml down

elif [ "$1" == "build" ]
then
    sudo docker-compose -f ../deployments/prod-behavior/docker-compose.yml build --force-rm

elif [ "$1" == "pull" ]
then
    sudo docker-compose -f ../deployments/prod-behavior/docker-compose.yml pull

elif [ "$1" == "restart" ]
then
    sudo docker-compose -f ../deployments/prod-behavior/docker-compose.yml restart -t 10

elif [ "$1" == "logs" ]
then
    echo -e " Follow log output? (y/n, default: n): \c"
    read isF
    echo ""
    if [ "$isF" == "y" ] || [ "$isF" == "Y" ]
    then 
        sudo docker-compose -f ../deployments/prod-behavior/docker-compose.yml logs -f
    else
        sudo docker-compose -f ../deployments/prod-behavior/docker-compose.yml logs
    fi

elif [ "$1" == "help" ] || [ "$1" == "-h" ] || [ "$1" == "--help" ]
then
    echo -e " Usage:
  ./prod-docker-compose.sh [COMMAND]
  ./prod-docker-compose.sh -h|--help

 Commands:
   build      Build Behavior prod container images
   down       Down Behavior prod containers
   help       Show this help message
   logs       View output from prod containers
   pull       Pull Behavior prod container images
   restart    Restart Behavior prod containers
   up         Up Behavior prod containers
   "
else
    echo -e " Cannot match the command \"$1\", please type \"help\" command for help."
fi
