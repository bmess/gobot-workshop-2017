# Drone - Parrot Minidrone

## Prereqs

### OS X

```
brew install pkg-config
brew install sdl2
```

### Linux

Not yet...

## Installation

```
go get -d -u gobot.io/x/gobot/...
```

## Running the code
When you run any of these examples, you will compile and execute the code on your computer. When you are running the program, you will be communicating with the drone using the Bluetooth Low Energy (LE) interface.

To compile/run the code, substitute the name of your drone as needed:

### OS X

```
$ GODEBUG=cgocheck=0 go run 5-freeflight.go RS_1234
```

### Linux

```
go run 5-freeflight.go RS_1234
```

## Code

### 1-takeoff.go

### 2-data.go

### 3-move.go

### 4-flips.go

### 5-flightcontrol.go

## License

Copyright (c) 2015-2017 The Hybrid Group. Licensed under the MIT license.
