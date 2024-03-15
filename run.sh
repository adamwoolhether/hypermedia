#!/bin/zsh
#wgo -file=.js npx tailwindcss -i view/css/app.css -o public/styles.css :: \
wgo -file=.go -file=.templ -file=.css -file=.js -xfile=_templ.go go run app/main.go | go run app/tooling/main.go &
PID1=$!
printf "wgo\t\tPID: %s\n" $PID1

templ generate --watch --proxy=http://localhost:42069 &
PID2=$!
printf "templ\t\tPID: %s\n" PID2$

cleanup() {
  for pid in $PID2 $PID1 $PID0; do
    if kill -0 $pid 2>/dev/null; then
      printf "PID\t\t%s stopping...\n" $pid
      kill $pid || printf "PID\t\t%s not stopped\n" $pid
    else
      printf "PID\t\t%s stopped\n" $pid
    fi

    # ensure wgo is stopped
    pkill wgo
  done
}

trap cleanup SIGINT

wait $PID2 $PID1

osascript -e 'tell application "Google Chrome" to close (tabs of window 1 whose URL contains "http://127.0.0.1:7331/")'