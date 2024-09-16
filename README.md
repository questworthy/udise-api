# udise-api

## Project Structure

- `bin` contains compiled app binaries, ready for deployment to a production server.
- `cmd/api` contains app-specific code ~ running server, reading & writing requests etc.
- [internal](https://go.dev/doc/modules/layout#package-or-command-with-supporting-packages) contains adjunct packages used by the api ~ custom queries, data validation etc.
- `remote` contains config files & setup scripts for our production server.
