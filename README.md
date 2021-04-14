# bbb-player
BigBlueButton Download Plugin + BigBlueButton Player

#### Docker Installation
```
docker run -d -v $(pwd)/config.yml:/config/config.yml -v <your published dir>/var/bigbluebutton/published/presentation:/published  -p 127.0.0.1:3535:3535  maaami98/bbb-download-api
```
