# TodoList Backend Application

TodoList is backend application that serving create and fetch a list. Application developed with Golang.

Hexagonal architecture choosed for application which developed using TDD.
Standard golang net/http package was used to server and http supports.

## Table of content
- [APIs](#apis)
    - [Fetch todo items](#fetch-todo-items)
    - [Create Item](#create-item)
- [Database](#database)
    - [InnerDB](#inner-db)
    - [Things to do to use different databases](#things-to-do-to-use-different-databases)
- [Tests](#tests)
    - [Unit](#unit)
    - [Pact](#pack)
- [Deploy](#deploy)
    - [Dockerfile](#dockerfile)
    - [GitLab](#gitlab)
- [Things Do Better](#things-do-better)
  - [Pact Versioning](#pact-versioning)
  - [Can I Deploy Stage](#can-i-deploy)
  - [Added Two More UI Stage](#added-wo-ore-ui-stage)
  - [One More Pact Statement](#one-more-pact-statement)
- [Links](#links)

## APIs

There are two REST endpoint in application serves. APIs were described below.

### Fetch todo items

It is the Get endpoint that Fetches-todo-items. It returns the todo item list in the order of addition.

Usage:<br>
Http Method Type: ``Get``

Endpoint: ``/api/todo-items``

Response Header:
``Content-Type:application/json; charset=UTF-8``

Response Body:

```
[
  {
    "id": 1,
    "title": "test"
  }
]
```

Response status: ``200``

### Create Item

It is the service that adds an element to the todo list.  The element is added to the end of the list.

Usage:<br>
Http Method Tipi: `Post`

Endpoint: ``/api/todo-item``

Request Header:
``Content-Type:application/json; charset=UTF-8``

Request Body:

```
{
  "title": "test"
}
```

Response status: ``201``

## Database

### InnerDB

A list as database was used to hold elements. Due to hard to deploy external database on kubernetes, list was choose for database.

### Database selection rule

External database can be used optionally. A struct that implement the ITodoRepository interface should be created for the external database to be used.

ITodoRepository interface:
```
Insert(item *model.TodoItem) *model.TodoItem
FindAll() model.TodoItems
```

## Tests

The application developed using TDD. Two different tests were writen as Unit and Pact tests for application. Unit and pact tests were seperated execution time by using ``tags`` keyword.

The tests runs with the following commands:

- [ ] Unit: ``go test ./... -tags=unit``
- [ ] Pact: ``go test ./... -tags=pact``

### Unit

Unit tests were written for handler and service classes. Tests written using Golang test util were positioned at the same level as classes and mock services that created by ``mockgen`` served under mock package.

To use mockgen:
mockgen kullanımı için:

Install mockgen with ``go install github.com/golang/mock/mockgen@v1.6.0`` command

Commands that creating mock services used in project:

``mockgen -source .\handler\todoController.go -destination ./mock/mock_todoController.go -package mock`` <br>
``mockgen -source .\handler\todoService.go -destination ./mock/mock_todoService.go -package mock`` <br>
``mockgen -source .\handler\todoRepository.go -destination ./mock/mock_todoRepository.go -package mock``

- ``-source``: Source file path (A interface must be in file)
- ``-destination``: Mock service destination path
- ``-package``: Describe package of mock service

### Pact

Pact was used to verify endpoints to serve consumer. Pact CLI Tools were installed to use pact's utils.
And then to achieve pact in golang, pact packages was gets with command ``go get github.com/pact-foundation/pact-go@v1``
UI projesinde kullanılacak endpointerin doğruluğu için pact kullanıldı. Pacti kullanabilmek için ``Pact CLI Tools``
kuruldu. Sonrasında ``go get github.com/pact-foundation/pact-go@v1`` komutu ile golange pact tanıtıldı.

Pact files were placed at Pactflow. To verify the pact file, pact broker ```(https://sak-assignment.pactflow.io/)``` has been accessed with access token.
Pact file that last created was verified.

The settings used in pact tests:

```
Host: "127.0.0.1" // Mock server address
Provider: "todo-backend" // Provider project name described by consumer
Consumer: "todo-frontend" // Consumer project name described by consumer
```

```
ProviderBaseURL: fmt.Sprintf("http://localhost:%d", port) // Mock server address
BrokerURL: "https://sak-assignment.pactflow.io" // where pact files are stored.
ProviderVersion: "1.0.1" // provider version
ConsumerVersionSelectors: []types.ConsumerVersionSelector{types.ConsumerVersionSelector{Latest: true}} // Consumer selector
BrokerToken: "XXXXX" // Pact broker access token
PublishVerificationResults: true // publish verification results configuration
StateHandlers: // Pact states
```

### Deploy
Project stored at GitLab and deploy steps were created using GitLab pipeline. Project is deployed on kubernetes provided by
`Google Cloud Platform`. Kubernetes configuration files are stored at `deployment-*-env`.

A parameter named `APP_ENV` was declared at docker container to determine which environment the project will run on.
This parameter can be change from deployment configuration.


Addresses of applications:<br>
Test environment : ``http://35.190.154.212`` <br>
Production environment: ``http://34.148.116.140``

Note : Reason why Google Cloud Platform choose is $300 free credit.

### Dockerfile
Dockerfile was created to deploy applications to kubernetes. Dockerfile consist of application execute file and config files.


```
FROM alpine   # alpine base image
WORKDIR /app  # create and redirect folder named app

RUN mkdir .config       # create folder named '.config' to obtain config files
RUN mkdir target        # create folder named 'target' to store application built file
COPY .config .config/   # copy configuration files into .config folder

COPY ./target/karaca-assignment ./target/   # copy executable file to target folder
ENV APP_ENV=local_env                       # set default environment variable. It's used to determine which configuration application run with

CMD ["./target/karaca-assignment"]    # execute application when docker run command enter
```

<br>
APP_ENV variable usage:

```
  docker run -d -p <HOST PORT>:3000 -e APP_ENV=<local_env | test_env | production_env>  <image name>
```

Note : if variable APP_ENV didn't specify in command, variable set default value is local_env

### GitLab

A pipeline declared on GitLab with ``.gitlab-ci.yml`` file. Stages on pipeline described above.
<br>

- [ ] build : application builds and stores in folder ``target/karaca-assignment``
- [ ] unit_test : runs unit tests with command ``go test -v ./... -tags=unit``
- [ ] dockerize : creates an image from Dockerfile and push GitLab repository
- [ ] deploy2TestEnv : pushes kubernetes test environment configuration files to Google Cloud Platform to kubernetes
  pull images stored GitLab repository
- [ ] consumer-driven-contract-tests : runs Pact Provider Tests with command ``go test -v ./... -tags=pact``
- [ ] Deploy2Prod : pushes kubernetes production environment configuration files to Google Cloud Platform to kubernetes
  pull images stored GitLab repository

## Things Do Better
This section about what things I could have done better.
### Pact Versioning
I was pushed pact file that I created to pactflow without set new version.
So, I was changed manually the version of pact file. 
### Can I Deploy Stage
I could add a stage before deploy2Prod but I did not do that. 
Because I couldn't set environments variable on pact file and can_i_deploy util want it.
### Added Two More UI Stage
I was added two more stage to UI stages. 
The reason why I was added it that I wanted to seperate built application with production and test.
### One More Pact Statement
I could add one more pact statement about empty input. I did not it because it is loo late.

## Links

* [Mockgen](https://github.com/golang/mock)
* [Pact-go](https://github.com/pact-foundation/pact-go)
* [Pactflow](https://pactflow.io/)
* [GitLab CI/CD](https://docs.gitlab.com/ee/ci/)
* [Google Cloud](https://cloud.google.com/)
