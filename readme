# Storage Metrics to InfluxDB Injection Tool

This tool is designed to collect storage metrics from the local system and inject them into an InfluxDB database. It retrieves information about the file systems, such as the file system name, size, used space, available space, and usage percentage, and sends this data to an InfluxDB instance.

## Overview

The tool is written in Go and uses the `influxdb-client-go` library to interact with the InfluxDB database. It reads configuration settings from a JSON file, including the InfluxDB host, port, organization, bucket, and authentication token. The tool then executes the `df` command to retrieve storage metrics and parses the output to extract the relevant information. Finally, it connects to the InfluxDB instance and writes the storage metrics to the specified bucket.

## Prerequisites

- Go programming language (version 1.16 or later)
- An InfluxDB instance (version 2.x)

## Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/your-repo/storage-metrics-influxdb.git
   ```

2. Navigate to the project directory:

   ```bash
   cd storage-metrics-influxdb
   ```

3. Build the project:

   ```bash
   go build
   ```

## Configuration

Before running the tool, you need to create a configuration file named `config.json` in the `config` directory. The file should contain the following fields:

```json
{
  "InfluxHost": "your-influxdb-host",
  "InfluxPort": 8086,
  "InfluxHttpSchema": "http",
  "InfluxToken": "your-influxdb-token",
  "InfluxOrg": "your-influxdb-organization",
  "InfluxBucket": "XXXXXXXXXXXXXXXXXXXX"
}
```

Replace the placeholders with the appropriate values for your InfluxDB instance.

## Usage

Run the tool:

```bash
./storage-metrics-influxdb
```

The tool will output the current time and indicate that it's trying to connect to the InfluxDB instance. If the connection is successful, it will print a message indicating that the connection was successful.

The tool will then retrieve the storage metrics and write them to the specified InfluxDB bucket.

Once the data is written, the tool will print a message indicating that the InfluxDB was updated successfully and flush any remaining data.

## Contributing

If you find any issues or have suggestions for improvements, feel free to open an issue or submit a pull request.

## License

This project is licensed under the MIT License.