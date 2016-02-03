FROM docker.io/centos:latest
MAINTAINER Bowen Denning <bdenning@gmail.com>

ENV PATH=$PATH:/opt/splunk/bin

RUN curl http://download.splunk.com/products/splunk/releases/6.3.0/splunk/linux/splunk-6.3.0-aa7d4b1ccb80-linux-2.6-x86_64.rpm -O && \
    yum install -y splunk*.rpm && \
    rm splunk*.rpm

RUN splunk version --accept-license
RUN splunk edit user admin -password p@ssword -role admin -auth admin:changeme

COPY entrypoint.sh /sbin/entrypoint.sh
RUN chmod +x /sbin/entrypoint.sh

EXPOSE 8000/tcp 8089/tcp 9997/tcp 9514/tcp
WORKDIR /opt/splunk

USER splunk

ENTRYPOINT ["/sbin/entrypoint.sh"]
