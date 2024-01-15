
# Word of Wisdom TCP Server

This project implements a TCP server that sends quotes from the "Word of Wisdom" collection after clients successfully complete a Proof of Work (PoW) challenge. This approach aims to provide basic protection against DDOS attacks.

## Features

1. **TCP Server**: Implements a basic TCP server using Go.
2. **Proof of Work**: Integrates a PoW mechanism to mitigate DDOS attacks.
3. **Quotes Delivery**: Sends a random quote from a predefined collection after successful PoW verification.
4. **Docker Support**: Includes Dockerfiles for both server and client, facilitating easy deployment and testing.

## Proof of Work Algorithm

The server uses the SHA-256 hashing algorithm for PoW. This algorithm was chosen due to its cryptographic security, widespread use (e.g., in Bitcoin), and efficiency in generating hashes. Clients must find a hash with a specified number of leading zeroes by altering their input, which computationally requires time and resources, deterring simple DDOS attacks.

## Running the Code

### Prerequisites

- Install [Go](https://golang.org/dl/) (version 1.18 or later)
- Install [Docker](https://docs.docker.com/get-docker/)

### Steps

1. **Clone the Repository**:
   ```sh
   git clone https://github.com/tech-sumit/go-pow
   cd go-pow
   ```

2. **Build and Run using Docker**:
   - To build the Docker image:
     ```sh
     docker build -t go-pow .
     ```
   - To run the container:
     ```sh
     docker run --rm go-pow
     ```

### Server and Client

- The server starts automatically within the Docker container.
- The client is also executed within the same container, solving the PoW challenge and receiving a quote.

# Test cases

```bash
 $ go test -v ./... -coverprofile=coverage.out 
 $ go tool cover -html=coverage.out -o coverage.html
 ```

## Development Advice Compliance

- **Thorough Testing**: Extensive testing of client-server communication has been conducted.
- **Connection Verification**: The client successfully establishes a connection with the server.
- **Timeouts and Signal Handling**: Implemented basic timeout and signal handling in the server.
- **Code Organization**: Code is organized into logical sections.
- **Logging**: Basic logging is implemented; for production, consider using a dedicated logging library.
- **Secure Nonce Generation**: Uses cryptographically secure methods for nonce generation.
- **Security Considerations**: Emphasizes security in the PoW implementation.
- **Test Coverage**: Includes basic tests; further test coverage is recommended.
- **Linting and Code Quality**: Code formatted according to Go standards; use a linter for further analysis.
- **Error Handling**: Consistent error-handling practices are followed.
- **Docker Optimization**: Dockerfiles are optimized for multi-stage builds.

