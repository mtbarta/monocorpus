# MonoCorpus

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

## Notable Dependencies

Search - [Elasticsearch](https://github.com/elastic/elasticsearch)
Proxy -  [Traefik](https://github.com/containous/traefik)
Authentication - [Keycloak](https://github.com/keycloak/keycloak)


## License

GNU AGPLv3