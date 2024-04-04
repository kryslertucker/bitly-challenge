# Bitly challenge solution

This solution requires Golang v1.22.1 or higher. I have included a Makefile with useful targets for running and development.

The code supports 4 log levels (DEBUG=-4,INFO=0, WARNING=4, and ERROR=8). It defaults to `INFO` level, but it can be configured by setting the `LOG_LEVEL` environment variable in the `.env` file. See `.env.example` for example values.

```bash
$ make run
```

Run in debug mode:
```bash
$ LOG_LEVEL=4 make run
```

To run the set of unit tests created:
```bash
$ make clean test
```

## Dependencies

All dependencies are listed in the `go.mod` file. I have mostly used libraries from the golang standard library, with the exception of `github.com/stretchr/testify` used in the unit tests and `github.com/joho/godotenv` to read environment variables.

## Design decisions

I decided to use a hash with the key being the combination of domain and the bitly hash and the value being the long URL, in order to quickly identify if the bitlink provided when iterating over the clicks belongs to a relevant encode or not. 

In order to log the final result in the expected format, I made the `click.Results` and `click.Result` structs implement the Stringer interface with a custom string representation where each object is separated by a comma and the key of each object is the long URL and the value is the amount of clicks. 