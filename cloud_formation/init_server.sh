#!/bin/bash
# Until I can figure out how to get this to work as UserData, I'll simply copy this script to the server
# This needs to be run line by line because for some reason bash does not pick up the updates in the bashrc

if [[ $# -ne 4 ]]; then
    echo "Illegal number of arguments: $#. Usage is:"
    echo "$0 database-password database-host database-port sendgrid-password"
    exit 1
fi


sudo apt-get update
sudo apt-get --assume-yes install git-core curl zlib1g-dev build-essential libssl-dev libreadline-dev libyaml-dev libsqlite3-dev sqlite3 libxml2-dev libxslt1-dev libcurl4-openssl-dev software-properties-common libffi-dev postgresql postgresql-contrib libpq-dev npm
git clone https://github.com/rbenv/rbenv.git .rbenv
echo 'export PATH=$HOME/.rbenv/bin:$PATH' >> ~/.bashrc
echo 'eval "$(rbenv init - bash)"' >> ~/.bashrc
git clone https://github.com/rbenv/ruby-build.git ~/.rbenv/plugins/ruby-build
echo 'export PATH=$HOME/.rbenv/plugins/ruby-build/bin:$PATH' >> ~/.bashrc
source ~/.bashrc
rbenv install -v 3.3.0
rbenv global 3.3.0
echo 'gem: --no-document' > ~/.gemrc
gem install rails -v 7.1
rbenv rehash
sudo npm install -g n
sudo n 18
git clone https://github.com/cgledezma1101/uwr-tournament.git uwr-tournament
cd uwr-tournament/
bundle install
cd app/javascript
npm install
npm run build

screen

DATABASE_NAME=uwr_tournament \
DATABASE_USER=uwr_tournament \
DATABASE_PASSWORD=Alberto1101 \
DATABASE_HOST=uwr-tournaments-db.chdty7r6aebs.us-east-1.rds.amazonaws.com \
DATABASE_PORT=5432 \
SENDGRID_USERNAME=cgledezma1101@gmail.com \
SENDGRID_PASSWORD= \
bundle exec puma --bind tcp://0.0.0.0:8000 --environment development