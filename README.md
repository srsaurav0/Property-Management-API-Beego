# Beego API Service

## Table of Contents
- [Introduction](#introduction)
- [Features](#features)
- [Installation](#installation)
- [Usage](#usage)
- [Endpoints](#endpoints)
- [Testing](#testing)

## Introduction
Beego API Service is a web application built using the [Beego](https://beego.me/) framework. It provides APIs to manage property details, including fetching property information based on property ID.

## Features

### 1. Get Property Details

**Description:** Fetches details of a property based on the provided property ID.

### 2. Get Bulk Property Details

**Description:** Fetches details of a number of properties based on the provided property IDs.

### 3. Get Property Images

**Description:** Fetches images of a property based on the provided property ID. Images are sorted by their ***labels***.

## Requirements

- Go 1.19+
  - Run `go version` to check if Go is installed in your system. If not, then install it and after that, go forward.
  - In Windows, ensure that **GOROOT** is added in the environment variables.
- Beego v2.3.4
  - Run `bee version` to check if Beego is installed in your system. If not, then install it and after that, go forward.
  - Instructions are provided in the [Beego Setup](#beego-setup) section for Beego Installation.

---

## Setup and Installation

### Clone the Repository
```bash
git clone https://github.com/srsaurav0/Property-Management-API-Beego.git
cd Property-Management-API-Beego
```

### Beego Setup
- **For Linux**:
  - Download Beego
    ```bash
    go get github.com/beego/bee/v2@latest
    ```
  - Set Path
    ```bash
    go env GOPATH
    export PATH=$PATH:$(go env GOPATH)/bin
    echo 'export PATH=$PATH:$(go env GOPATH)/bin' >> ~/.bashrc
    ```
  - Check if installed successfully
    ```bash
    bee version
    ```
-**For Windows**:
  - Download Beego
    ```bash
    go get github.com/beego/bee/v2@latest
    ```
  - Check if installed successfully
    ```bash
    bee version
    ```

### Install Dependencies
```bash
go mod init beego-api-service
go mod tidy
```

### Configuration
1. Create a folder named **conf** at the root directory. 
2. Create a file named **app.conf** inside that directory.
3. Enter these configurations inside that directory.
   ```bash
   appname = beego-api-service
   httpport = 8080
   runmode = dev

   externalAPIBaseURL = "http://192.168.0.44:8085/dynamodb-s3-os"
   ```


### View Test Coverage
Generate a coverage report:
```bash
go test ./... -coverprofile=coverage.out
go tool cover -html=coverage.out -o coverage.html
open coverage.html
```
View total coverage in terminal:
```bash
go tool cover -func=coverage.out | grep total: | awk '{print $3}'
```
Open `coverage.html` in a browser to view detailed coverage.