# Builds the Docker image for the youtube-subtitles application
build:
	@echo "\033[92mBuilding Docker Image...\033[0m"
	docker build -t youtube-subtitles .
	@echo "\033[92mDone\033[0m"

# Runs the Docker image for the youtube-subtitles application in detached mode with the specified port number or default port 8080
ifeq ($(PORT),)
    PORT := 8080
endif

run:
	@echo "\033[92mRunning Docker Image...\033[0m"
	docker run -d --name youtube-subtitles -p $(PORT):8080 youtube-subtitles
	@echo "\033[92mListening on port $(PORT)\033[0m"
	@echo "\033[92mDone\033[0m"


# Stops the running youtube-subtitles Docker container
stop:
	@echo "\033[92mStopping Docker Image...\033[0m"
	docker stop youtube-subtitles
	@echo "\033[92mDone\033[0m"

# Shows the logs of the youtube-subtitles Docker container
logs:
	@echo "\033[92mShowing logs...\033[0m"
	docker logs youtube-subtitles

# Removes the youtube-subtitles Docker image and container
clean:
	@echo "\033[92mRemoving Docker Image...\033[0m"
	docker rm -f youtube-subtitles
	docker rmi -f youtube-subtitles
	@echo "\033[92mDone\033[0m"

# Displays help message about available commands
help:
	@echo "\033[92mCommands:\033[0m"
	@echo "\033[92mbuild - build the image\033[0m"
	@echo "\033[92mrun - run the image in detached mode with the specified port number or default port 8080\033[0m"
	@echo "\033[92mstop - stop the image\033[0m"
	@echo "\033[92mlogs - show logs\033[0m"
	@echo "\033[92mclean - remove the image\033[0m"
	@echo "\033[92mhelp - display this help\033[0m"


# Declares these targets as phony to ensure they are always rebuilt
.PHONY: build run stop logs clean help