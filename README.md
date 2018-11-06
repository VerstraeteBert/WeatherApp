<h1 align="center">WeatherApp</h1>

<p align="center">
  <a href="https://goreportcard.com/report/github.com/VerstraeteBert/WeatherApp">
  	<img src="https://goreportcard.com/badge/github.com/VerstraeteBert/WeatherApp" alt="">
  </a>
  <a href="https://circleci.com/gh/VerstraeteBert/WeatherApp/tree/master">
  	<img src="https://circleci.com/gh/VerstraeteBert/WeatherApp/tree/master.svg?style=svg" alt="">
  </a>
</p>    


### Description
Creating a weather app that receives weather data from IoT devices via the MQTT protocol,    
which then get persisted and can be consulted through a web interface.

### Prerequisites
- Go 1.10.1+
- Mysql 5.7+
- Dep package manager for Go


### Setup
Checkout   
Run dep ensure in project root   
Copy .env.example to .env and enter your MySql config / desired port   
Make sure you have a database with name "weatherdb" and then run make migrate
Run make build to start the server

### Endpoints
Readings:

| URL        								| Method           	| Info  |
| ------------- 							|:-------------:	| -----:|
| http://127.0.0.1:1337/readings     			| GET 				| Returns an array of readings |
| http://127.0.0.1:1337/readings     			| POST 				| Endpoint to create a reading |



# To do
- [X] Set up CI
- [ ] Add endpoint to change reading interval
- [ ] Allow MQTT communication (from proxy arduino)
- [ ] Expand current endpoints for more weather data
- [ ] Authenticate devices and their users (JWT)
- [ ] Build Web FrontEnd (React, vue?)
- [ ] Dockerize 
- [ ] Host 
- [ ] Expand CI/CD
- [ ] More / better designed tests

