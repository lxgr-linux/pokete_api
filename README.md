# Pokete API

An rest API to get (Pokete, Attack etc.) data from [Pokete](https://github.com/lxgr-linux/pokete).

## Building
```shell
go build
```

## Usage
To lauch the server:
```shell
./pokete_api
```

Extended usage:
```
Downloads data and starts server
Usage: ./pokete_api [ARG1] [ARG2] ...
Options:
	--no-download    : Does not download the data
	--update         : Just downloads the data
	--port <port>    : The port the server will run on, default is 8000
	--help           : Shows this dialog
```

## Accessing data
The API provides access to Poketes, types and attacks. They can be reached via:

```
host:port/cathegory/subcathegory
```

So for example to get data about all Poketes on your localhost and the standart port:
```
localhost:8000/poketes
```

To just see data about `steini`:
```
localhost:8000/poketes/steini
```
