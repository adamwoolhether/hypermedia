dev.setup:
	go install github.com/bokwoon95/wgo@latest
	go install github.com/a-h/templ/cmd/templ@latest

dev.setup.mobile:
	git clone https://github.com/Instawork/hyperview.git
	sed -i'' -e "s|export const ENTRY_POINT_URL = 'http://0.0.0.0:8085/index.xml';|export const ENTRY_POINT_URL = 'http://0.0.0.0:42069/';|" hyperview/demo/src/constants.ts
	cd hyperview && yarn
	cd hyperview/demo && yarn && \
		yarn add react-native-communications && \
		yarn add react-native-root-toast && \
		yarn add react-native-swipeable
	cp app/hypermedia/mobile/static/js/{email.js,phone.js,toast.js,swipeable.js,HyperviewScreen.tsx} hyperview/demo/src
	xcodebuild -runFirstLaunch
	xcodebuild -downloadPlatform iOS
	cd hyperview && npm update fsevents --force; npm audit fix --force
	cd hyperview/demo && npm update fsevents --force; npm audit fix --force

up:
	go run app/main.go

templ:
	templ generate

dev:
	./run.sh
	@#wgo -file=.go -file=.templ -file=.js -file=.css -xfile=_templ.go templ generate :: go run app/contact/main.go

mobile:
	cd hyperview/demo && yarn ios

run: templ
	@trap 'osascript -e "tell application \"Google Chrome\" to close (tabs of window 1 whose URL contains \"http://localhost:42069/\")"' INT TERM EXIT && \
	open -a "Google Chrome" http://localhost:42069/ && \
	wgo run -xdir=hyperview -xdir=vendor -file=.go -file=.templ -file=.css -file=.js -xfile=_templ.go -verbose app/main.go | go run app/tooling/main.go

tidy:
	go mod tidy
	go mod vendor

# curl -i -X POST http://localhost:42069/api/v1/contacts
android:
	/Users/adam/Library/Android/sdk/emulator/emulator -avd Pixel_8_Pro_API_UpsideDownCakePrivacySandbox -dns-server 8.8.8.8


#import OpenPhone from './phone';
#import OpenEmail from './email';
#import ShowToast from './toast';
#import SwipeableRow from "./swipeable";
#console.log(HyperviewScreen.Behaviors);

#HyperviewScreen.Behaviors = [OpenPhone, OpenEmail, ShowToast];
#let components = [SwipeableRow];

#components={components}
