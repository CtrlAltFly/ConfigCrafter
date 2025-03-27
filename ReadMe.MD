ConfigCrafter

Overview

ConfigCrafter is a simple and efficient tool that converts between YAML and Properties file formats. It provides both a CLI and a web API to handle conversions seamlessly.

Features

Convert YAML (.yaml/.yml) to Properties (.properties)

Convert Properties (.properties) to YAML (.yaml)

REST API for easy integration

Docker support for deployment

Built with Go for speed and reliability

Installation

Prerequisites

Go 1.21+

Docker (optional for containerized deployment)

Build from Source

git clone https://github.com/your-repo/configcrafter.git
cd configcrafter
go build -o configcrafter

Run as a Docker Container

docker build -t configcrafter .
docker run -p 8080:8080 configcrafter

Usage

CLI

You can convert a file using the CLI by running:

./configcrafter -input config.yaml -output config.properties

Web API

Run the server:

./configcrafter

Send a file for conversion:

curl -X POST -F "file=@config.yaml" http://localhost:8080/convert -o config.properties

API Endpoints

POST /convert - Accepts a YAML or Properties file and returns the converted format.

License

This project is licensed under the MIT License.

Contributing

Feel free to submit issues or pull requests to improve the tool!

Author

CtrlAltFly