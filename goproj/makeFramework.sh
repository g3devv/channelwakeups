go get -d golang.org/x/mobile
go get -d golang.org/x/mobile/bind

FRAMEWORK_DESTINATION_PATH=../frameworks/goprojlib.xcframework

gomobile bind -target ios -o "$FRAMEWORK_DESTINATION_PATH" ./goprojlib

