#!/usr/bin/bash

## services
services=("places" "forums" "auth" "api")
for idx in "${!services[@]}"; do
if [[ $idx -eq 0 ]]; then
        tmux new -d -s rundev "cd ${services[$idx]} && go run main.go";
    else
        tmux split-window -d -t rundev:1 -p20 -v "cd ${services[$idx]} && go run main.go";
    fi
done
tmux select-layout -t rundev:1 main-vertical
tmux attach-session -t rundev