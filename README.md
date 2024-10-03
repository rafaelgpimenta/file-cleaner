# File Cleaner
The application runs a routine to delete files from S3. I created it to explore the fundamentals of the Go language.

## Build

```bash
go build cmd/cleaner/main.go
```

## Run

### Using your Go setup

Keeps dependencies running after execution
```bash
./scripts/start-app.sh --keep
```

Removes dependencies after execution
```bash
./scripts/start-app.sh
```

### Using docker compose

```bash
# starts cleaner application with dependencies
docker compose -f docker-compose/docker-compose.yml up -d
```

```bash
# removes containers
docker compose -f docker-compose/docker-compose.yml down
```
