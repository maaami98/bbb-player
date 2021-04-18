# bbb-player
BigBlueButton Download Plugin + BigBlueButton Player

#### Docker Installation
```
docker run -d  -v <your published dir>/var/bigbluebutton/published/presentation:/published  -p 127.0.0.1:3535:3535  maaami98/bbb-download-api

cd /etc/bigbluebutton/nginx/ && wget https://raw.githubusercontent.com/maaami98/bbb-player/main/download.nginx

 cd /var/bigbluebutton/playback/presentation/2.0/acornmediaplayer/themes/bigbluebutton/ && wget https://github.com/maaami98/bbb-player/raw/main/bigbluebutton-download.png

service nginx restart

```

#### Usage
```
https://example.com/download/presentation/<meeting_id>

```

