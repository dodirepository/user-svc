
# User Service

This service running to manage user and also get some token for authentication
### Workflows
![workflow](https://github.com/github/docs/actions/workflows/go.yml/badge.svg?branch=main)

## Development Environment
### Local Development
```bash
- copy and rename .env-example to `.env`
- adjust the value to your config **[optional]**
```
### Using docker
```bash
docker compose -f deployment/docker-compose.yml up 
```
### Running application
By default the application reads the .env file
### STEP 
Makesure you all dependency have been installed

#### 1. Running Migration 
- Note: make sure you have a database name same value with `DB_MYSQL_DATABASE` value in `.env` file
```bash
go run main.go db:migrate up
```
#### 2. Running application
- By default application running on port `:8080` like value `APP_PORT` in `.env` file

```bash
go run main.go serve
```



## API Reference

#### health check

```http
  GET /in/health
```
#### User Seeder for login test
```bash
{
    "username":"admin@gmail.com",
    "password":"admin123"
}
```

#### Postman Collection
- [Collection]((https://awesomeopensource.com/project/elangosundar/awesome-README-templates))

