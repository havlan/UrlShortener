# UrlShortener

### A simple url shortening server
[![Build Status](https://travis-ci.org/havlan/UrlShortener.svg?branch=master)](https://travis-ci.org/havlan/UrlShortener)

- desired:
    - cache most used urls from a cron service in a map to avoid uneccesary reads
    - force https in redirects
    
- future work:
    - block url-shortening for malicious sites, if not possible, hook up site checker for verification
    - create an index.html site
    - use https in general (server, db)

### To run this simple project
- create a database (WEB_URL.sql)
- create a config file or use env
- in src dir (with go workspace, go get...)
- go build .
- ./binary
- postman or similar
    - post "/url" 
    ```
    {
        "url":"www.example.com"
    }
    ``` 
    - localhost:8080/5 (example) provided by the server
    - user tries to access localhost:8080/5 and should then be redirected to url provided
- In a real life setting this can be useful, as long as a short domain name is specified.



### To run this in a docker container
- compile
    - CGO_ENABLED=0 GOOS=linux go build -a -tags netgo -ldflags '-w' .
- build docker image
    - docker build -t ushort . // builds image from dockerfile with name ushort
- run docker container
    - docker run -d -p 8080:8080 --name ushort_ab ushort // forward 8080 => 8080 and container with name ushort_ab from image ushort

### To run this with docker compose
- compile
    - CGO_ENABLED=0 GOOS=linux go build -a -tags netgo -ldflags '-w' .
