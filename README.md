# redirectSSRF

A simple Go web server to redirect a SSRF request back to localhost if target application is denying certain hostname, such as localhost:

http://example.com:8000/redirect?url=http://localhost/

To run:

`go run main.go`
