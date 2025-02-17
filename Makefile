# Makefile for building the toggle app in Go

# Target executable
TARGET = toggle

# Source files
SRCS = main.go

# Build the target
build:
	go build -o $(TARGET) $(SRCS)

# Clean the build directory
clean:
	rm -f $(TARGET)
