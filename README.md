# Hook

![go-mod](https://img.shields.io/github/go-mod/go-version/cyan903/c-share) ![go-report](https://goreportcard.com/badge/github.com/cyan903/c-share) ![last-commit](https://img.shields.io/github/last-commit/cyan-store/Shop)

The stripe webhook for cyan-store. This is required for orders to process and payments to complete. Requires an SQL server along with redis.

```sh
$ make build # build for production
$ make dev # run for development
$ make format # format & lint code
$ make listen # listen to webhook with stripe-cli
```

## Features

- Insert orders from webhook
- Read cached incomplete orders
- Email error reporting
- Email successful purchases
- Validate order information

## License

[MIT](LICENSE)
