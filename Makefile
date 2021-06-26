TARGET_NAME			= cdn
TARGET_EXTENSION 	=
TARGET				= $(TARGET_NAME)$(TARGET_EXTENSION)
GO_LD_FLAGS			= -s -w

ifeq ($(OS),Windows_NT)
	TARGET_EXTENSION = .exe
endif

all: clean lint
	go build -ldflags="$(GO_LD_FLAGS)" -o $(TARGET)

lint:
	go vet

.PHONY: clean
clean:
	rm -f $(TARGET)

.PHONY: clean-content
clean-content:
	rm -rf ./content ./*/content ./*/**/content
