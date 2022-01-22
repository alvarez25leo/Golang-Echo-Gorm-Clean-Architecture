# Golang-Echo-Gorm-Clean-Architecture

Implementation of Clean Architecture in Golang RestAPI projects.


![Untitled Diagram](https://user-images.githubusercontent.com/15135199/150567366-5ea535bf-07ac-4049-a24b-44a0cd5c21cf.png)

### Project structure

    ├── main.go
    ├── db
    ├── configs
    ├── middlewares
    ├── routes
    │   └── route.go
    ├── features
    │   ├── user
    │   │   ├── data.go
    │   │   ├── service.go
    │   │   └── handler.go
    │   ├── role
    │   │   ├── data.go
    │   │   ├── service.go
    │   │   └── handler.go
    ├── repositories
    │   ├── models.go
    │   ├── user_repo.go      // interface
    │   ├── role_repo.go      // interface
    │   ├── mysql
    │   │   ├── user_repo.go  // implement
    │   │   └── role_repo.go  // implement
    ├── config.yaml
    ├── docker-compose.yaml
    ├── Makefile


