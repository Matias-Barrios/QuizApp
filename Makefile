# Go parameters
GOCMD=go 
GOBUILD=$(GOCMD) build
BINARY_NAME=quiz
BINARY_NAME_MACOS=quiz_darwin

clean: 
	rm -f $(BINARY_NAME)
	rm -f $(BINARY_NAME_MACOS)
   
# Specific compilation

build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_NAME) -v

build-macos:
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 $(GOBUILD) -o $(BINARY_NAME_MACOS) -v

build-install-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_NAME) -v
	mkdir -p /opt/quizapp
	mv static /opt/quizapp/static
	mv quiz /opt/quizapp/quiz
	chmod -R 550 /opt/quizapp
	mv quizappservice.service /lib/systemd/system/quizappservice.service
	chmod ag+r /lib/systemd/system/quizappservice.service

