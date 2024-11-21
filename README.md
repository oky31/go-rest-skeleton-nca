# GO RESTfull Skeleton Project Non Clean Architecture

This project is skeleton project for create rest full API, without using clean architecture concept,
this code acualy from book [https://lets-go-further.alexedwards.net/](https://lets-go-further.alexedwards.net/) but i am only adjust with my need


## `.envrc`

- this file contain config like database 
- config in this file are : 

```
export DBNAME_DB_DSN='drivername://dbuser:password@host/dbname'
```
- for value depend DSN driver you use

- example real value : 

```
export GREENLIGHT_DB_DSN='postgres://greenlight:pa55word@localhost/greenlight'
```

## Running This project 

- this project required with `.envrc`
- for running you can use `make` command :

```
make run/api
```

## Env for Connect to db

- using env variable 

```
export GREENLIGHT_DB_DSN='postgres://greenlight:pa55word@localhost/greenlight'
```
