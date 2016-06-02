# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOINSTALL=$(GOCMD) install
GOTEST=$(GOCMD) test
GODEP=$(GOTEST) -i
GOFMT=gofmt -w

TOPLEVEL_PKG=pinball
GOARGS=GOARCH=arm GOOS=linux
DIST=dist
BINARY=pinball
TARGET=$(DIST)/$(BINARY)

# Set the pi user
RPI_USER?=pi
# Set the rpi ip address to hostname rpi in /etc/hosts
RPI=rpi
all: build

local:
	$(eval GOARGS = )
	$(eval BIN_ARGS = "--no-check --no-hardware" )

run: build
	$(TARGET) $(BIN_ARGS)

build:
	mkdir -p $(DIST)
	$(GOARGS) $(GOBUILD) -o $(TARGET) $(TOPLEVEL_PKG)

launch: build
	scp shell/clean.sh $(RPI_USER)@$(RPI):~/clean.sh
	ssh $(RPI_USER)@$(RPI) "sudo ./clean.sh"
	scp $(TARGET) $(RPI_USER)@$(RPI):~/$(BINARY)
	ssh $(RPI_USER)@$(RPI) "sudo ./$(BINARY)"
