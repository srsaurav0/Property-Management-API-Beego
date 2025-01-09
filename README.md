# Beego API Service

## Table of Contents
- [Introduction](#introduction)
- [Features](#features)
- [Setup and Installation](#setup-and-installation)
- [API Description and Usage](#api-description-and-usage)
- [Tests](#tests)

---

## Introduction
Beego API Service is a web application built using the [Beego](https://beego.me/) framework. It provides APIs to manage property details, including fetching property information based on property ID.

---

## Features

### 1. Get Property Details

**Description:** Fetches details of a property based on the provided property ID.

### 2. Get Bulk Property Details

**Description:** Fetches details of a number of properties based on the provided property IDs.

### 3. Get Property Images

**Description:** Fetches images of a property based on the provided property ID. Images are sorted by their ***labels***.

---

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

### Run the Application

Start the server:
```bash
bee run
```

---


## API Description and Usage

### Get Property Details 

**Endpoint:** GET /v1/api/property/details/:propertyId (e.g. HG-72485269838878)

**Description:**
This endpoint will:
- Call an external API to get property reference data 
- Extract detailed property information from S3 
- Transform and combine the data 
- Return a formatted response

**Usage:**
- Open postman app and create a new ***GET*** request setup.
- Enter the url: *`http://localhost:8080/v1/api/property/details/:propertyId`*
- Replace `:propertyId` with a valid property id. For example: `HA-121156550`
- Press the `Send` button to generate the response.

### Bulk Property Fetch

**Endpoint:** GET /v1/api/propertyList?propertyIds=prop-1,prop-2,prop-3 (prop-1, prop-2, prop-3 will be replaced with real property)

**Description:**
This endpoint will:
- Accept comma-separated property IDs 
- Fetch details for each property in parallel using goroutines 
- Prepare response date and return as a list of property details 

**Usage:**
- Open postman app and create a new ***GET*** request setup.
- Enter the url: *`http://localhost:8080/v1/api/propertyList?propertyIds=porp-1,prop-2,prop-3,...`*
- Replace `prop-*` with a valid property id. For example: `BC-4672180,HA-121156550,HA-321568120,HA-3212331066`.
- Add as many property as you want. These property information will be fetched concurrently.
- Press the `Send` button to generate the response.

### Property Images 

**Endpoint:** GET /v1/api/property/:propertyId /gallery/  (*:propertyId* will be replaced with real property)

**Description:**
This endpoint will:
- Fetch image metadata through API 
- Filter images with confidence score > 95  
- Group images by their labels (other, kitchen, bathroom, etc.) 
- Return filtered and grouped image URLs as a list

**Usage:**
- Open postman app and create a new ***GET*** request setup.
- Enter the url: *`http://localhost:8080/v1/api/property/gallery/:propertyId`*
- Replace `:propertyId` with a valid property id. For example: `BC-4672180`.
- Press the `Send` button to generate the response.

---

## Tests

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