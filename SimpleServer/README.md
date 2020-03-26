# SimpleWebServer
A simple web server in go programming language running on a docker container.

## Usage
```sh
cd <path/to/the/repository>
cd SimpleServer
docker build -t <username>/<image-name> .
docker run -p 8080:8080 <username>/<image-name>
```

Open your browser and goto localhost:8080
