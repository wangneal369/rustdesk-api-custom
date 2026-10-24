# 自定义版本构建说明

本文说明如何构建本定制版 RustDesk API 服务端镜像。镜像内包含 Go 后端与管理后台前端，不依赖上游发布的镜像。

## 默认配置要点

- `app.web-client: 0`：网页客户端及访客分享功能默认关闭（接口不注册、静态文件不挂载）。
- `admin.title`：管理后台标题，默认 `RustDesk 管理后台`。
- `rustdesk.id-server` / `relay-server` / `api-server`：为占位符 `your.server.ip`，部署前请改为自己的地址。
- 服务器命令（`/api/admin/rustdesk/*`）与 `shareByWebClient` 接口在当前版本中不注册。

配置文件位于 `conf/config.yaml`，镜像构建时会打包进去。若需要在部署时覆盖，请挂载外部文件，而不是修改镜像内容。

## 构建步骤

### 1. 构建管理后台前端

在前端仓库 `rustdesk-api-web-custom` 中执行：

```bash
cd E:/projects/rustProjects/RustDesk/rustdesk-api-web-custom
npm ci
npm run build
```

产物位于该仓库的 `dist/` 目录。

### 2. 将前端产物复制到后端 `resources/admin`

`resources/admin` 已在后端 `.gitignore` 中，不会被提交。

```bash
mkdir -p E:/projects/rustProjects/RustDesk/rustdesk-api-custom/resources/admin
cp -r E:/projects/rustProjects/RustDesk/rustdesk-api-web-custom/dist/. E:/projects/rustProjects/RustDesk/rustdesk-api-custom/resources/admin/
```

若 `resources/admin` 缺失，`docker build` 会在 `COPY resources/admin/` 步骤失败，这是预期行为。

### 3. 构建镜像

在后端仓库根目录执行：

```bash
cd E:/projects/rustProjects/RustDesk/rustdesk-api-custom
docker build -f Dockerfile.custom -t rustdesk-api-custom:local .
```

`Dockerfile.custom` 为多阶段构建：第一阶段在 `golang:1.23-alpine` 中以 CGO 编译 `apimain`（sqlite 驱动需要），第二阶段基于 `alpine` 组装运行环境。构建过程需要能访问 Go 模块代理。

### 4. 运行

```bash
docker run -d --name rustdesk-api -p 21114:21114 -v rustdesk-api-data:/app/data rustdesk-api-custom:local
```

- 数据目录挂载在 `/app/data`（sqlite 数据库位于其中）。
- `rustdesk.key-file` 默认为 `/data/id_ed25519.pub`，该路径不在镜像内，需要自行挂载 hbbs 生成的公钥文件，或修改该配置项。
- 如需使用外部配置文件，追加 `-v /绝对路径/config.yaml:/app/conf/config.yaml`。

## 注意事项

- 如需开启网页客户端，将 `app.web-client` 改为 `1` 后重新构建镜像，并确认 `resources/web` 存在。
- 本地 Windows 环境若 CGO 编译失败，可直接使用上述 Docker 构建，避免依赖本机 C 编译器。
