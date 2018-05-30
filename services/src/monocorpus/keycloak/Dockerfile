FROM jboss/base-jdk:8-alpine

ENV KEYCLOAK_VERSION 4.0.0.Beta1
ENV MAVEN_VERSION 3.3.9
# Enables signals getting passed from startup script to JVM
# ensuring clean shutdown when container is stopped.
ENV LAUNCH_JBOSS_IN_BACKGROUND 1

USER root

RUN curl -fsSL https://archive.apache.org/dist/maven/maven-3/$MAVEN_VERSION/binaries/apache-maven-$MAVEN_VERSION-bin.tar.gz | tar xzf - -C /usr/share \
  && mv /usr/share/apache-maven-$MAVEN_VERSION /usr/share/maven \
  && ln -s /usr/share/maven/bin/mvn /usr/bin/mvn

ENV MAVEN_HOME /usr/share/maven

USER jboss

ADD docker-entrypoint.sh /opt/jboss/
COPY realm-export.json /opt/jboss/monocorpus/realm-export.json

EXPOSE 8080

ENTRYPOINT [ "/opt/jboss/docker-entrypoint.sh" ]

CMD ["-b", "0.0.0.0", "-Dkeycloak.import=/opt/jboss/monocorpus/realm-export.json"]
