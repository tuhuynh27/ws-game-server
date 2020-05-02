# OddGame.io Game Server

## Development

- Install [Go](https://golang.org/) and [Docker](https://docs.docker.com/install/).
- Run `docker-compose up mongo -d` to start MongoDB in background (or you can start with local MongoDB).
- Run `go run ./cmd/odd-game-server/main.go` to start development.

## Production

- Install [Docker](https://docs.docker.com/install/).
- Run `make build` to build.
- Run `make start` to start.
- Run `make stop` to stop.

## Workflow

### Branch naming:

- Feature: feature/add-new-button
- Hotfix: hotfix/fix-bug-abc
- Improvement: improve/improve-button-abc

### Commits:

- Commit should not capitalize first character, example: migrate to hooks

## Stack:

- Go
- MongoDB
- Docker
