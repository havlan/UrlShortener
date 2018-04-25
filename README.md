# UrlShortener

### A simple url shortening server


- desired:
    - cache most used urls from a cron service in a map
    - force https in redirects
    
- future work:
    - block url-shortening for malicious sites, if not possible, hook up site checker for verification
    - create an index.html site
    - use https in general (server, db)

### To run this simple project
- create a database
- create a config file or use env




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
