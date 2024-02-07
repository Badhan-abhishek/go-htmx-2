<!-- Write README for this project -->

# Go HTMX Task Manager

Based on Go, HTMX, Gorm, tailwindcss & atlas for migrations. A simple project-task manager for personal use.
Create projects and tasks to those projects

# Setup

1. Clone the repository
2. Run `go mod tidy` to install the dependencies
3. Install atlas and go ($GOPATH should be set)
    ```bash
    go get github.com/pressly/atlas/cmd/atlas
    ```

    and

    ```bash
    go get github.com/cosmtrek/air`
    ```

4. Verify DB credentials in `cmd/models/db.go` & `atlas.hcl`
5. Run migrations
   ```bash
   atlas migrate apply --env gorm
   ```
6. Install node packages using pnpm (for tailwindcss)
   ```bash
   pnpm install
   ```
7. Run the server
   ```bash
   air # for development
   ```
8. Open the browser and go to `http://localhost:8080`
