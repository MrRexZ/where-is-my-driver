# WHERE IS MY GO-JEK DRIVER v3.0- Anthony Tjuatja (anthonytjuatja@gmail.com)

A web server that provides features to insert & update drivers coordinates,
 and to find drivers.
 Also contains a simulator that insert/update 50000 drivers every 1 minute.

## Table of Contents
1.Tech Stack

2.Rationale for tech stack

3.Infrastructure

4.Instructions

## 1. Tech Stack

Language : Go

Frameworks : https://github.com/stretchr/testify, https://github.com/mongodb/mongo-go-driver, https://github.com/umahmood/haversine
             , https://github.com/gorilla/mux, https://github.com/icrowley/fake
             , https://github.com/vektra/mockery

Database : MongoDB             
 
 
## 2. Rationale for tech stack

I was searching through the web for language that is simple to learn given the time, with considerations to my background, and the context of the application.
As this is a web app and concurrency seems to be one of the key requirement, I decided to seek for language that can is capable of handling concurrency well in terms of simplicity and performance.
Go was among one of the languages, and knowing its reputations among top companies (such as GO-JEK, whose business domain problem fits with the one in this assignment), I decided to go for this language. 
