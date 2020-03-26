# Pando

## Environtment Variables
### General :
```
GO111MODULE
PANDO_TIMEZONE
PANDO_PORT

optional:
PANDO_CORS
PANDO_URL_REWRITE
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

## Installation
to be writen.

## API Documentation
RESTful APIs documentation can be opened in url /swagger/index.html. for writing documentation guide, please refer to [guide](https://github.com/swaggo/swag)