# SimpleNodeServer
A simple web server in node programming language which is being run on a docker container.

## Usage
```sh
cd <path/to/the/repository>
cd SimpleNodeServer
docker build -t <username>/<image-name> .
docker run -p 8081:8081 <username>/<image-name>
```

Open your browser and goto localhost:8081

