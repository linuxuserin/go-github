# go-github

A simple cli to search github users.

## Build and Run

```sh
go get github.com/linuxuserin/go-github

go-github -u username
```

## Manual

Query a user,

```sh
go-github -u username
```

Query a user with auth,

```sh
go-github -u username -x $token
```

Query more than one users,

```sh
go-github -u username1,username2,username3
```

Login to github,

```sh
go-github -x $token
```

    Usage: go-github [options]
    Options:
    -x, --auth string   Login to Github
    -h, --help          help message
    -u, --user string   Search Users

## License

- MIT
