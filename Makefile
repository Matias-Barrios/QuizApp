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

install-linux:
	systemctl stop quizappservice
	mkdir -p /opt/quizapp
	cp -r static /opt/quizapp/.
	mv quiz /opt/quizapp/quiz
	cp quizapp.env /opt/quizapp/quizapp.env
	chmod -R 550 /opt/quizapp
	cp quizappservice.service /lib/systemd/system/quizappservice.service
	chmod ag+r /lib/systemd/system/quizappservice.service
	systemctl start quizappservice

