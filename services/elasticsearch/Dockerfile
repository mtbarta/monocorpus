FROM docker.elastic.co/elasticsearch/elasticsearch-oss:6.2.2

ENV CONSUL_VERSION 0.9.3

# install consul
RUN yum update -y && yum install -y ca-certificates wget unzip && \
    wget https://releases.hashicorp.com/consul/${CONSUL_VERSION}/consul_${CONSUL_VERSION}_linux_amd64.zip -O /tmp/consul.zip && \
    unzip -d /usr/local/bin /tmp/consul.zip && \
    rm -f /tmp/consul.zip && \
    wget https://github.com/kreuzwerker/envplate/releases/download/v0.0.8/ep-linux -O /usr/local/bin/ep && \
    chmod +x /usr/local/bin/ep && \
    yum remove -y wget unzip

COPY entry-consul.sh /usr/local/bin/entry-consul.sh
COPY consul.tpl.json /etc

CMD '/usr/local/bin/entry-consul.sh'