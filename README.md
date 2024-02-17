# uwr-tournaments

This is an application that allows the smooth management of Underwater Rugby tournaments. This repository contains
all the code manage the application. This README outlines how to develop the application.

## Setup

All of the development environment is setup in Docker, so you shouldn't need to manually install anything. However, for
the sake of integrations with IDEs and such, it's useful to know the major dependencies of the project:

1. Ruby 3.3
1. Ruby on Rails 7.1
1. NodeJS 18
1. Postgres 14.1

## Running the application

The dev loop is fully defined using Docker compose. To initialise the development server use:

```bash
docker compose up
```

This will start three containers:

1. A Postgres database
1. A `puma` development server that will initialise the database when launched, and watch for changes.
1. A Webpack build that will watch for changes on the [app/javascript](./app/javascript) folder.

The web server will listen on the container's port 8080, and Docker compose will map it to the host's port 8080. The
web server will also seed the database, if it's empty, using the [seeds.rb file](./db/seeds.rb)
