# Location Service

The Location Service provides an api for saving and retrieving the gps location of any entity.

It's uses [go-micro](https://github.com/micro/go-micro) for the microservice core and Hailo's [go-geoindex](https://github.com/hailocab/go-geoindex) for fast point tracking and K-Nearest queries. 

## Usage

Run the service

```
go run main.go 
```

Test

```
go run examples/client.go
```

Output

```
Saved entity: id:"id123" type:"runner" location:<latitude:51.516509 longitude:0.124615 timestamp:1425757925 > 
Read entity: id:"id123" type:"runner" location:<latitude:51.516509 longitude:0.124615 timestamp:1425757925 > 
Search results: [id:"id123" type:"runner" location:<latitude:51.516509 longitude:0.124615 timestamp:1425757925 > ]
```
