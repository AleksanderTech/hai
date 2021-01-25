# README

## Overview

Simple application exposing a REST API.
Docker compose has been used with two services to simplify spinning it up without having to configure individual parts on your own.

## Quick and simple start with Docker

### Requirements

* Docker Engine
* Docker Compose

### Run

    git clone <repo_url>
    cd hai/
    docker-compose up -d

## Manual

#### Application made with

* Go version: 1.15

#### Using relational database

* PostgreSQL 12.5

### Run

    git clone <repo_url>
    cd hai/database/init
    source setup_db.sh <user> <password> <host> <port>
    cd ../../app/
    go run main.go

### Build

    cd hai/app
    go build -o main

#### Run build app

    ./main

### Configuration

Configuration loaded from yaml files localized in hai/config/

## Testing

    cd hai/app
    go test -v

## Docker

### Build Docker image

    cd hai
    docker build -t hai:<version> .