#!/bin/bash
# Until I can figure out how to get this to work as UserData, I'll simply copy this script to the server
# This needs to be run line by line because for some reason bash does not pick up the updates in the bashrc

if [[ $# -ne 4 ]]; then
    echo "Illegal number of arguments: $#. Usage is:"
    echo "$0 database-password database-host database-port sendgrid-password"
    exit 1
fi


sudo apt-get update
sudo apt-get --assume-yes install git-core curl zlib1g-dev build-essential libssl-dev libreadline-dev libyaml-dev libsqlite3-dev sqlite3 libxml2-dev libxslt1-dev libcurl4-openssl-dev python-software-properties libffi-dev postgresql postgresql-contrib libpq-dev
git clone https://github.com/rbenv/rbenv.git .rbenv
echo 'export PATH=$HOME/.rbenv/bin:$PATH' >> ~/.bashrc
echo 'eval "$(rbenv init - bash)"' >> ~/.bashrc
git clone https://github.com/rbenv/ruby-build.git ~/.rbenv/plugins/ruby-build
echo 'export PATH=$HOME/.rbenv/plugins/ruby-build/bin:$PATH' >> ~/.bashrc
source ~/.bashrc
rbenv install -v 2.7.8
rbenv global 2.7.8
echo 'gem: --no-document' > ~/.gemrc
gem install rails -v 6.0
rbenv rehash
curl -o- https://raw.githubusercontent.com/nvm-sh/nvm/v0.34.0/install.sh | bash
. ~/.nvm/nvm.sh
nvm install 16
git clone https://github.com/cgledezma1101/uwr-tournament.git uwr-tournament
cd uwr-tournament/
bundle install

DATABASE_NAME=uwr_tournament \
DATABASE_USER=uwr_tournament \
DATABASE_PASSWORD=$1 \
DATABASE_HOST=$2 \
DATABASE_PORT=$3 \
SENDGRID_USERNAME=cgledezma1101@gmail.com \
SENDGRID_PASSWORD=$4 \
rails server -e production -p 8000 -d