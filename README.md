# File Cleaner
The application runs a routine to delete files from S3. I created it to explore the fundamentals of the Go language.

## Run

Keeps dependencies running after execution
```bash
./scripts/start-app.sh --keep
```

Removes dependencies after execution
```bash
./scripts/start-app.sh
```

## Build

```bash
go build cmd/cleaner/main.go
```
