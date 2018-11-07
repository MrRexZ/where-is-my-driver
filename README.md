# WHERE IS MY GO-JEK DRIVER v3.0- Anthony Tjuatja (anthonytjuatja@gmail.com)

A web server that provides features to insert & update drivers coordinates,
 and to find drivers.
 Also contains a simulator that insert/update 50000 drivers every 1 minute.

## Table of Contents

1.Project Structure

2.Tech Stack

3.Rationale for tech stack

4.Architecture

5.Requirements

6.Instructions


## 1.Project Structure
In the root folder, there contains several multiple folders :
- `cmd` contains main application code
- `pkg` contains business logic, with framework and DB required for the running of the application.
- `controller` contains code to handle connection between incoming request(s) to server and the defined app business rule.
- `vendor` contains dependencies/libraries used in the app.  

Inside `app` folder:
- `main.go` is the main file to start server and simulator to update 50000 drivers latlng info every 1 minute.
- `app/app.go` is the app file to start the server.

Inside `pkg` folder:
- `driver` folder containing frameworks & the business logic for drivers
- `entity` contains enterprise wide business rule and object.

Inside `driver` folder:
- `mocks` contains generated mock data for testing
- `repository` contains code for DB 
- `usecase` contains domain rules, business logic, and serves as an interface between network handler and database.


## 2. Tech Stack

- Language : Go
- Test Frameworks : [Testify](https://github.com/stretchr/testify), [Mockery](https://github.com/vektra/mockery), [Fake](https://github.com/icrowley/fake)
- Network Frameworks : [gorilla/mux](https://github.com/gorilla/mux) 
- Business Logic Framework : [Haversine](https://github.com/umahmood/haversine)
- DB Framework : [MongoGoDriver](https://github.com/mongodb/mongo-go-driver)
- Database : MongoDB
- Containerization Platform : Docker and Docker Compose
      
 
## 3. Rationale for tech stack

I was searching through the web for language that is simple to learn given the time, with considerations to my background, and the context of the application.
As this is a web app and concurrency seems to be one of the key requirement, I decided to seek for language that can is capable of handling concurrency well in terms of simplicity and performance.
Go was among one of the languages, and knowing its reputations among top companies (such as GO-JEK, whose business domain problem fits with the one in this assignment), I decided to go for this language. 
It has simple concurrency models that is more easier to grasp compared to other major languages, and its native library `net/http` provides support to handle requests concurrently by starting a handler for each request that comes in,
which will allow me to spend my time efficiently in other areas.

In testing frameworks, `Mockery` is used to Mock interface, allowing testing of seperate layers without strictly depending on the completion of the other end of layer.
`Testify` is used for assertion and mocking used in Mock generated from `Mockery` 
`Fake` is a mocker used to generate random LatLng.

In network framework, `mux` is used as multiplexer as it is more powerful & its simple syntax in handling URL matching of request to handler compared to the native way.

In business logic framework, only `Haversine` is used to correctly calculate the distance between 2 LatLng coordinates.

DB type used is a NoSQL DB. Reason for choosing NoSQL type is NoSQL is well-known for its speed compared to SQL and given the data needs to be stored frequently, and accessed concurrently.
MongoDB is chosen as I have very basic familiarity with it beforehand, which means more time to learn other stuff.
`MongoGoDriver` was chosen as the interface between Go and MongoDB because it's more simpler than the other variant `mgo`.

## 4. Architecture

Architecture in this project is based on [Clean Architecture](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html).
Relationship between my project structure and the architecture:
- `/pkg/entity` represents `Entity`
- `/pkg/driver/usecase` represents `Usecase`
- `/pkg/driver/repository` and `/cmd/app.go` represents `Frameworks & Drivers`
- `/controller` represents `Interface Adapters`

## 5. Requirements
1.  Go 1.11 or above
2.  Docker and Docker Compose
3.  MongoDB (if running locally w/o Docker image)

## 6.Instructions




