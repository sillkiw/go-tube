# go-tube

Lightweight YouTube‑style app in Go — **upload → convert → stream** locally. It ships a minimal web UI, stores configuration in YAML, and includes a Dockerfile for containerized deployment.

---

## Features

- Secure file uploads with optional **TLS** (HTTP can be enabled too)
- Optional **password protection** (admins/users) for uploads and viewing
- **Video conversion** to multiple resolutions and formats (DASH/WebM)
- Simple static **HTML/CSS/JS** web interface

---

## Tech stack

- **Go** (standard library for HTTP server & templates)
- **HTML / CSS / JavaScript** (frontend)
- **FFmpeg** and **MP4Box (GPAC)** for media processing
- **Docker** (containerization)


---

## Repository layout

```
.
├── cmd/
│   └── web/            # Application entry point (main package)
├── internal/           # Internal packages (handlers, services, utils)
├── pages/              # HTML templates / pages for UI
├── static/             # Static assets (JS, CSS, images)
├── uploads/            # Upload target (created if missing)
├── Dockerfile          # Container build recipe
├── config.yaml         # Main application configuration
├── users.yaml          # Users & roles (if auth is enabled)
├── go.mod / go.sum     # Go modules
└── README.md           # Project documentation
```


---

## Prerequisites

- **Go** 1.21+
- **FFmpeg**
- **MP4Box (GPAC)**

Debian/Ubuntu:

```bash
sudo apt update
sudo apt install -y ffmpeg gpac
```

macOS (Homebrew):

```bash
brew install ffmpeg gpac
```

---

## Configuration

All runtime settings live in `config.yaml`.

### Users & roles

`users.yaml` defines admin/user credentials and roles used for gated uploads or viewing. Keep this file outside the image when using Docker and **mount it** at runtime.

---

## Run locally

```bash
# 1) Fetch modules
go mod download

# 2) Configure the app
cp config.yaml config.local.yaml   # optional copy for local edits
# edit config.local.yaml and/or users.yaml

# 3) Start the server
go run ./cmd/web
# or build a binary
go build -o bin/go-tube ./cmd/web
./bin/go-tube
```

Visit:

```
http://<host>:<ServerPort>/
# or, if TLS is enabled
https://<host>:<ServerPortTLS>/
```

---

## Docker

Build from the included Dockerfile and run with mounted config and media folders:

```bash
# Build
docker build -t go-tube .

# Run (adjust ports to match your config)
docker run -p 8085:8085 go-tube  
```

---

