Created some native types from C header files.

This is to create Go code that mirrors the structure of the C code, and which
can be consumed by higher level interfaces

```bash
cd pkg/southbound/e2proxy/e2ctypes/
protoc -I. --go_out=. *.proto
```

```bash
docker run -it -v `pwd`:/go/src/github.com/onosproject/onos-ric -w /go/src/github.com/onosproject/onos-ric --entrypoint pkg/southbound/e2proxy/ctypes/internal-protos.sh onosproject/protoc-go:v0.6.4
```
