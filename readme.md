# Go URL Shortener

A simple URL shortener service written in Go.

## Overview

This project is a URL shortening service similar to bit.ly or tinyurl.com, allowing users to create shorter aliases for long URLs. It uses Go, the Gorilla Mux router for HTTP routing, and the Bun ORM for database interactions.

## Getting Started

### Prerequisites

- Go 1.22.4 or higher
- PostgreSQL

### Installation

1. Clone the repository:

    ```sh
    git clone https://github.com/madushas/go-url-shortner.git
    cd go-url-shortner
    ```

2. Copy the `.example.env` file to `.env` and update the environment variables according to your setup:

    ```sh
    cp .example.env .env
    ```

3. Install the dependencies:

    ```sh
    go mod tidy
    ```

4. Run the application:

    ```sh
    go run cmd/urlshortner/main.go
    ```

The application will be available at `http://localhost:8080`.

## Usage

### Creating a Short URL

To create a short URL, send a POST request to the `/shorten` endpoint with the `url` parameter:

```json
{
  "url": "https://example.com",
  "customURL": "example",
  "expiresAt": "2023-12-31T23:59:59Z"
}
```

The `customURL` and `expiresAt` parameters are optional. If `customURL` is not provided, a random short URL will be generated. If `expiresAt` is not provided, the short URL will never expire.

### Redirecting to the Original URL

Access the short URL generated in a web browser or send a GET request to the short URL. The application will redirect you to the original URL.

### Contributing

Contributions are welcome! Feel free to open an issue or submit a pull request if you find a bug or want to add a new feature.

## License

Distributed under the MIT License. See [LICENSE](LICENSE) for more information.
