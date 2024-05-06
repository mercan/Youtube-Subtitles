build:
	docker build -t go-youtube-subtitles .

run:
	docker run -p 8080:8080 go-youtube-subtitles

help:
	@echo "Commands:"
	@echo "build - build the image"
	@echo "run - run the image"
	@echo "help - display this help"

clean:
	docker rm -f go-youtube-subtitles
	docker rmi go-youtube-subtitles
