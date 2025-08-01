<img width="1378" height="1036" alt="image" src="https://github.com/user-attachments/assets/1f591cbe-7787-4516-afca-3e7acc99b99f" />
# GoTube

A self-hosted, Go-based video streaming application inspired by YouTube’s core functionality. GoTube enables secure video upload, on-the-fly processing, and modern, responsive playback.

---
## Screenshots


<img width="1378" height="1036" alt="image" src="https://github.com/user-attachments/assets/bc6b1003-f8a2-4c40-a898-cdf14097d595" />
<img width="1378" height="1036" alt="image" src="https://github.com/user-attachments/assets/c5c456e7-f7a5-469b-adb7-19f1016e5fab" />



---

## Overview

GoTube is designed to provide a straightforward, performant platform for video hosting and streaming. Written in Go, the application offers:

- **Secure uploads** via HTTP handlers  
- **Automated processing** (thumbnail generation, format conversion)  
- **Responsive UI** with modern CSS styling  
- **Pagination** and **administrative controls**

---

## Features

- **Video Upload**: Support for multipart form uploads with file size limits.  
- **FFmpeg Integration**: Automatic thumbnail extraction and video transcoding.  
- **Video Listing**: Grid layout, hover animations, duration badges.  
- **Playback**: HTML5 video element with controls and poster image.  
- **Administration**: Delete functionality and basic admin panel.  
- **Configuration**: All settings managed via `config.yaml`.  
- **Docker Ready**: Provided Dockerfile for containerized deployment.

---
## Project structure
```bash
go-tube/
├── cmd/web          # Application entrypoint and HTTP router
├── internal/        # Core business logic (handlers, services)
├── pages/           # HTML templates
├── static/          # CSS, images, client-side assets
├── uploads/         # User-uploaded video files
├── converted/       # Generated thumbnails and converted media
├── config.yaml      # Application configuration
├── users.yaml       # User credentials and roles
├── Dockerfile       # Docker image definition
└── README.md        # Project documentation
```

---

## Prerequisites

- Go **1.18** or higher  
- **FFmpeg** installed and available in `PATH`  
- (Optional) **Docker** and **Docker Compose** for containerized setup

---

## Build Docker image

1) Clone Repository

```sh
git clone https://github.com/sillkiw/go-tube.git
cd go-tube
```
2) Create Docker container
```sh
docker build -t <name-container> .  
```
3) Run container
```sh
docker run -p 8085:8085 <name-container>
```

## Frameworks
FFMpeg, MP4Box


