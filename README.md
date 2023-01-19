# wallet-api
wallet-api tracks wallet's changes and shows a history of wallet’s wealth to everyone.


## Getting Started

Following configurations needs to be set to configure `wallet-api`.

- Set below environment variables to configure postgres database.
```shell
export POSTGRES_HOST=127.0.0.1 # Running host of the postgres database.
export POSTGRES_PORT=10104 # Running port of the postgres database.
export POSTGRES_USER=webxg # Username of the postgres database
export POSTGRES_PASS=m3tr0d1me # Password of the postgres database.
export POSTGRES_DB=dfj_webxg # Postgres database name.
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