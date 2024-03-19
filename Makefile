dev.setup:
	go install github.com/bokwoon95/wgo@latest
	go install github.com/a-h/templ/cmd/templ@latest

dev.setup.mobile:
	git clone https://github.com/Instawork/hyperview.git
	sed -i'' -e "s|export const ENTRY_POINT_URL = 'http://0.0.0.0:8085/index.xml';|export const ENTRY_POINT_URL = 'http://0.0.0.0:42069/mobile';|" hyperview/demo/src/constants.ts
	cd hyperview && yarn
	cd hyperview/demo && yarn
	xcodebuild -runFirstLaunch
	xcodebuild -downloadPlatform iOS

templ:
	templ generate

dev:
	./run.sh
	@#wgo -file=.go -file=.templ -file=.js -file=.css -xfile=_templ.go templ generate :: go run app/contact/main.go

mobile:
	cd ../hyperview/demo && yarn ios

run: templ
	@trap 'osascript -e "tell application \"Google Chrome\" to close (tabs of window 1 whose URL contains \"http://localhost:42069/\")"' INT TERM EXIT && \
	open -a "Google Chrome" http://localhost:42069/ && \
	go run main.go | go run app/tooling/main.go

tidy:
	go mod tidy
	go mod vendor

# curl -i -X POST http://localhost:42069/api/v1/contacts 