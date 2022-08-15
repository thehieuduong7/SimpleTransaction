<!-- PROJECT LOGO -->

<div id="header" align="center">
  <img src="https://media3.giphy.com/media/k0ijJhqrUP4T2EvmJ1/giphy.gif?cid=790b7611e84211e835bf53fbbaf6faab8b0562325a2db0f2&rid=giphy.gif&ct=g" width="200"/>
</div>
<br>
<a name="readme-top"></a> 
<div align="center">
  <h1 align="center">Features of Banking</h1>
  <hr>
  <p align="center">
    A research of Golang and PostgreSQL. 
    Including: Transactions, CRUD Users and CRUD Accounts.
  </p>
  
</div>

## Introduction

<hr>
We are Team 1 Intern at TP&P Technology company

## Build with

<hr>
What we used to build our project.
<br>
<br>
<p align="left"><a href="https://azure.microsoft.com/en-in/" target="_blank" rel="noreferrer"> <img src="https://www.vectorlogo.zone/logos/microsoft_azure/microsoft_azure-icon.svg" alt="azure" width="40" height="40"/> </a> <a href="https://git-scm.com/" target="_blank" rel="noreferrer"> <img src="https://www.vectorlogo.zone/logos/git-scm/git-scm-icon.svg" alt="git" width="40" height="40"/> </a> <a href="https://golang.org" target="_blank" rel="noreferrer"> <img src="https://raw.githubusercontent.com/devicons/devicon/master/icons/go/go-original.svg" alt="go" width="40" height="40"/> </a> <a href="https://www.postgresql.org" target="_blank" rel="noreferrer"> <img src="https://raw.githubusercontent.com/devicons/devicon/master/icons/postgresql/postgresql-original-wordmark.svg" alt="postgresql" width="40" height="40"/> </a> <a href="https://postman.com" target="_blank" rel="noreferrer"> <img src="https://www.vectorlogo.zone/logos/getpostman/getpostman-icon.svg" alt="postman" width="40" height="40"/> </a><a href="https://www.docker.com/" target="_blank" rel="noreferrer"> <img src="https://raw.githubusercontent.com/devicons/devicon/master/icons/docker/docker-original-wordmark.svg" alt="docker" width="40" height="40"/> </a>

Based on

- Clean Architecture
- Scrum and Agile methodology

The list of major frameworks/libraries in the project.

- Gin - Go framework
- GORM - GORM library
- PostgreSQL - PostgreSQL database

## Installation

<hr>
_Step by step to run our project._

1.  First thing first Clone the repo
    ```sh
    git clone https://tucnguyen@dev.azure.com/tucnguyen/Internship_072022/_git/Internship_072022
    ```
2.  Install all the frameworks and libraries needed for the project. They will automatically be installed.
    ```sh
    go mod tidy
    ```

Or pull from Docker Hub:

    ```sh
    docker pull dabadaxx/intern-golan
    ```

3.  Check all the requirements has been satisfied. Custom your server and database

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

5.  Using Seed Database to create randomly database for Simple Transaction account.

    ```sh
    // Seed data
    GetSeedData()
    // get data
    storage.GetDB()
    ```

## Usage

<hr>

        import{
         github.com/brianvoe/gofakeit/v6 v6.18.0
         github.com/gin-gonic/gin v1.8.1
         github.com/joho/godotenv v1.4.0
         github.com/mashingan/smapping v0.1.17
         gorm.io/driver/postgres v1.3.8
         gorm.io/gorm v1.23.8
        }

## Contributor

<hr>

- Tuc Nguyen
- Truong Nguyen Xuan
- Hieu Duong The
- Tram Anh Vu
