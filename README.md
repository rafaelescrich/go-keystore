# Go Keystore

Keystore written in golang to create and store safely your keys.

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development, testing or using purposes.

### Prerequisites

* Go version 1.5 at least
* Some linux distro with make

### Installing

A step by step series of examples that tell you have to get a development env running

Clone the project

```
git clone git@github.com:rafaelescrich/go-keystore.git $GOPATH/src/github.com/rafaelescrich/go-keystore
```

Build bin with make tool

```
make
```
Then if everything runned smoothly you should have a binary
To run it just type

```
./go-keystore
```

## Running the tests

make test

## Built With

* [IShell](https://github.com/abiosoft/ishell) - Library for creating interactive cli applications.
* [BoltDB](https://github.com/boltdb/bolt) - An embedded key/value database for Go.
* [Argon2](https://github.com/golang/crypto/argon2) - Go supplementary cryptography libraries
* [Fastrand](https://github.com/NebulousLabs/fastrand) - 10x faster than crypto/rand

## Author

* **Rafael Escrich** - [github.com/rafaelescrich](https://github.com/rafaelescrich)

## License

This project is licensed under the GPL v2 License - see the [LICENSE.md](LICENSE.md) file for details