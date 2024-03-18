# Multithreading-Challenge-Go

### This repository is a Multithreading challenge developed in Golang as a partial assessment for the completion of the Postgraduate Degree in Golang

### In this work, concepts such as the use of goroutines, HTTP requests and Channels

The requirements are:

1. Do a request to 2 APIS:
- https://brasilapi.com.br/api/cep/v1/01153000 + cep
- http://viacep.com.br/ws/" + cep + "/json/

2. Adopt the API that delivers the fastest response and discard the slower one.

3. The result of the request must be printed on the command line with the address data and the API that provides it

4. The response time is 1 second, otherwise the timeout message must be printed