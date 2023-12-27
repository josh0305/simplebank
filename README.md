# Simple Bank backend

This codebase constitutes a straightforward bank backend implemented in [Golang](https://go.dev/). It enables you to perform various operations, including user creation, account creation, and money transfers between accounts. All services are exposed in a RESTful API format, ensuring easy accessibility. Furthermore, this repository incorporates JSON Web Token (JWT) and Paseto to secure the claims exchanged between endpoints. The project encompasses CRUD functionality for SQL and adheres to the ACID rules for the transaction service. Enjoy exploring and experimenting with this codebase!

## Requirement Tool

- [Docker desktop](https://www.docker.com/products/docker-desktop)
- [Golang 1.21.5](https://golang.org/)
- [Migrate] (https://github.com/golang-migrate/migrate)
- [SQLc](https://docs.sqlc.dev/en/latest/overview/install.html)
- [Gomock](https://pkg.go.dev/github.com/stretchr/testify/mock)

## Setup command

- Start Postgres container

```bash
make postgres
```

- Generate CRUD SQL

```bash
make sqlc
```

- Generate Mock code for test

```bash
make mock
```

- Migrate DB

```bash
make migrateup
```

```bash
make migratedown
```

## Start Server

Before running the server, must run postgres server on port 5432 and migrate the DB!

```bash
make server
```

Command for test

```bash
make test
```

## Note

This repository draws inspiration from the [project](https://github.com/techschool/simplebank). However, certain sections of the code were outdated due to updates in some packages. Consequently, I've made the following modifications:

JWT Package: The original package has been deprecated, and in version 5, a different architecture for claims is utilized. I've implemented the JWT token maker for version 5, and it has successfully passed the unit tests.

Docker Compose File: In the original repository, a third-party package was employed to monitor the status of the Postgres container. However, a new feature called healthy check status is now available, and I have incorporated it into my version. If you're interested, you can find more details in the [link](https://docs.docker.com/compose/compose-file/05-services/#long-syntax).
