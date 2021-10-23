## T2T Server - GO

> A GO gin implementation of Tutor2Tutee Server

![](header.jpeg)

### Folder Structure
```
.
├── README.md
├── header.jpeg
├── controllers
│   └── student.go
│   └── teacher.go
├── db
│   └── db.go
├── middlewares
│   └── cors.go
├── tests
│   └── test.go
├── models
│   └── student.go
│   └── teacher.go
└── routers
|   ├── router.go
│   └── student.go
│   └── teacher.go
├── main.go
```

## Installation

Install the `go.mod` dependencies using the following command

```
go build
```


## Development setup

Use `gow` to watch over changes

```
gow run .
```
## Run

##### Now enter the address into your browser 

`http://localhost:8080`
