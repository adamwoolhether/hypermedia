install:
	go install github.com/bokwoon95/wgo@latest
	go install github.com/a-h/templ/cmd/templ@latest

tidy:
	go mod tidy
	go mod vendor

run:
	./run.sh
	@#wgo -file=.go -file=.templ -file=.js -file=.css -xfile=_templ.go templ generate :: go run app/contact/main.go

templ:
	templ generate