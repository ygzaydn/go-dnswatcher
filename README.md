# Contents of /dns-watcher/dns-watcher/README.md

# DNS Watcher

DNS Watcher is a simple application that monitors DNS records for changes and provides a web interface for viewing the current status.

## Features

- Monitors DNS records for changes
- Web interface for real-time updates
- Configurable polling intervals

## Installation

1. Clone the repository:
   ```
   git clone https://github.com/yourusername/dns-watcher.git
   cd dns-watcher
   ```

2. Install dependencies:
   ```
   go mod tidy
   ```

## Configuration

Edit the `config.yaml` file to set your DNS server addresses and polling intervals.

## Usage

Run the application:
```
go run cmd/main.go
```

Visit `http://localhost:8080` in your browser to access the web interface.

## License

This project is licensed under the MIT License.