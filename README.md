# Go Auth Server
[![go report card][grc-badge]][grc-url]
[![license][license]](LICENSE)

Authentication server in Go.

### Install
Installing EdgeX Auth Server is trivial [`go get`](https://golang.org/cmd/go/):
```bash
go get github.com/drasko/edgex-auth
cd edgex-auth/cmd
go buld -o edgex-auth
./edgex-auth
```

### Features
- JWT
- User account creation
- MongoDB persistance
- Login and logout
- Token revocation

### License
[Apache License, version 2.0](LICENSE)

[grc-badge]: https://goreportcard.com/badge/github.com/drasko/edgex-auth
[grc-url]: https://goreportcard.com/report/github.com/drasko/edgex-auth
