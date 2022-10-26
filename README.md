# Campfire Reads

Prerequisits for local development:
* Install [Docker Desktop](https://www.docker.com/products/docker-desktop/)

## Docker Setup
Add a `.api.env` and `.ui.env` file that is a copy of `.api.env.exaple` and `.ui.env.example` filled in with your values.

**IMPORTANT:** the ui.env file is a dotenv file and the api.env file is a docker env file, the ui.env file should us `=` syntax and the api.env file should use `:` syntax to assign values. 

Then simply run `docker-compose up` in the root folder of this project

The hosts are as follows:
* API: localhost:8080
* UI: localhost:4200
* Postgres: localhost:5432

Default Postgres username is `root` default password is `secret`

### Local Setup (you should probably just use docker instead)
#### API
You'll need to install https://github.com/cosmtrek/air

Set up an air config in a file called `.air.toml`, here is an example:

```
root = "."
testdata_dir = "testdata"
tmp_dir = "tmp"

[build]
  args_bin = []
  bin = "./tmp/api"
  cmd = "go build -o ./tmp/api ./cmd/api/main.go"
  delay = 1000
  exclude_dir = ["assets", "tmp", "vendor", "testdata"]
  exclude_file = []
  exclude_regex = ["_test.go"]
  exclude_unchanged = false
  follow_symlink = false
  full_bin = "GOOGLE_API_KEY={yourAPIKey} SERVER_PORT=8080 ./tmp/api"
  include_dir = []
  include_ext = ["go", "tpl", "tmpl", "html"]
  kill_delay = "0s"
  log = "build-errors.log"
  send_interrupt = false
  stop_on_error = true

[color]
  app = ""
  build = "yellow"
  main = "magenta"
  runner = "green"
  watcher = "cyan"

[log]
  time = false

[misc]
  clean_on_exit = false

[screen]
  clear_on_rebuild = false
```

Once that exists just run `air`


#### UI
Just CD into static and run `npm run dev`

##### Environments

You'll need to set your environment values for the API server by adding a `.env` file in the `static` folder with your environment variables.  You can use `.ui.example` as an example. 