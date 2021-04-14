# bbb-player
BigBlueButton Download Plugin + BigBlueButton Player

#### Docker Installation
```
docker run -d  -v <your published dir>/var/bigbluebutton/published/presentation:/published  -p 127.0.0.1:3535:3535  maaami98/bbb-download-api

mv download.nginx /etc/bigbluebutton/nginx/

service nginx restart

```

#### Usage
```
https://example.com/download/presentation/<meeting_id>

```

