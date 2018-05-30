VERSION?=latest

build:
	cd notes && $(MAKE) build
	cd gateway && $(MAKE) build
	cd frontend && $(MAKE) build
	cd keycloak && $(MAKE) build

containers:
	cd notes && $(MAKE) container
	cd gateway && $(MAKE) container
	cd frontend && $(MAKE) container
	cd keycloak && $(MAKE) container