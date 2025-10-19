# Go CRUD Task with MongoDB

This project is a Go web application using the [Fiber](https://github.com/gofiber/fiber) framework and MongoDB for data storage. It includes live-reloading for development using [Air](https://github.com/air-verse/air).

## Project Structure

- `main.go`: Main application entry point.
- `go.mod` / `go.sum`: Go module files for dependency management.
- `air.toml`: Configuration file for Air live-reloading.

## Key Tools & Dependencies

### Fiber

- **Install:**
  ```bash
  go get github.com/gofiber/fiber/v2
  ```
- **Purpose:** Fiber is a fast, Express-inspired web framework for Go. It simplifies building REST APIs and web servers.
- **Usage:** Import it in your code:
  ```go
  import "github.com/gofiber/fiber/v2"
  ```
  Example minimal server:
  ```go
  app := fiber.New()
  app.Get("/", func(c *fiber.Ctx) error {
      return c.SendString("Hello, Fiber!")
  })
  app.Listen(":4000")
  ```

### Air

- **Install:**
  ```bash
  go install github.com/air-verse/air@latest
  ```
- **Purpose:** Air is a live-reloading tool for Go applications. It watches your source files and automatically rebuilds/restarts your app when you make changes.
- **Usage:**
  1. Make sure `$HOME/go/bin` is in your PATH:
     ```bash
     echo 'export PATH="$HOME/go/bin:$PATH"' >> ~/.bashrc
     source ~/.bashrc
     ```
  2. Start Air in your project directory:
     ```bash
     air
     ```
  3. Air uses `air.toml` for configuration. By default, it watches `.go` files and runs your app on changes.

### MongoDB

- **Purpose:** Used as the database for CRUD operations. You need a running MongoDB instance and the Go MongoDB driver (add to `go.mod` as needed).

## Common Issues & Solutions

- **Port already in use:** If you see `failed to listen: listen tcp4 :4000: bind: address already in use`, another process is using the port. Find and kill it:
  ```bash
  sudo lsof -i :4000
  kill <PID>
  ```
- **Air not found:** Make sure Air is installed and `$HOME/go/bin` is in your PATH.

## How Dependencies Work

- `go get` adds libraries to your project and updates `go.mod`/`go.sum`.
- `go install` installs CLI tools (like Air) to your Go bin directory.
- `air.toml` configures how Air watches and rebuilds your app.

## Quick Start

1. Clone the repo and enter the directory.
2. Install dependencies:
   ```bash
   go get github.com/gofiber/fiber/v2
   go install github.com/air-verse/air@latest
   ```
3. Make sure your PATH is set:
   ```bash
   export PATH="$HOME/go/bin:$PATH"
   ```
4. Start live-reloading:
   ```bash
   air
   ```
5. Open your browser to `http://localhost:4000` (or the port you set).

## References

- [Fiber Documentation](https://docs.gofiber.io/)
- [Air Documentation](https://github.com/air-verse/air)
- [Go Modules](https://blog.golang.org/using-go-modules)

---

**This README documents the setup and usage of Fiber, Air, and other essentials for future reference.**
