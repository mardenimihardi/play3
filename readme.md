# Play 3 test

## The Case
## Objective
We need to build a service for users to find delicious dishes in New York. Users should
be able to search by dish name or restaurant name.
From a user experience perspective, searching must be enjoyable, including a feature
like auto-completion when typing on the search box. On the results, users should be
able to filter based on several categories like price range, events, and venues.

## Brief
Use What's on the Menu? Data from New York Public Library
Deliver the service in REST API or other interfaces you think appropriate.
Share the codebase on a code repository.
A running service that we can try is a plus.

## Requirements
* Git v1.7.10 or later version
* Go v1.18.0 or later version
* Mockgen 1.4 or later version

## Installation
Clone the project
``` bash
git clone https://github.com/mardenimihardi/play3.git
```
Restore dependencies
``` bash
go mod tidy
go mod vendor
```
Build and run the project
``` bash
go run ./app.go 
```
or run makefile to run the project & create testfile
```
make
``` 

### How to see
```
1. Open browser
2. Type localhost:9000
3. Done 
```

### Host

| Environment | Host |
| ------------ | ------------- |
| localhost | localhost:9000|
| Staging |  |
| Production | |  