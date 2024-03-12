# Music Streaming App Backend API
This project is a backend API for a music streaming application written in **GoLang**. It provides functionalities for user-management, music-management, and etc.

### Features
- User Management
    - User registration (signup)
    - User login (signin)
    - User logout
    - Authentication (jwt)
- Music Management
    - Adding new music to the streaming library
    - Retrieving existing music information
    - Deleting music from the library
    - Update the music information
 
### Technologies Used
[![My Skills](https://skillicons.dev/icons?i=golang,docker,nginx,git,github,postman,postgres,aws)](https://skillicons.dev)

### Getting Started
Clone the repository:
```bash
git clone https://github.com/tahadostifam/MusicStreamingApp.git
```
Set up the development environment:
```bash
cd MusicStreamingApp
go mod tidy
sudo docker-compose up -d postgres minio
```
Check the configs:
```bash
nano ./config/configs.yml
```
Run the API:
```bash
go run ./cmd/server.go
```
API Documentation:   
https://www.postman.com/crimson-equinox-208211/workspace/musicstreamingapp    

Author:
Taha. Dostifam (gihtub.com/tahadostifam)
