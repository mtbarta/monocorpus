VERSION?=latest

build:
	cd services/notes && $(MAKE) build
	cd services/gateway && $(MAKE) build
	cd web && $(MAKE) build
	cd services/keycloak && $(MAKE) build

containers:
	cd services/notes && $(MAKE) container
	cd services/gateway && $(MAKE) container
	cd services/frontend && $(MAKE) container
	cd servioces/keycloak && $(MAKE) container