# MonoCorpus
https://mtbarta.github.io/monocorpus/

**DEPRECATED**   
Unfortunately, I accepted a job where this project is a conflict of interest. This project will be in maintanence mode from here on out. 

***
MonoCorpus is a tool to record code, notes, and papers during
the development of software and machine learning algorithms. It
lets you lookup previous bugs and relevant notes to quickly solve new problems and bugs.

<p align="center">
  <img src="./notebook.gif" alt="MonoCorpus example"
       width="654" height="450">
</p>

## Key Features

* Full Text Search
* Chronological ordering of notes
* Title Filters
* import Arxiv abstracts
* KaTeX support
* Markdown support
* Image support (beta)

## How it Works

**gateway proxying into gRPC**
MonoCorpus uses Vuejs on the frontend, GraphQL as a gateway layer between HTTP and grpc calls, and golang microservices in the backend.

**async messaging**
As notes are written and sent to the backend, we save them and then send them to a messaging platform for other services to pick up asynchronously. 

**robust authentication**
We use keycloak to provide authentication and authorization of users. Every user is able to enable two-factor auth for their account.

Microservices:
* Notes
  - CRUD functions to Mongo.
* Search
  - Async CRUD functions to add notes to Elasticsearch.
* Gateway
  - GraphQL interface to notes and search.
  
### Setup

Monocorpus should be installed inside of your $GOPATH at `github.com/mtbarta/monocorpus`.

create an .env file in `monocorpus/docker`. It should have the following variables:
```
ENV=dev
NETWORK=docker
HOST=localhost

POSTGRES_DATA_LOC=/data/postgres
POSTGRES_USER=loginUserAdmin
POSTGRES_PASSWORD=adminpw

MONGO_DATA_LOC=/data/mongo
SEARCH_DATA_LOC=/data/es

KEYCLOAK_USER=admin
KEYCLOAK_PASSWORD=admin
```

There are a couple snags that I've run into deploying this:

1. the system needs to be aware of the host name. This is for proper routing of calls, decryption of tokens, and ACME support in traefik.
3. Elasticsearch may need permissions to write to your elasticsearch data directory. Run `chown -R 1000:1000 es` to fix this.
4. Keycloak's realm requires knowledge of valid redirect urls. If you are not running on localhost, this needs to be changed. You can boot up keycloak, and go it it's endpoint at `/auth/admin` and login with the above user and password. https://www.keycloak.org/docs/latest/server_admin/index.html#_clients has more information about where to navigate to change client urls.
5. `docker-compose up -d` to bring the whole system up.

## Notable Dependencies

Search - [Elasticsearch](https://github.com/elastic/elasticsearch)

Proxy -  [Traefik](https://github.com/containous/traefik)

Authentication - [Keycloak](https://github.com/keycloak/keycloak)

## Contributors

@mtbarta

## How to Contribute

Please feel free to send me a PR. There's a lot of low-hanging fruit across this project -- refactoring, documentation, testing, new features. Let me know if there's something you want to work on and we can discuss.

Please rebase PRs if necessary -- https://github.com/edx/edx-platform/wiki/How-to-Rebase-a-Pull-Request

## License

GNU AGPLv3
