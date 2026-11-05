# Gator

Gator is a multi-user command-line RSS feed aggregator written in Go. You add RSS feeds, follow the ones you like, and Gator collects their posts into a PostgreSQL database so you can browse them from your terminal.

## Requirements

You'll need these installed:

- **Go** 1.23 or later: <https://go.dev/doc/install>
- **PostgreSQL** 15 or later: <https://www.postgresql.org/download/>

## Installation

Install the `gator` CLI with `go install`:

```bash
go install github.com/AtlasRegap/gator@latest
```

This builds the program and puts a `gator` binary in `~/go/bin`. Make sure that folder is on your `PATH`:

```bash
echo 'export PATH=$PATH:$HOME/go/bin' >> ~/.bashrc
source ~/.bashrc
```

## Setup

### 1. Create the database

Start Postgres and create a database called `gator`:

```bash
sudo service postgresql start      # Linux / WSL
sudo -u postgres psql -c "CREATE DATABASE gator;"
```

### 2. Run the migrations

Gator uses [Goose](https://github.com/pressly/goose) to set up its tables. Install it, clone this repo, and run the migrations:

```bash
go install github.com/pressly/goose/v3/cmd/goose@latest
git clone https://github.com/AtlasRegap/gator.git
cd gator/sql/schema
goose postgres "postgres://postgres:postgres@localhost:5432/gator" up
```

Replace `postgres:postgres` with your own Postgres username and password.

### 3. Create the config file

Gator reads its settings from `~/.gatorconfig.json`. Create it with your database connection string:

```json
{
  "db_url": "postgres://postgres:postgres@localhost:5432/gator?sslmode=disable"
}
```

Gator adds the `current_user_name` field itself when you register or log in.

## Usage

```bash
gator <command> [arguments]
```

### Users

| Command | What it does |
|---|---|
| `gator register <name>` | Create a new user and log in as them |
| `gator login <name>` | Switch to an existing user |
| `gator users` | List all users (the current one is marked) |

### Feeds

| Command | What it does |
|---|---|
| `gator addfeed <name> <url>` | Add a feed and follow it |
| `gator feeds` | List every feed and who added it |
| `gator follow <url>` | Follow a feed someone else added |
| `gator following` | List the feeds you follow |
| `gator unfollow <url>` | Stop following a feed |

### Posts

| Command | What it does |
|---|---|
| `gator agg <time_between_reqs>` | Keep fetching feeds in a loop, e.g. `gator agg 1m`. Stop it with Ctrl+C. |
| `gator browse [limit]` | Show the newest posts from feeds you follow (default 2) |

### Other

| Command | What it does |
|---|---|
| `gator reset` | Delete all users, feeds and posts. Useful for development; don't run it on data you want to keep. |

## Example

```bash
gator register sagar
gator addfeed "Hacker News" "https://news.ycombinator.com/rss"
gator addfeed "Boot.dev Blog" "https://blog.boot.dev/index.xml"

# In one terminal, leave the aggregator running:
gator agg 1m

# In another terminal, read the latest posts:
gator browse 5
```

## Development

From inside the repo, you can run the code without installing it:

```bash
go run . <command> [arguments]
```

`go run .` is for development. Once you've installed it, use `gator`: it's a compiled binary and doesn't need the Go toolchain to run.
