Allows users to quickly find the desired text in YouTube videos by showing the time, minute, and second intervals where the desired word or sentence appears in videos with subtitles or auto-generated subtitles.
### Run on Your Computer

#### 1. Clone the Project:

```bash
  git clone https://github.com/mercan/Go-Youtube-Subtitles.git
```

#### 2. Go to Project Directory:

```bash
    cd Go-Youtube-Subtitles
```

### Using Docker

#### Build the Docker Image:

This command builds a Docker image named go-youtube-subtitles containing your project code.

```bash
    make build
```

#### Run the Docker Image:

This command starts the Docker container based on the built image. It maps the container port 8080 to the host port 8080.

```bash
    make run
```

### Available Commands:

You can use the following Makefile commands to manage the Docker image:

* make build: Builds the Docker image.
* make run: Runs the Docker image.
* make help: Displays this help message.
* make clean: Stops and removes any running container and the built image.

&nbsp;

**Note:** Make sure you have Docker installed and running on your system to use these commands.


### API Usage

#### Search Subtitles:

```http
  GET /subtitles?url=YoutubeVideoURL&text=SearchTerm
```

| Parameter | Type     | Description                  |
| :-------- | :------- | :--------------------------- |
| `url` | `string` | **Required.** Youtube Video URL. |
| `text` | `string` | **Required.**. The word/sentence you want to find in the video. |

#### Example Usage

Find the timestamps of the word "Hello" in the following YouTube video:

```bash
    curl -X GET "http://localhost:8080/subtitles?url=https://www.youtube.com/watch?v=YQHsXMglC9A&text=Hello"
```

This request will return a JSON response with the timestamps where the word "Hello" appears in the video.