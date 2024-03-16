install:
	go install github.com/bokwoon95/wgo@latest
	go install github.com/a-h/templ/cmd/templ@latest

tidy:
	go mod tidy
	go mod vendor

dev:
	./run.sh
	@#wgo -file=.go -file=.templ -file=.js -file=.css -xfile=_templ.go templ generate :: go run app/contact/main.go

run: templ
	@trap 'osascript -e "tell application \"Google Chrome\" to close (tabs of window 1 whose URL contains \"http://localhost:42069/\")"' INT TERM EXIT && \
	open -a "Google Chrome" http://localhost:42069/ && \
	go run main.go | go run app/tooling/main.go

templ:
	templ generate

# curl -i -X POST http://localhost:42069/api/v1/contacts -d '{"first_name":"adam","last_name":"woo","phone":"1234567489","email":"email@example.com"}'