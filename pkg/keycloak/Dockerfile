FROM jboss/keycloak-postgres:4.0.0.Beta1

ENV DB_VENDOR=POSTGRES
ENV PROXY_ADDRESS_FORWARDING=true

ADD realm-export.json /opt/jboss/monocorpus/realm-export.json

CMD ["-b", "0.0.0.0", "-Dkeycloak.import=/opt/jboss/monocorpus/realm-export.json"]