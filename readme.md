# GISPRIME

## Description

A simple web application written in go to work with my labs. 

Expose a REST http server to check if a number is prime. 

Also has a healthcheck endpoint

## Usage

- Clone this project and enter inside

- Build the image

```
docker build -t gisprime .
```

- Run the image 
```
docker run --rm -p 8000:8000 gisprime
```

- Make a request putting in the path the number you want to check if is prime
```
curl -s localhost:8000/31            
> {"isPrime":true}
```
