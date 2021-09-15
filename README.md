## Start project with Docker

```
docker run --rm -v "$PWD":/usr/src/myapp -w /usr/src/myapp  -p 4567:4567 golang:1.17 ./miniapi

```