#!/bin/sh

# join listed consul hosts
if [ -z $CONSUL_JOIN ]; then
    echo "CONSUL_JOIN is empty"
    exit 1
fi

JOIN_STR=""
for node in $CONSUL_JOIN; do
    JOIN_STR="${JOIN_STR} -retry-join ${node}"
done

cp /etc/consul.tpl.json /etc/consul.json
/usr/local/bin/ep /etc/consul.json
ret=$?
if [ $ret != 0 ]; then
    echo "envplate failed, some env vars not set"
    exit 1
fi
/usr/local/bin/consul agent -config-file /etc/consul.json $JOIN_STR &

set -- docker-entrypoint.sh $@

exec "$@"