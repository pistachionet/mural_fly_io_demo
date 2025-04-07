# Fly.io Go Demo

This is a simple Go web server that runs on Fly.io and tells you what region is serving your request.

## 🚀 Quick Start

### 1. Install Fly CLI
```bash
curl -L https://fly.io/install.sh | sh
```

### 2. Login to Fly
```bash
fly auth login
```

### 3. Deploy the App
```bash
fly launch --name flyio-go-demo --region iad --no-deploy
fly deploy
```

### 4. Visit Your App
Fly will print a public URL where you can view the app.

## 🌍 What It Does
- Serves "Hello from Fly.io!"
- Displays the Fly.io region handling the request via `FLY_REGION`

## 🛠 Customize It
Edit `main.go` to add routes, API logic, or templates.

## 📦 Structure
- `main.go`: App logic
- `Dockerfile`: Container definition
- `fly.toml`: Fly.io app config

---
Enjoy deploying Go globally with zero ops!
