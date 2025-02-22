TARGET = toggle

SRCS = main.go

build:
	go build -o $(TARGET) $(SRCS)

clean:
	rm -f $(TARGET)
