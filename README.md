# redirectSSRF

A simple Go web server to redirect a SSRF request to a specified host if target application is denying hostnames directly:

http://example.com:8000/redirect?url=http://localhost:8080/

To run:

`go run main.g0`
