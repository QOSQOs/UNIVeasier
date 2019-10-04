# Layout and Application Project
## Layout Project

```bash
├── cmd
├── build
├── pkg
│   ├── api
│   │   ├── mockRouter.go
│   │   ├── mockRouter_test.go
│   ├── app
│   │   ├── mockService.go
│   │   ├── mockService_test.go
│   ├── db
│   │   ├── mockDao.go
│   │   ├── mockDao_test.go
│   ├── models
│   │   └── mock
│   │       ├── mockModel.go
│   │       └── mockModel_test.go
├── internal
│   │   ├── configs
│   │   └── UI
│   │       ├── mock.html
│   │       └── mock.js
├── deployments
├── vendor
├── third_party
├── docs
├── test
├── README.md
└── .gitignore
```
1. _cmd_ : keep the binaries.
2. _build_ : scripts to setting up the project.
3. _pkg_ : contain the logic model, there is no restriction for its internal or external use.
4. _internal_ : only available in the scope of the project. External project can not access any file from here.
5. _docs_: official documentation

_NOTE: How is this different between `third_party` and `vendor` directory? If you import and use external code as-is then it should go into the `vendor` directory. If you are using a modified version of an external project then put it in the `third_party` directory._

## Application Layer
### API Layer
This is a _lightweight_ interface to our business logic. When a client interacts with our application, it will do so through an API. Common interface are:
* GraphQL:  that is the most recomendable
* gRPC
* REST

This layer is only responsible for translating the data returned from the business logic layer into a consumable format for clients.

### Business Logic Layer
This is where business logic is handle:
```
Business logic or domain logic is that part of the program which encodes the real-world business rules that determine how data can be created, displayed, stored, and changed. It prescribes how business objects interact with one another, and enforces the routes and the methods by which business objects are accessed and updated.
```
Authorization and authentication are also handle in this layer.

### Persistence Layer
This is where data is persisted to a database or other store.

In this layer, our business logic layer makes these requests to our database. The only thing this layer should be responsible for is database operations; there shouldn't be any business logic.


References:
1. Application Structure: https://aaf.engineering/go-web-application-structure-pt-1/
2. Go Project-Layout: https://medium.com/golang-learn/go-project-layout-e5213cdcfaa2
3. Standard Project-Layout: https://github.com/golang-standards/project-layout
4. GraphQL: https://medium.com/@bradford_hamilton/building-an-api-with-graphql-and-go-9350df5c9356