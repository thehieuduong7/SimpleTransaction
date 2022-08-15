<!-- PROJECT LOGO -->
<br />
<div align="center">
  <!-- <a href="https://www.tpptechnology.com/">
    <img src="internBE.com/pkg/logo.png" alt="Logo" width="80" height="80">
  </a> -->
    <a name="readme-top"></a>

  <h1 align="center">The Future of Banks</h1>

  <p align="center">
    A research of Golang and PostgreSQL. Including: Transactions, CRUD Users and CRUD Accounts.
    <br />
  </p>
</div>

## Introduction

We are Team 1 Intern at TP&P Technology company

## Build with

What we used to build our project.

<p align="left"><a href="https://azure.microsoft.com/en-in/" target="_blank" rel="noreferrer"> <img src="https://www.vectorlogo.zone/logos/microsoft_azure/microsoft_azure-icon.svg" alt="azure" width="40" height="40"/> </a> <a href="https://git-scm.com/" target="_blank" rel="noreferrer"> <img src="https://www.vectorlogo.zone/logos/git-scm/git-scm-icon.svg" alt="git" width="40" height="40"/> </a> <a href="https://golang.org" target="_blank" rel="noreferrer"> <img src="https://raw.githubusercontent.com/devicons/devicon/master/icons/go/go-original.svg" alt="go" width="40" height="40"/> </a> <a href="https://www.postgresql.org" target="_blank" rel="noreferrer"> <img src="https://raw.githubusercontent.com/devicons/devicon/master/icons/postgresql/postgresql-original-wordmark.svg" alt="postgresql" width="40" height="40"/> </a> <a href="https://postman.com" target="_blank" rel="noreferrer"> <img src="https://www.vectorlogo.zone/logos/getpostman/getpostman-icon.svg" alt="postman" width="40" height="40"/> </a>

Besides, there is a list of major frameworks/libraries used.

- Gin - Go framework
- GORM - GORM library
- PostgreSQL - PostgreSQL database

## Installation

_Step by step to run our project._

1.  First thing first Clone the repo
    ```sh
    git clone https://tucnguyen@dev.azure.com/tucnguyen/Internship_072022/_git/Internship_072022
    ```
2.  Install all the frameworks and libraries needed for the project. They will automatically be installed.
    ```sh
    go mod tidy
    ```
3.  Check all the requirements has been satisfied. Custom your server and
    database

    ```sh
    Host = "localhost"
    Port = "5432"
    Password = "123"
    User = "postgres"
    DBName = "demo1"
    ```

4.  Run the project by command.

    ```sh
    go run main.go
    ```

5.  Please noted that we using Seed Database to create randomly database for Simple Bank account.

    ```sh
    // Seed data
    GetSeedData()
    // get data
    storage.GetDB()
    ```

    <p align="right">(<a href="#readme-top">back to top</a>)</p>

## Usage

    import{
        github.com/brianvoe/gofakeit/v6 v6.18.0
        github.com/gin-gonic/gin v1.8.1
        github.com/joho/godotenv v1.4.0
        github.com/mashingan/smapping v0.1.17
        gorm.io/driver/postgres v1.3.8
        gorm.io/gorm v1.23.8
    }

## Contribute

TODO:
