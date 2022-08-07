# Final-Project-Hacktiv8

FINAL PROJECT DOCUMENTATION
URL: localhost:8080

git clone https://github.com/vinix997/Final-Project-Hacktiv8.git
cd final-project
go run main.go

USER API:
POST:

url: /users/register
[request body]
{
    "Username": "brokins",
    "Email": "brokin@gmail.com",
    "Password": "brokin",
    "Age": 22
}

POST:
url: /users/login
[request body]
{
    "Email": "brokin@gmail.com",
    "Password": "brokin",
}

PUT:
url: /users/{userId}
[request body]
{
    "Username": "brokinsupdated",
    "Email": "brokinupdated@gmail.com",
}

DELETE:
/users/{userId}

PHOTO API:

POST:
url: /photos
[request body]
{
    "title": "Title",
    "caption": "Caption",
    "photo_url": "http:www"
}

PUT:
url: /photos/{photoId}
[request body]
{
    "title": "TitleUpdated",
    "caption": "CaptionUpdated",
    "photo_url": "http:wwwEdit"
}

GET:
url: /photos

DELETE:
url: /photos/{photoId}

SOCIAL MEDIA API:

POST:
url: /socialmedias
[request body]
{
    "name": "twitter media",
    "social_media_url": "twitter.com"
}

PUT:
url: /socialmedias/{socialMediaId}
[request body]
{
    "name": "twitter media",
    "social_media_url": "twittermedia.com"
}

GET:
url: /socialmedias

DELETE:
url: /socialmedias/{socialMediaId}

COMMENT API:

POST:
url: /comments
[request body]
{
    "message": "Keren bro",
    "photo_id": 1
}

PUT:
url: /comments/{commentId}
[request body]
{
    "message": "Keren banget bro",
    "photo_id": 1
}

GET:
url: /comments

DELETE:
url: /comments/{commentId}
