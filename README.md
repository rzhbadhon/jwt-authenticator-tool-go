
# Go-JWT Authenticator Tool 🚀

<p align="center">
  A sleek, modern, and performant tool for decoding and verifying JSON Web Tokens (JWTs), powered by a Go backend and a responsive frontend.
</p>

<p align="center">
  <img src="https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white" alt="Go" />
  <img src="https://img.shields.io/badge/Tailwind_CSS-38B2AC?style=for-the-badge&logo=tailwind-css&logoColor=white" alt="Tailwind CSS" />
  <img src="https://img.shields.io/badge/License-MIT-blue.svg?style=for-the-badge" alt="License: MIT" />
</p>

---

<p align="center">
  <a href="https://go-jwt-tool.github.io/"><strong>View Live Demo »</strong></a>
</p>

<br/>

<img src="https://i.ibb.co.com/kg6mScXk/Screenshot-2025-10-07-000436.png" alt="Screenshot-2025-10-07-000436" border="0">

## ✨ Features

* **Instant JWT Decoding**: Paste a token to instantly see the Header and Payload.
* [cite_start]**HS256 Signature Verification**: Verify token integrity against a secret key[cite: 2].
* **Modern Techy UI**: A beautiful, dark-themed interface with neon highlights and responsive design.
* **Blazing Fast Backend**: Built with Go for maximum performance and efficiency.
* **Pure Vanilla JS**: No heavy frameworks, ensuring a lightweight and fast frontend.

---

## 🛠️ Tech Stack

* **Backend**: Go (Golang)
* **Frontend**: HTML, Tailwind CSS, Vanilla JavaScript
* **Icons**: Feather Icons

---

## ⚙️ Getting Started (Running Locally)

To get a local copy up and running, follow these simple steps.

### Prerequisites

* Go installed on your machine.
* A modern web browser.

### Installation & Running

1.  **Clone the repository:**
    ```sh
    git clone [https://github.com/rzhbadhon/jwt-tool.git](https://github.com/rzhbadhon/jwt-tool.git)
    cd jwt-tool
    ```
2.  **Run the Go Backend:**
    Open a terminal in the project root and run:
    ```sh
    go run main.go
    ```
    [cite_start]Your server will start and listen on `http://localhost:8080`[cite: 1].

3.  **Open the Frontend:**
    Navigate to the `frontend` directory and open the `index.html` file in your browser.

---

## 📁 Project Structure

The project is organized with a clear separation between the frontend and the backend logic.

````

jwt-tool/
├── frontend/
│   └── index.html         \# The single-file frontend application
├── handlers/
│   ├── cors.go            \# CORS middleware
│   └── create\_jwt.go      \# Main JWT handler
├── models/
│   └── jwt.go             \# Structs for JSON requests/responses
├── util/
│   ├── base64.go          \# Base64 encoding utilities
│   ├── decode.go          \# JWT decoding logic
│   └── verify.go          \# Signature verification logic
├── go.mod
└── main.go                \# Main application entrypoint

```
````


## 👤 Contact

**Razibul Hasan Badhon**

* **GitHub**: [@rzhbadhon](https://github.com/rzhbadhon)
* **LinkedIn**: [/in/rzhbadhon](https://linkedin.com/in/rzhbadhon)
* **Instagram**: [@ig_rzhbadhon](https://www.instagram.com/ig_rzhbadhon)

