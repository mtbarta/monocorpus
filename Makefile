VERSION?=latest

build:
	cd services/notes && $(MAKE) build
	cd services/gateway && $(MAKE) build
	cd web && $(MAKE) build
	cd services/keycloak && $(MAKE) build

containers:
	cd services/notes && $(MAKE) container
	cd services/gateway && $(MAKE) container
	cd services/search && $(MAKE) container
	cd web && $(MAKE) container
	cd services/keycloak && $(MAKE) container
	cd services/mongo && $(MAKE) container
	cd services/elasticsearch && $(MAKE) container
