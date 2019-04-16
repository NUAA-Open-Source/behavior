#! /bin/bash
# Author:   Jinjin Feng, Zhenzhen Zhao
# Date:     2019-04-14

if [ "$1" == "-h" ] || [ "$1" == "--help" ]
then
    echo -e " Usage:
   ./getcert.sh [DOMAIN]    Generate domain HTTPS certification
   ./getcert.sh -h|--help   Get this help message
   "
elif [ "$1" != "" ]
then
    domain=$1
    read -s -p "Authorization password: " password
    echo

    mkdir ~/.secrets/
    mkdir ~/.secrets/certbot/
    chmod 700 ~/.secrets/

    header="authorization: Basic "
    auth="a2os:"$password
    auth=$(echo -n $auth | base64)
    header+="$auth"

    curl -o ~/.secrets/certbot/cloudflare.ini https://api.vvzero.com/certbot/cloudflare.ini -H "$header"
    chmod 400 ~/.secrets/certbot/cloudflare.ini

    # just for Ubuntu/Debian
    echo -e "Install certbot-dns-cloudflare package"
    sudo apt update -y
    sudo apt-get install build-essential libssl-dev libffi-dev python-dev python-pip -y
    # sudo pip install --upgrade --ignore-installed certbot-dns-cloudflare -y
    sudo pip install certbot-dns-cloudflare -y

    certbot certonly --dns-cloudflare --dns-cloudflare-credentials ~/.secrets/certbot/cloudflare.ini -d $domain
else
    echo "Cannot get the domain value, please use \"-h\" or \"--help\" to get help."
fi
