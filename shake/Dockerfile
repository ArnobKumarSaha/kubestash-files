FROM ubuntu:20.04
ENV DEBIAN_FRONTEND noninteractive
WORKDIR /opt
RUN apt-get clean && apt-get update && apt-get install -y --no-install-recommends --no-install-suggests wget curl openjdk-21-jdk \
    && wget https://github.com/alibaba/MongoShake/releases/download/release-v2.8.4-20230425/mongo-shake-v2.8.4.tgz \
    && tar -xzf mongo-shake-v2.8.4.tgz \
    && mv mongo-shake-v2.8.4 mongo-shake \
    && rm -rf mongo-shake/collector.linux \
    && rm -rf mongo-shake-v2.8.4.tgz \
    && apt-get autoremove -y \
    && rm -rf /var/lib/apt/lists/*

COPY bin/collector.linux /opt/mongo-shake/collector.linux

RUN chmod a+x /opt/mongo-shake/collector.linux

WORKDIR /opt/mongo-shake
CMD ["/opt/mongo-shake/collector.linux", "-conf=/opt/mongo-shake/collector.conf", "-verbose=4"]