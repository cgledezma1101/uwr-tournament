FROM ruby:2.7.8

# The Ruby package does not ship with a NodeJS runtime, which we require in Rails. So we install Node here.
# We install NodeJS using an NPM package called n (https://github.com/tj/n)
RUN apt-get update
RUN apt-get install -y software-properties-common npm
RUN npm install -g n
RUN n 18

# This makes sure that bundler fails if Gemfile was modified after generating Gemfile.lock.
# You can re-generate Gemfile.lock by running ./update_gemfile.sh
RUN bundle config --global frozen 1

# Install Rails and its dependencies
RUN gem install nokogiri -v 1.15.5
RUN gem install rails -v 6.0

# Configure and install the project's dependencies
WORKDIR /uwr-tournaments
COPY Gemfile Gemfile.lock ./
RUN bundle install

# This is the environment required for the server to run
ENV DATABASE_NAME=uwr-tournament
ENV DATABASE_USER=uwr-tournament
ENV DATABASE_PASSWORD=admin
ENV DATABASE_HOST=tournaments-database
ENV DATABASE_PORT=5432
ENV SENDGRID_USERNAME=cgledezma1101@gmail.com
ENV SENDGRID_PASSWORD=foo

CMD rails db:schema:load;bundle exec puma --bind tcp://0.0.0.0:8080 --environment development
EXPOSE 8080
