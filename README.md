# baldeweg/desk_storage

An app to administer shifts, staff and more.

## Repositories

- storage <https://github.com/abaldeweg/desk_storage> - Backend
- shift <https://github.com/abaldeweg/desk_shift> - UI

## Requirements

- [Go](https://go.dev/)

## Getting Started

First, you need to install [Go](https://go.dev/).

Download the project archive from the [git repository](https://github.com/abaldeweg/desk_storage).

Inside the project directory, you can build the app with the `go build` command. If you have [GoReleaser](https://goreleaser.com/) installed, instead run `goreleaser build --snapshot --rm-dist`.

Run the command `storage`. Depending on the OS you need to add a file extension.

The app will create files where you can edit the staff and their shifts.

## Storage

Create a `.env` file to define some settings.

```env
// .env

ENV=prod
STORAGE=gcp-bucket
FILE_PATH=.
GCP_BUCKET_NAME=name
GOOGLE_APPLICATION_CREDENTIALS=service-account-file.json
CORS_ALLOW_ORIGIN=http://localhost:8081
```

- ENV - Set to `prod`, `dev` or `test`
- STORAGE - Can be `file` or `gcp-bucket`
- FILE_PATH - Path where to store the files, only for file storage
- GCP_BUCKET_NAME - If `gcp-bucket` was chosen as storage, then define the bucket name.
- GOOGLE_APPLICATION_CREDENTIALS - Key file, for auth and buckets
- CORS_ALLOW_ORIGIN - Allowed origins
