# FIAP - TechChallenge - User Service

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

### Running directly from go

Option 1: $```go run cmd/client/main.go```

Option 2: $```make run-app```

## Manually testing the API

On directory ```/api``` there's a collection that can be imported on Insomnia or similar so you can test manually the application's API.

## Running the unit tests

Simply run ```make run-tests``` and let the magic happens. At the end it will automatically open an html with the coverage % for every package.
If you don't have Go installed on your machine, don't worry. We've created a container stage that runs the tests and build the application in a separeted environment. The only thing you need to do is:

```make run-tests-in-docker```

We also have the most recently applied unit tests file in this [folder](/unit-tests-results/unit-tests.png) too.

## Infrastructure

This application runs in a k8s cluster. The manifests about the configuration of this application are in this [repository](link-to-gitops).