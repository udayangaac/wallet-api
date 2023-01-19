# wallet-api


[![CircleCI](https://dl.circleci.com/status-badge/img/gh/udayangaac/wallet-api/tree/main.svg?style=svg&circle-token=af27c8c67844c384ef4abb4f7c3dd42eb4a6323c)](https://dl.circleci.com/status-badge/redirect/gh/udayangaac/wallet-api/tree/main)

wallet-api tracks wallet's changes and shows a history of wallet’s wealth to everyone.

## Getting Started

###  Without Docker

Following configurations needs to be set to configure `wallet-api`.

- Set below environment variables to configure postgres database.
```shell
export POSTGRES_HOST=127.0.0.1 # Running host of the postgres database.
export POSTGRES_PORT=5432 # Running port of the postgres database.
export POSTGRES_USER=walletapiuser # Username of the postgres database
export POSTGRES_PASS=walletapipwd # Password of the postgres database.
export POSTGRES_DB=walletapidb # Postgres database name.
export POSTGRES_LOGLEVEL=4 # Define the postgres log level.
export POSTGRES_AUTOMIGRATE=Y # Creates all tables automatically. (Y/y = enable)
```
- Download  dependencies.
```zsh
go mod download
```
- Build the project.

```zsh
go build -o bin/wallet-api cmd/wallet-api/wallet_api.go
```
- Run the application.
```zsh
 bin/wallet-api --port=8081
```

### With Docker

Please run the below command to run via docker.  
__Note:__
- Docker runtime should be installed, up and running in the Host computer
- Postgres Database has been integrated with same container and do not need to configure it separately.

```docker
docker compose up
```

