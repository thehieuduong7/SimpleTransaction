# Introduction

We are TP&P Technology Team 1 Intern.

## Build with

Our major frameworks/libraries used on our project.

<p align="center"><p>Programming Languages: </p><a href="https://azure.microsoft.com/en-in/" target="_blank" rel="noreferrer"> <img src="https://www.vectorlogo.zone/logos/microsoft_azure/microsoft_azure-icon.svg" alt="azure" width="40" height="40"/> </a> <a href="https://git-scm.com/" target="_blank" rel="noreferrer"> <img src="https://www.vectorlogo.zone/logos/git-scm/git-scm-icon.svg" alt="git" width="40" height="40"/> </a> <a href="https://golang.org" target="_blank" rel="noreferrer"> <img src="https://raw.githubusercontent.com/devicons/devicon/master/icons/go/go-original.svg" alt="go" width="40" height="40"/> </a> <a href="https://www.postgresql.org" target="_blank" rel="noreferrer"> <img src="https://raw.githubusercontent.com/devicons/devicon/master/icons/postgresql/postgresql-original-wordmark.svg" alt="postgresql" width="40" height="40"/> </a> <a href="https://postman.com" target="_blank" rel="noreferrer"> <img src="https://www.vectorlogo.zone/logos/getpostman/getpostman-icon.svg" alt="postman" width="40" height="40"/> </a>

### Installation

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
    Please noted that we using Seed Database to create randomly database for Simple Bank account.

        ```sh
        // Seed data
        GetSeedData()
         // get data
        storage.GetDB()
        ```

        Then Run

        ```sh
        go run main.go
        ```

<p align="right">(<a href="#readme-top">back to top</a>)</p>

# Contribute

TODO: Explain how other users and developers can contribute to make your code better.
