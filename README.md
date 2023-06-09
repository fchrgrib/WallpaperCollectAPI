# WallpaperCollectionAPI

## Overview
This project is an example of a REST API built using Golang and the Gin Framework, with SQLite as the database and JWT for authentication. The API allows users to create and manage a collection of wallpapers, everyone can use this rest api and give some feedback to me about this code.

## Getting Started

1. Make sure you have Go and SQLite installed on your machine.

2. Clone the repository to your local machine:
````bash
Copy code
git clone https://github.com/fchrgrib/wallpaperCollectRestAPI.git
````
3. Install the necessary dependencies:
````
go get -v github.com/gin-gonic/gin
go get -v github.com/jinzhu/gorm
go get -v github.com/jinzhu/gorm/dialects/sqlite
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

1. `POST`   /register: Register a new user
2. `POST`   /login: Login and get a JWT
3. `GET`    /logout : logout
4. `GET`    /wallpaper: Get a list of all wallpapers
5. `POST`   /wallpaper/upload: Upload a new wallpaper
6. `GET`    /wallpaper/profile: to get all user info
7. `PUT`    /wallpaper/profile/update_profile: for updating your profile
8. `PUT`    /wallpaper/profile/upload_profile_picture: to upload your profile picture
