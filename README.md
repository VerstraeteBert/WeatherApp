# WeatherApp

### Prerequisites
- Go 1.10.1+
- Mysql 5.7+
- Dep package manager for Go


### Setup
Checkout   
Run dep ensure in project root   
Manually run migrations (in ./migrations folder)
Copy .env.example to .env and enter your MySql config    
Build and run main.go (in ./cmd folder)       

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
- [ ] Build Web FrontEnd (React, vue?)
- [ ] Build Web FrontEnd (React, vue?)
- [ ] Dockerize 
- [ ] Host 
- [ ] Expand CI/CD
- [ ] More / better designed tests

