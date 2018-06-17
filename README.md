# MonoCorpus
https://mtbarta.github.io/monocorpus/

MonoCorpus is a tool to record code, notes, and papers during
the development of software and machine learning algorithms. It
lets you lookup previous bugs and relevant notes to quickly solve new problems and bugs.

This repo is a collection of microservices and scripts to deploy the service. It is being reorganized during its transition from a private repo, so some things may be nonfunctional or pointing to nonexistent code.

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

MonoCorpus uses Vuejs on the frontend, GraphQL as a gateway layer between HTTP and grpc calls, and golang microservices in the backend.

MonoCorpus relies on Keycloak for user auth, traefik as a proxy to the backend, and Elasticsearch for full text search.

Microservices:
* Notes
  - CRUD functions to Mongo.
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

KEY=INSERT_KEYCLOAK_REALM_PUBLIC_KEY
```

Deployment requires the gateway to be aware of keycloak's public key for token decryption.

1. `docker-compose up -d keycloak` to bring up the keycloak instance.
2. Find the notes realm public key and replace the dummy key in the `.env` file.
3. Elasticsearch may need permissions to write to your elasticsearch data directory. Run `chown -R 1000:1000 es` to fix this.
4. Keycloak's realm requires knowledge of valid redirect urls. If you are not running on localhost, this needs to be changed.
5. `docker-compose up -d` to bring the whole system up.

## Notable Dependencies

Search - [Elasticsearch](https://github.com/elastic/elasticsearch)

Proxy -  [Traefik](https://github.com/containous/traefik)

Authentication - [Keycloak](https://github.com/keycloak/keycloak)

## Contributors

@mtbarta

## How to Contribute

Please feel free to send me a PR. There's a lot of low-hanging fruit across this project -- refactoring, documentation, testing. Let me know if there's something you want to work on and we can discuss.

Please rebase PRs if necessary -- https://github.com/edx/edx-platform/wiki/How-to-Rebase-a-Pull-Request

## License

GNU AGPLv3
