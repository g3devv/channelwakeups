## gomobile `wakeups_resource` bug minimal example

This repo follows up on the request in https://github.com/golang/go/issues/29284#issuecomment-447816700 to provide a way to reliably reproduce that error.

The go project creates a "few" channels which causes the `wakeups_resource` to trigger, the consequence being a complete crash of the app.

### Instructions to set up the Go layer

- `mkdir GoChannelCrash`

- `cd GoChannelCrash`

- Clone the project 

- `cd` into the `goproj` subdirectory

- `chmod +x makeFramework.sh`

- `./makeFramework.sh`

The framework will be written to `../frameworks/goprojlib.xcframework`

### Instructions to set up the Swift layer

No complicated build settings are at play here, the only thing the project has to do is to make that single function call to the library. The simplest way to set this up is the following:

- Create an Xcode project (default template `iOS` -> `App`), Storyboard (no SwiftUI)

- Drag the newly-generated `goprojlib.xcframework` into Xcode (I suggest to deselect `Copy items if needed` - so whenever you run the `makeFramework.sh` again it will automatically be updated in Xcode too)

- Replace the contents in the Xcode project's `ViewController` with the contents in `ViewController.swift`

### Running the application

- Install the app on your physical device. When ran via Xcode, the `wakeups_resource` issue won’t come up (same in my "actual app", never happens with connected Xcode). So to reproduce:

- Disconnect the USB cable, then run the app by tapping it on the home screen. On my old iPhone it crashes after about 30 seconds.

### Gathering crash logs

On your iPhone, enter `Settings > Privacy > Analytics & Improvements > Analytics Data` and you will find the crash report there.

Note: The `Analytics Data` screen will not refresh new reports automatically if you let the app run and crash again - to refresh the page go back to the `Analytics & Improvements` screen, then tap `Analytics Data` and the new report will be visible.

### Further notes

Basically all the go project does is to spawn channels, write to them which results into a lot of blocked and interrupted threads which according to a response in https://developer.apple.com/forums/thread/124180 leads up to that issue.

I am not familiar with golang at all, it's possible that there are way smarter ways to trigger this issue but I think this demonstration is fairly concise in what it does.

Since the `wakeups_resource` issue won't trigger when ran via Xcode (maybe the WatchDog is disabled during debug sessions?) it will eventually hit a memory shortage but I suppose these issues are unrelated. Feel free to change anything around in the code.
