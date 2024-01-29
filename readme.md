# FIAP - TechChallenge - User Service

## Description

This service is responsible to menage users and customers. We decided to have 2 entities, user is about a login, password and access level in the system, while customer is an extension for user with informations about the person, like document, email, name, etc, with a link to user entitie. This decision was made to have better control about users. Create requests create the both entities. It exposes a grpc channel to comunicate with the other services in the context that requires user datas. We have a diagram about a flow of this service [here](./docs/diagrams/user-service-diagram.png)

## Features

- Create User
- Get Users
- Get User by login
- Get Customers
- Get Customer By UserId
- Get Customer By Document

## How To Run Locally

First of all we need the DataBase. To set it up you have 2 options:

Option 1: $```docker-compose -f docker/db-docker-compose.yml up -d```

Option 2: $```make run-db```

Both are going to have the same result.

Then you can run the application:

### VSCode - Debug
The launch.json file is already configured for debuging. Just hit F5 and be happy.

## Manually testing the API

On directory ```/api``` there's a collection that can be imported on Insomnia or similar so you can test manually the application's API.

## Running the unit tests

Simply run ```make run-tests``` and let the magic happens. At the end it will automatically open an html with the coverage % for every package.
We also have the most recently applied unit tests file in this [folder](./docs/unit-tests-results/unit-tests-user.png) too. And there is a html file about the last unit tests [execution](./docs/unit-tests-results/coverage.html).

## Test + Build + Bake Image

Simply run ```make test-build-bake``` and let the magic happens. The docker file will run the unit-tests, build the application and bake the docker image for the application.

## Infrastructure

This application runs in a k8s cluster. The manifests about the configuration of this application are in this [repository](https://github.com/mauriciodm1998/user-service-gitops).