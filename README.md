# URLv2

## Overview
URLv2 is a project designed to handle and process URLs efficiently. It provides tools for parsing, validating, and manipulating URLs in various formats. Additionally, it includes a URL shortening service implemented in Go.

## Features
- URL parsing and validation
- URL encoding and decoding
- Support for multiple protocols
- Easy-to-use API
- URL shortening service with a RESTful API

## Installation
Clone the repository:
```bash
git clone https://github.com/lenardjombo/URLv2.git
cd URLv2
```


### Go URL Shortener
Run the Go-based URL shortener service:
1. Ensure you have PostgreSQL installed and running.
2. Update the database connection string in the `models` package.
3. Start the server:
    ```bash
    go run main.go
    ```
4. Use the following endpoints:
    - **Generate Short URL**: `POST /v1/short` with JSON body `{"url": "https://example.com"}`
    - **Retrieve Original URL**: `GET /v1/short/{encoded_string}`

## Contributing
Contributions are welcome! Please fork the repository and submit a pull request.

## License
This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.