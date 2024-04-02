# <img src="icons\go.svg" alt="go-icon" width="50" height="50"> <img src="icons\api.svg" alt="go-icon" width="50" height="50"> Go REST API Template 

Icons were free from https://iconscout.com/icons <br>

This repository provides a template for building a RESTful API in Go. It includes basic infrastructure, CRUD operations, authentication using OAuth2, Dockerization, and hosting as a microservice.

## Project Description
This project is a REST API written in Go that serves as a template for anyone looking to build their own RESTful API. It provides a starting point with common features and best practices.

## Endpoints
- users/all

## Goals
- [X] Create a basic infrastructure
- [X] Create a startup DB setup cpabilities
- [X] Create a build.bat
- [ ] Create a Makefile
- [ ] Add testing capcbilities
- [ ] Implement all CRUD operations
    - [ ] Add user
    - [ ] Get user
    - [ ] Delete User
- [ ] Implement subrouting
- [ ] Add authentication using OAuth2
- [ ] Dockerize the application
- [ ] Host it as a microservice

## Getting Started
1. Clone this repository.
2. Customize the code to fit your specific requirements.
3. Create a Database by the specifications the API that is provided (For now a postgres)
3. Run the application locally for testing and development.
    1. build.bat
    2. Makefile (soon)
4. Deploy the application to your preferred hosting service.

## Usage
1. Install Go and any necessary dependencies.
2. Build the application using `go build`.
3. Run the application using `.cmd/rest_api/rest_api` (or rest_api.exe if using Windows).

## Contributing
Contributions are welcome! Feel free to open issues or submit pull requests.

## License
This project is licensed under the MIT License - see the LICENSE file for details.
