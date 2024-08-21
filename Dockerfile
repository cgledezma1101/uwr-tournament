FROM ruby:3.3.0 as base

# We first install packages that are required to run Rails and its dependencies
RUN apt-get update && apt-get install -y software-properties-common npm make automake gcc

# The Ruby package does not ship with a NodeJS runtime, which we require in Rails. So we install Node here.
# We install NodeJS using an NPM package called n (https://github.com/tj/n)
RUN npm install -g n
RUN n 18

# This makes sure that bundler fails if Gemfile was modified after generating Gemfile.lock.
# You can re-generate Gemfile.lock by running ./update_gemfile.sh
RUN bundle config --global frozen 1

# Install Rails and its dependencies
RUN gem install nokogiri -v 1.16.2
RUN gem install rails -v 7.1

# Create the PID file that Puma needs to run
WORKDIR /uwr-tournaments/tmp/pids
RUN touch server.pid

# Configure and install the project's dependencies
WORKDIR /uwr-tournaments
COPY Gemfile Gemfile.lock ./
RUN bundle install

WORKDIR /uwr-tournaments/app/javascript/
COPY app/javascript/package.json app/javascript/package-lock.json ./
RUN npm install --force

FROM base as development

WORKDIR /uwr-tournaments
ENV DATABASE_NAME=uwr-tournament
ENV DATABASE_USER=uwr-tournament
ENV DATABASE_PASSWORD=admin
ENV DATABASE_HOST=tournaments-database
ENV DATABASE_PORT=5432

CMD rails db:setup;bundle exec puma --bind tcp://0.0.0.0:8080 --environment development
EXPOSE 8080

FROM base as production
WORKDIR /uwr-tournaments
COPY . ./

WORKDIR /uwr-tournaments/app/javascript
RUN npm run build

WORKDIR /uwr-tournaments
# You can get the host from CloudFormation with:
# aws cloudformation describe-stacks --stack-name uwr-tournaments | jq --raw-output '.Stacks[0].Outputs.[] | select( .OutputKey | contains("DatabaseDns") ) | .OutputValue'
ARG databaseHost
ARG databasePort
ARG databasePassword
ARG sendgridPassword

ENV DATABASE_NAME=uwr-tournament
ENV DATABASE_USER=uwr-tournament
ENV DATABASE_PASSWORD=${databasePassword}
ENV DATABASE_HOST=${databaseHost}
ENV DATABASE_PORT=${databasePort}
ENV SENDGRID_USERNAME=cgledezma1101@gmail.com
ENV SENDGRID_PASSWORD=${sendgridPassword}

CMD bundle exec puma --bind tcp://0.0.0.0:8080 --environment production
EXPOSE 8080
