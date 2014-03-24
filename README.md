llano
=====
![Build Status](https://travis-ci.org/MindTwister/llano.png)

[![GoDoc](https://godoc.org/github.com/MindTwister/llano?status.png)](http://godoc.org/github.com/MindTwister/llano)

The [llano pocket gopher](https://en.wikipedia.org/wiki/Llano_pocket_gopher).

Llano is a mock server, useful when testing other libraries during development.

See [godoc](https://godoc.org/github.com/MindTwister/llano) for library usage.

### To install the binary

    go get github.com/MindTwister/llano/llano

### Example usage

    llano -http=":2020" -body="OK"

### Command line flags

**-http** `string` the address on which to listen, default: **127.0.0.1:2020**

**-body** `string` Response for `/200` default: **OK**


