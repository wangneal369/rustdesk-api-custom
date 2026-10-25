# RustDesk API 服务端（定制版）

RustDesk 远程桌面的 API 服务端，提供用户、设备、地址簿与管理后台功能。管理后台前端构建后位于 `resources/admin/`，通过 `/console/` 访问。

## 构建与部署

参见 [docs/custom-build.md](docs/custom-build.md)。

```shell
docker build -f Dockerfile.custom -t rustdesk-api-custom:local .
```

## 配置

配置文件为 `conf/config.yaml`，默认关闭网页客户端（`app.web-client: 0`）。

基于 lejianwen/rustdesk-api（MIT License）二次开发
