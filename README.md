[![GoDoc](http://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)](http://godoc.org/github.com/gigovich/queryqute)
[![Go Report Card](https://goreportcard.com/badge/github.com/gigovich/lazysetup)](https://goreportcard.com/report/github.com/gigovich/queryqute)
# QueryCute
Simple Go query tool for PostgreSQL

# Test
Install test dependencies:
```
go get -t .
```

Create test user and database:
```
sudo -u postgres psql -c "CREATE USER test WITH PASSWORD 'test123'"
sudo -u postgres psql -c "CREATE DATABASE test WITH OWNER test"
```

Copy create `.env` file from `.env.example` and run tests:
```
cp .env.example .env
source .env
go test -v .
```
