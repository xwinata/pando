# Pando

## Environtment Variables
### General :
```
GO111MODULE
PANDO_TIMEZONE
PANDO_PORT
PANDO_NODE ('server' / 'worker')
PANDO_QUEUE_NAME

optional:
PANDO_CORS
PANDO_URL_REWRITE
PANDO_WORKER_NAME (if PANDO_NODE is set to worker mode only)
```
more about GO111MODULE can be found [here](https://github.com/golang/go/wiki/Modules)

### JWT :
```
PANDO_JWT_EXPIRES
PANDO_JWT_SECRET
```
### Database :
```
PANDO_DB_ADAPTER
PANDO_DB_HOST
PANDO_DB_PORT
PANDO_DB_USERNAME
PANDO_DB_PASSWORD
PANDO_DB_LOGMODE(true/false)

optional:
PANDO_DB_SSL(postgres)
PANDO_DB_MAXLIFETIME
PANDO_DB_MAXIDLECONNECTION
PANDO_DB_MAXOPENCONNECTION
```

### AMQP
```
PANDO_AMQP_USER
PANDO_AMQP_PASS
PANDO_AMQP_HOST
PANDO_AMQP_PORT
PANDO_AMQP_RECONNECT_FOREVER
PANDO_AMQP_RECONNECT_RETRIES
PANDO_AMQP_RECONNECT_INTERVAL
PANDO_AMQP_RECONNECT_DEBUGMODE
PANDO_AMQP_EXCHANGE_NAME
```

## Installation
to be writen.

## API Documentation
RESTful APIs documentation can be opened in url /swagger/index.html. for writing documentation guide, please refer to [guide](https://github.com/swaggo/swag)

## Coverages
to generate coverage file use command :
```
go test tests/* -coverprofile=[filename] -coverpkg=./...
```
to read code coverage use command :
```
go tool cover -html=[filename]
```
read more about golang code coverage [here](https://blog.golang.org/cover)