Go-JWT Authenticator Tool 🚀
<p align="center">
A sleek, modern, and performant tool for decoding and verifying JSON Web Tokens (JWTs), powered by a Go backend and a responsive frontend.
</p>

<p align="center">
<img src="https://www.google.com/search?q=https://img.shields.io/badge/Go-00ADD8%3Fstyle%3Dfor-the-badge%26logo%3Dgo%26logoColor%3Dwhite" alt="Go" />
<img src="https://www.google.com/search?q=https://img.shields.io/badge/Tailwind_CSS-38B2AC%3Fstyle%3Dfor-the-badge%26logo%3Dtailwind-css%26logoColor%3Dwhite" alt="Tailwind CSS" />
<img src="https://www.google.com/search?q=https://img.shields.io/badge/License-MIT-blue.svg%3Fstyle%3Dfor-the-badge" alt="License: MIT" />
</p>

<p align="center">
<a href="https://www.google.com/search?q=https://rzhbadhon.github.io/jwt-tool/"><strong>View Live Demo »</strong></a>
</p>

<p align="center">
<img src="https://www.google.com/search?q=https://i.imgur.com/uG9kSza.png" alt="Go-JWT Tool Screenshot" width="800"/>
</p>

✨ Features
Instant JWT Decoding: Paste a token to instantly see the Header and Payload.

HS256 Signature Verification: Verify token integrity against a secret key.

Modern Techy UI: A beautiful, dark-themed interface with neon highlights and responsive design.

Blazing Fast Backend: Built with Go for maximum performance and efficiency.

Pure Vanilla JS: No heavy frameworks, ensuring a lightweight and fast frontend.

🛠️ Tech Stack
Backend: Go (Golang)

Frontend: HTML, Tailwind CSS, Vanilla JavaScript

Icons: Feather Icons

⚙️ Getting Started (Running Locally)
To get a local copy up and running, follow these simple steps.

Prerequisites
Go installed on your machine.

A modern web browser.

Installation & Running
Clone the repository:

git clone [https://github.com/rzhbadhon/jwt-tool.git](https://github.com/rzhbadhon/jwt-tool.git)
cd jwt-tool

Run the Go Backend:
Open a terminal in the project root and run:

go run main.go

Your server will be running at http://localhost:8080.

Open the Frontend:
Navigate to the frontend directory and open the index.html file in your browser.

Note: You may need to enable CORS in the backend to allow requests from file:///. The README.md from our previous conversation has instructions on how to do this.

📁 Project Structure
jwt-tool/
├── frontend/
│   └── index.html         # The single-file frontend application
├── handlers/
│   ├── cors.go            # CORS middleware
│   └── create_jwt.go      # Main JWT handler
├── models/
│   └── jwt.go             # Structs for JSON requests/responses
├── util/
│   ├── base64.go          # Base64 encoding utilities
│   ├── decode.go          # JWT decoding logic
│   └── verify.go          # Signature verification logic
├── go.mod
└── main.go                # Main application entrypoint

👤 Contact
Razibul Hasan Badhon

GitHub: @rzhbadhon

LinkedIn: /in/rzhbadhon

Instagram: @ig_rzhbadhon

Facebook: Razibul Hasan