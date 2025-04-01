# Key-Value Store (WIP)

This is an experimental Key-Value store which I am building while studying distributed systems. Currently it has 0 dependencies and I will try to keep it that way.

It is currently usable as an in-memory cache for simple usecases.

## Usage

### Start Server
```
// default 3000
./kvs -p 6379
```

### Operations
```
// set "foo" key to value "bar" without TTL (previous value will be replace if key already exists)
curl -X PUT localhost:3000/foo -d "bar"

// set "foo" to "bar" with 10s TTL
curl -X PUT localhost:3000/foo/10 -d "bar"

// returns value "bar" (404 if expired or does not exist)
curl localhost:3000/get/foo

// Set Integer values
curl -X PUT localhost:3000/counter -d "0"

// increments counter value by 5
curl localhost:3000/inc/counter/5

// decrements counter value by 1
curl localhost:3000/dec/counter

// delete counter
curl -X DELETE localhost:3000/delete/counter

// list all key-value pairs
curl localhost:3000/list

// flush database
curl -X DELETE localhost:3000/flush
```

## Roadmap

- [x] PUT/GET over http
- [x] expiration time
- [x] DELETE/LIST/FLUSH
- [x] integers and inc/dec
- [ ] Snapshots and logs
- [ ] replication
- [ ] explore protocols other than http