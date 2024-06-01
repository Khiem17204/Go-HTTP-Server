# Go Basic HTTP Server
This is a simple HTTP server written in Go programming language from scratch, without using any third-party libraries or frameworks.

## Overview
The server listens for incoming HTTP requests on a specified port and handles the following routes:

GET /: Returns a simple "Hello, World!" message.
GET /echo/message=<message>: Returns the provided message in the response body(support gzip compression).
GET /files/<filename>: Reads the files in the system and return its content in the response
POST /files/<filename>: Sends a file to the server

## Prerequisites
To run this server, you need to have Go installed on your machine. You can download and install Go from the official website: https://golang.org/dl/
Getting Started

Clone this repository or create a new Go project.
Copy the main.go file into your project directory.
Open a terminal or command prompt and navigate to your project directory.
Run the server with the following command:

**./your_server.sh**

By default, the server will listen on http://localhost:4221. You can change the port by modifying the PORT constant in the server.go file.

## Usage
Once the server is running, you can send HTTP requests using tools like curl, a web browser, or a dedicated testing tool like Postman.

## Recognition
This is a challenge of [CodeCrafters](https://codecrafters.io/), feel free to check out others challenges!
