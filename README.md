# tmux-watcher

### Installation: 
##### Requirements: 
- Latest version of [go](https://go.dev/) 
-  tmux > version 3.5

### Introduction

This is a tool to track the time spent working on any given repository / session indexed by tmux session name.  It tracks your active sessions and writes to a json file or if you'd like, a google sheet through the google api (requires google cloud project).  

Also an alternative to [aw-watcher-tmux](https://github.com/akohlbecker/aw-watcher-tmux) if you don't want touse [ActivityWatcher](https://activitywatch.net/) . Or if your're using wsl.

It's not required but I recommend using this alongside [tmux-sessionizer](https://github.com/jrmoulton/tmux-sessionizer).  

### Configuration: 

Create a file at ~/.config/tmux-watcher/config.json

Config options: 

enabled_repositories: A list of session names you want to watch


```js

{
    "enabled_repositories": ["repo1", "repo2"], 
    "write_location": "/path/to/desired/out-file" 
    "sheets-api-options": { 
        "enabled": false,
        "sheet_id": "abc123",
    }
}

```

1. Clone the repo
```sh
git clone https://github.com/cjodo/tmux-watcher.git
cd tmux-watcher
```

2. Run either of the following.

```sh
go build .
./tmux-watcher
```
If you want it to run in the background. 
```sh 
./build-and-run.sh
```
