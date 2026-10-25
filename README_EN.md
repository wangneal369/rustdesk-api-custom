# RustDesk API Server (Custom Build)

API server for RustDesk remote desktop: users, devices, address books and an admin panel. The built admin frontend is served from `resources/admin/` under `/console/`.

## Build and deploy

See [docs/custom-build.md](docs/custom-build.md).

```shell
docker build -f Dockerfile.custom -t rustdesk-api-custom:local .
```

## Configuration

The config file is `conf/config.yaml`. The web client is disabled by default (`app.web-client: 0`).

基于 lejianwen/rustdesk-api（MIT License）二次开发
