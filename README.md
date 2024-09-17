# UDISE API

API to fetch schools details (from a big query table) based on provided udise id.

## Pre-requisite

- Set your GOOGLE_APPLICATION_CREDENTIALS to a GCP service account key.

```bash
export GOOGLE_APPLICATION_CREDENTIALS=[project-name-xyz].json
```

## Building

From the root of the source tree, run:

```bash
make build/api
```

## Using

```bash
make run # runs server
curl localhost:4000/v1/schools/[udise-id] # pings endpoint
```

## Project Structure

- `bin` contains compiled app binaries, ready for deployment to a production server.
- `cmd/api` contains app-specific code ~ running server, reading & writing requests etc.
- [internal](https://go.dev/doc/modules/layout#package-or-command-with-supporting-packages) contains adjunct packages used by the api ~ custom queries, data validation etc.
- `remote` contains config files & setup scripts for our production server.
