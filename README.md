## go-todo-backend [![CircleCI](https://circleci.com/gh/circleci/circleci-docs.svg?style=svg)](https://app.circleci.com/pipelines/github/aasimsajjad22/go-todo-backend)
The backend API has been created with MVC pattern isolating each layer and follow the TDD approach.<br>
You can find frontend application consuming that API here [react-todo-frontend](https://github.com/aasimsajjad22/react-todo-frontend)

### How to run
##### - Run through Docker
- If you have docker installed then update `.env` file with your database credientials and type `docker-compose up` to setup the application in docker environment.
#### - Run Manually
- This app uses mod for dependencies, make sure to install **go >= 1.13** and also **mysql** database. <br>
You can check go version by typing ``` go --version ```on terminal and then download all the dependencies by typing ```go mod download``` this will install all the external libraries used in project. <br>
- Make sure to add env variables for your mysql credentials.
- After that cd into src folder and type ```go run main.go``` this will launch our server. <br>
- you can use postman tool to test the routes or simply use curl.
- ```curl -X GET 'http://localhost:8081/todos' ``` to get list of all todos 
- ```curl -X POST localhost:8081/todo -d '{"description": "Testing description"}' ```
- To test simply type ``` go test ./..``` this will run all your test cases, you can also append --cover flag to see coverage as well.

### Code Coverage
According to Goland Tool we have 100% code coverage (100% files, 100% statements)

### Project Structure.
- [ ] **src/app**
  - contains URL mappings and router setup to route our application to controller functions
- [ ] **src/controllers**
  - contains todo controller and respective test cases (Integration Tests)
- [ ] **src/datasources**
  - contains database configuration, in our case i have used mysql.
- [ ] **src/domain**
  - core domain for our application, it contains data access and data transfer files and all interaction with database layer also i have added unit testing to test each single line of code. for database i have used Mocking. 
- [ ] **src/services**
  - service layer is actually to perform business logic, validation etc take request from controller perform all business req and return data to controller
- [ ] **src/tests**
  - contains functional testing where we actually runs the application and perform testing.
- [ ] **src/utils**
  - support functionality like RestErr which manages rest errors for our API.
 
### Run Tests.
  - To test simply type ``` go test ./..``` this will run all your test cases, you can also append --cover flag to see coverage as well
