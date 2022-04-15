# Go Server Client
This project is a client server in Go. This project was done during a Go course with the objective of doing a client-server communication. The client sends an image using TCP and the server processes the image (here an RGB to gray scale conversion) and sends it back. Unfortunately, I did not find the final version of the project. So this project is only an image processing script and an example of TCP server and client. Maybe I will do the project again.
# Development
## Project
```
- images/
- src/
```
`src/` is a folder that contains the files `RGBtoGray.go` which is a file that processes a given image to convert the RGB color to grayscale, and the files `TCPClient.go` and `TCPServer.go` which are the files that initiate a TCP connection. `images/` is a folder that contains the PNG images that will be processed by the RGB converter
## Install
```zsh
$ git clone https://github.com/theolemague/GoServerClient.git
...
$ cd GoServerClient
```
## Run
### RGB to Gray scale
To run the RGB to gray scale converter, you need to run `RGBtoGray.go` with the command `go run` (my version: go1.18.1) followed by the name of the image you want to use (the image must be placed in `images/` and use the PNG format)
```zsh
$ go run src/RGBtoGray.go
Default image is used
Number of goroutine : 3
Number of channel : 4
```
If you want to challenge the program a little more, you can choose an image from the [Hubble bank](https://hubblesite.org/resource-gallery/images).
```zsh
$ go run src/RGBtoGray.go hubble.go
Default image is used
Number of goroutine : 25
Number of channel : 26
```
### Client Server example
To start the client-server connection, you must first start the server, then the client.
```zsh
(terminal 1) $ go run src/TCPServer.go
(terminal 2) $ go run src/TCPClient.go
```
After starting the client, the connection is created, a message is sent and the connection is closed and the client is terminated.
```zsh
(terminal 1) $ go run src/TCPServer.go
Launching server...
Client connected from 127.0.0.1:53161
&{Yo Hello}
(terminal 2) $ go run src/TCPClient.go
&{Yo Hello back}
```