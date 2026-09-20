# go-todos-api

Sample DockOps **instance-project**: a Go todos API plus MariaDB. Same contract as [dockops-instance](https://github.com/Dockops-Team/dockops-instance); this repo fills in the stub.

| Template (`dockops-instance`) | This sample |
|-------------------------------|-------------|
| `docker/webservice` = nginx identity page | `docker/webservice` = Go app image |
| No sidecars (commented examples) | `db` sidecar, config in `docker/db/` |
| `app/` absent | `app/` = Gin todos API |

Keep service **`webservice`** on **8080**. Extra services join `webservice` only, never `dockops`.

```bash
docker network create dockops
cp .env.dist .env
cp docker-compose.yml.dist docker-compose.yml
docker compose up --build
```

DockOps clones a branch, writes `.env` and `docker-compose.override.yml`, and copies `docker-compose.yml.dist` if `docker-compose.yml` is missing.

**Advanced env** on create is free-form `KEY=value`. Those lines go into `.env` and `env_file: .env` is attached to every service. They override `${VAR:-default}` in this compose file (`APP_ENV`, `GIN_MODE`, `DB_*`, `MYSQL_ROOT_PASSWORD`, …). Identity keys (`INSTANCE_*`, `ID`, `NAME`, `HOSTNAME`, `IP`) are reserved. Other instance-projects declare their own keys.

## Layout

```
docker-compose.yml.dist          # contract + this sample's db sidecar
docker-compose.override.yml.dist
.env.dist
docker/webservice/Dockerfile     # replaces the template nginx stub
docker/db/my.cnf                 # MariaDB sidecar
app/                             # Go source (listens on 8080)
```

## Swagger

```
swag init
```
