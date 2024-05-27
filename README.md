# Dummy Sensor

Dummy Sensor Server is a simple TCP server application written in Go (Golang). It simulates receiving data from sensors and handles incoming connections to read the data.

## Features

- Listens for incoming connections on port 8080.
- Handles multiple concurrent connections using goroutines.
- Reads data sent by clients (simulated sensor data) and prints it to the console.
- Gracefully handles errors such as connection issues and EOF (end of file).

## Requirements

- Go (Golang) installed on your machine.

## Usage

1. Clone the repository

2. Navigate to the project directory:

```bash
cd DummySensor
```

3. Build the executable:

```bash
go build
```

4. Run the server:

```bash
./dummy-sensor-server
```

5. The server will start listening on port 8080. You can now connect to it using a TCP client.

## Configuration

- Port: By default, the server listens on port 8080. You can modify the port in the source code if needed.

## Contributing

Contributions are welcome! Feel free to submit issues or pull requests.
