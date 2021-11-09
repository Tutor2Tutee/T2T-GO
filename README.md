## T2T Server - GO

> A GO gin implementation of Tutor2Tutee Server

![](assets/header.png)

### Folder Structure
```
.
├── assets
│   └── header.jpeg
├── controllers
│   └── class.go
│   └── quizzes.go
│   └── refresh_token.go
│   └── user.go
├── db
│   └── db.go
├── helpers
│   └── tokenHelper.go
├── middlewares
│   └── cors.go
│   └── jwt_auth.go
├── models
│   └── class.go
│   └── quizzes.go
│   └── user.go
├── repository
|   ├── classRepository.go
│   └── quizRepository.go
│   └── userRepository.go
│   └── init.go
├── routers
|   ├── classes.go
|   ├── quizzes.go
|   ├── routes.go
│   └── user.go
├── tests
│   └── mockDatabase.go
├── main.go
├── README.md
├── go.mod
└── go.sum
```

## Installation

Install the `go.mod` dependencies using the following command

```
go build
```

or using

```
go get
```


## Development setup

Use `gow` to watch over changes

```
gow run .
```
## Run

##### Now enter the address into your browser 

`http://localhost:8080`
