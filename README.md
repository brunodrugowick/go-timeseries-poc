# go-timeseries-poc

An attempt on the same PoC that I worked on recently, but now in Golang. The plan is to use:

- [ ] Tests
- [X] Makefiles
- [X] Only native http stuff for the server
- [ ] Only native code to read environment variables/config
  - [X] Read properties from file
  - [ ] Override with values from environment
- [X] `sqlc` to generate database code (my exception to the "all native stuff" because it's a code generator, not a library)
- [ ] Golang templates to build pages
- [ ] ...

## Development

Make sure you can run these in your terminal:

- make
- docker (configured to run with you user)
- go

### Quick start

To quick start running this project:

```commandline
make local-db run
```

which creates a new postgres database and initialize the schema, then runs the application.

### Database code

All code that access database is generated with [`sqlc`](https://sqlc.dev/). This is not a library, it's a code generator, meaning you can inspect the code afterwards.

To generate the whole `pkg/infrastructure/database` package based on the `pkg/infrastructure/sqlc.yaml` file, run:

```commandline
make generate-db-code
```

## Reusable packages

Each package has (or will have) its own README.md. The following is a list of the current reusable packages:

- `config-reader`
- `server`
