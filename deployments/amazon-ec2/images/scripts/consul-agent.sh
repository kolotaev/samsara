#!/bin/bash -e

if [ "$(id -u)" != "0" ] ; then
    echo "Running the script with root privilege's"
    sudo "$0"
    exit $?
fi

export CONSUL_VERSION=0.5.2

echo '------------------------------------------------------------------'
echo '                    Consul agent installation'
echo '------------------------------------------------------------------'
apt-get install -y unzip
echo "(*) Agent installation..."
wget "https://dl.bintray.com/mitchellh/consul/${CONSUL_VERSION}_linux_amd64.zip" \
     -O /tmp/consul.zip && \
    unzip /tmp/consul.zip -d /tmp && \
    mv /tmp/consul /usr/sbin/ && \
    rm -fr /tmp/consul.zip

echo "(*) Consul user creation"
useradd -r -s /bin/false consul
usermod -a -G logger consul
mkdir -p /data/consul
chown -R consul:consul /data/consul

echo "(*) Consul config dir"
mkdir -p /etc/consul.d

echo "(*) Consul upstart configuration"
# TODO: join ips must be passed as user-data
cat > /etc/init/consul.conf <<\EOF
description "Consul agent"
author "Bruno"
setuid consul
start on runlevel [2345]
stop on runlevel [016]
respawn
script
   exec >> /logs/consul.log 2>&1
   exec /usr/sbin/consul agent \
       -retry-join 10.10.1.5 \
       -retry-join 10.10.2.5 \
       -retry-join 10.10.3.5 \
       -data-dir /tmp/consul \
end script
EOF

echo "(*) Set DNS/Consul integration"
apt-get install -y dnsmasq
echo "server=/consul/127.0.0.1#8600" > /etc/dnsmasq.d/10-consul
