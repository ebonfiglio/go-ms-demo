# Go DB Demo

A **demo Go project** for learning how to build simple **console and web CRUD applications**, using:

- [`sqlx`](https://github.com/jmoiron/sqlx) for SQL database access with struct mapping
- [`gin`](https://github.com/gin-gonic/gin) for lightweight HTTP web framework
- PostgreSQL as the database

---

## What this project is

A **personal learning project** to explore:

- Structuring a Go application with `internal` packages
- Using `sqlx` for easy, type-safe database interactions
- Building a simple **console CLI** to manage users and organizations
- Adding a basic **REST API** using `gin` to serve the same data
- Managing database migrations with [`golang-migrate`](https://github.com/golang-migrate/migrate)
- Environment-based configuration management

---

## Configuration

The application uses environment variables for configuration, with support for `.env` files for local development.

### Setup

1. Copy the example environment file:
   ```bash
   cp .env.example .env
   ```

2. Edit `.env` with your database and server settings:
   ```bash
   # Database Configuration  
   # Set DB_HOST to your PostgreSQL VM's Tailscale IP
   DB_HOST=localhost
   DB_PORT=5432
   DB_USER=postgres
   DB_PASSWORD=your_password
   DB_NAME=go-db-demo
   DB_SSL_MODE=disable

   # Server Configuration
   # Set SERVER_HOST to your web app server's Tailscale IP (0.0.0.0 for all interfaces)
   SERVER_PORT=8080
   SERVER_HOST=0.0.0.0
   ```

### Configuration Options

| Variable | Description | Default |
|----------|-------------|---------|
| `DB_HOST` | Database host (use Tailscale IP for separate VM) | `localhost` |
| `DB_PORT` | Database port | `5432` |
| `DB_USER` | Database username | `postgres` |
| `DB_PASSWORD` | Database password | `admin` |
| `DB_NAME` | Database name | `go-db-demo` |
| `DB_SSL_MODE` | SSL mode for PostgreSQL | `disable` |
| `SERVER_PORT` | Web server port | `8080` |
| `SERVER_HOST` | Web server host (use 0.0.0.0 or Tailscale IP) | `localhost` |

The application will automatically load the `.env` file if present, otherwise it will use environment variables or the default values.

### Deployment Configuration

For deployment with separate VMs connected via Tailscale VPN, you'll need to configure the following secrets in your GitHub repository:

| Secret | Description |
|--------|-------------|
| `TS_HOST` | Tailscale IP of your web application server |
| `DB_HOST` | Tailscale IP of your PostgreSQL database server |
| `DEPLOY_SSH_KEY` | SSH private key for deployment |
| `DB_USER` | Database username |
| `DB_PASSWORD` | Database password |
| `DB_NAME` | Database name |

**Note**: The deployment workflow assumes that PostgreSQL is already set up on the database VM with the database and user created. The workflow will automatically run database migrations, but database and user creation must be handled separately on the database server.

---
