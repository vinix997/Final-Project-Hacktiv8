# Final-Project-Hacktiv8

FINAL PROJECT DOCUMENTATION <br />

git clone https://github.com/vinix997/Final-Project-Hacktiv8.git<br />
cd final-project<br />
go run main.go<br />

URL: localhost:8080 <br />
All endpoint requires auth bearer token except register and login <br />
to generate token, please do login or register and login first. <br />

USER API: <br />
POST: <br />

url: /users/register<br />
[request body]
```json
{
    "Username": "brokins",
    "Email": "brokin@gmail.com",
    "Password": "brokin",
    "Age": 22
}
```

POST:<br />
url: /users/login<br />
[request body]
```json
{
    "Email": "brokin@gmail.com",
    "Password": "brokin",
}
```

PUT:<br />
url: /users/{userId}<br />
[request body]<br />
```json
{
    "Username": "brokinsupdated",
    "Email": "brokinupdated@gmail.com",
}
```

DELETE:<br />
/users/{userId}<br />

PHOTO API:<br />

POST:<br />
url: /photos<br />
[request body]<br />
```json
{
    "title": "Title",
    "caption": "Caption",
    "photo_url": "http:www"
}
```

PUT:<br />
url: /photos/{photoId}<br />
[request body]<br />
```json
{
    "title": "TitleUpdated",
    "caption": "CaptionUpdated",
    "photo_url": "http:wwwEdit"
}
```

GET:<br />
url: /photos<br />

DELETE:<br />
url: /photos/{photoId}<br />

SOCIAL MEDIA API:<br />

POST:<br />
url: /socialmedias<br />
[request body]<br />
```json
{
    "name": "twitter media",
    "social_media_url": "twitter.com"
}
```

PUT:<br />
url: /socialmedias/{socialMediaId}<br />
[request body]<br />
```json
{
    "name": "twitter media",
    "social_media_url": "twittermedia.com"
}
```

GET:<br />
url: /socialmedias<br />

DELETE:<br />
url: /socialmedias/{socialMediaId}<br />

COMMENT API:<br />

POST:<br />
url: /comments<br />
[request body]<br />
```json
{
    "message": "Keren bro",
    "photo_id": 1
}
```

PUT:<br />
url: /comments/{commentId}<br />
[request body]<br />
```json
{
    "message": "Keren banget bro",
    "photo_id": 1
}
```

GET:<br />
url: /comments<br />

DELETE:<br />
url: /comments/{commentId}<br />
