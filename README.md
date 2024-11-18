# tmux-watcher

### Installation: 
##### Requirements: 
- Latest version of [go](url) 
-  tmux > version 3.5

### Introduction

This is a tool to track the time spent working on any given repository / session indexed by tmux session name.  It tracks your active sessions and writes to a json file or if you'd like, a google sheet through the google api (requires google cloud project).  

It's not required but I recommend using this alongside [tmux-sessionizer](https://github.com/jrmoulton/tmux-sessionizer).  

#### Examples:
- Student tracking time spent on each class / project
- Tracking hours worked on a client project as a freelance dev
- Anyone interested to see how long they spend writing code

### Configuration: 

Create a file at ~/.config/tmux-watcher/config.json

Config options: 

```json

{
    "enabled_repositories": ["repo1", "repo2"], // session-names to look for
    "write_location": "/path/to/desired/out-file" // default is ./sessions.json
    "sheets-api-options": { // unstable
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
