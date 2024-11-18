#!/bin/bash

echo "terminating previous tmux-watcher process..."
pkill tmux-watcher 

rm tmux-watcher

echo "building tmux-watcher..."

go build -o tmux-watcher || {
	echo "build failed!"
	exit 1
}

export PATH=$PATH:$(pwd)

source ~/.bashrc

echo "runnung tmux-watcher in the background..."
nohup ./tmux-watcher > app.log 2>&1 &


PID=$!
echo "tmux-watcher is running in the background with PID $PID"
echo "logs are being written to app.log"

zsh
