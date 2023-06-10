# WallpaperCollectionAPI

## Overview
an exclusive API for wallpaperCollectApp. you can login,register,collect image,update profile in this API. I use MySQL for database with using GORM for manipulate relation of database and also Gin Framework for build this API. I deploy this API in https://wallpapercollectapi-production-c728.up.railway.app/. feel free to feedback this API.

## Getting Started

1. Make sure you have Go and SQL installed on your machine.

2. Clone the repository to your local machine:
````bash
Copy code
git clone https://github.com/fchrgrib/wallpaperCollectRestAPI.git
````
3. Install the necessary dependencies:
````
go get -v github.com/gin-gonic/gin
go get -v github.com/jinzhu/gorm
go get -v github.com/jinzhu/gorm/dialects/sql
go get -v github.com/dgrijalva/jwt-go
````
4. Create a .env file in the root of the project and set your JWT secret key:
````
JWT_SECRET=yoursecretkey
````
5. Run the migration files to create the necessary tables in the SQLite database:
````
go run migrations/*.go
````
6. Start the server:
````
go run main.go
````
7. The API will be running on http://localhost:8080. Use a tool like Postman to test the endpoints.

## Endpoints

### Unprotected
1. `POST`   /register-email-default: Register a new user
2. `POST`   /login-email-default: Login and get a JWT
3. `GET`    /logout : logout
4. `GET`    /login-google-session
5. `GET`    /register-goolgle-session
6. `GET`    /login-facebook-session
7. `GET`    /register-facebook-session

### Protected
1. `GET`    /wallpaper: Get a list of all wallpapers
2. `POST`   /wallpaper/upload: Upload a new wallpaper
3. `GET`    /wallpaper/profile: to get all user info
4. `PUT`    /wallpaper/profile/update_profile: for updating your profile
5. `PUT`    /wallpaper/profile/upload_profile_picture: to upload your profile picture
6. `GET`    /images/:id
7. `GET`    /images/:id/download
8. `DELETE` /images/:id/delete
