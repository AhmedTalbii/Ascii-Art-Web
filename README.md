# ASCII Art Web

A web-based interface for converting text into ASCII art banners, built with Go and served via a clean browser interface.

## 🚀 Features

- Multiple ASCII art styles: Standard, Shadow, ThinkerToy.
- Real-time generation: input text and preview instantly.
- Download your output as a `.txt` file.
- Error handling with custom 404 and 500 pages.
- Responsive UI built with HTML, CSS, and JavaScript.

## 🧩 Project Structure

```
Ascii‑Art‑Web/
├── ascii-art/         # Banner font templates (Standard, Shadow, ThinkerToy)
├── handlers/          # HTTP route and processing logic
├── server/            # Server setup and routing
├── static/            # CSS and JS files
├── templates/         # HTML layouts, index, and error pages
├── go.mod             # Go module file
└── main.go            # Entry point for the Go web server
```

## 🛠️ Prerequisites

- Go 1.18 or newer
- Git (for cloning)

## 🏃‍♂️ Running Locally

```bash
git clone https://github.com/AhmedTalbii/Ascii-Art-Web.git
cd Ascii-Art-Web
go run main.go
```

Then open your browser and go to `http://localhost:3000`.

## 🎨 Usage

1. Access the app in your browser.
2. Select a banner style.
3. Enter the text you want to convert.
4. Click **Generate**, preview the ASCII art, and download as `.txt`.
