# Go URL Shortener

A simple URL shortener service built using Go. This project provides an API to shorten long URLs and redirect them to the original URL when accessed.

## Features
- Shorten long URLs into small, unique shortened URLs.
- Redirect to the original URL when the shortened URL is visited.

## Requirements
- Go 1.18 or higher

## Installation

1. Clone this repository:

    ```bash
    git clone https://github.com/PratikBhopi/go-url-shortener.git
    cd go-url-shortener
    ```
## Usage

1. Start the application:

    ```bash
    go run main.go
    ```

    This will start a web server on `localhost:3000` by default.

2. After starting the server, simply open your browser and go to:

    ```
    http://localhost:3000/shorturl
    ```

    This will automatically generate a shortened URL and redirect you to the original URL.

3. The shortened URL will be displayed on the browser, and it will redirect you to the original URL when accessed.

