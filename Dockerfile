FROM ubuntu:18.04

RUN \
	apt-get update && \
	apt-get upgrade -y && \
	apt-get install -y wget tar git

RUN \
	wget https://go.dev/dl/go1.18.4.linux-amd64.tar.gz && \
	rm -rf /usr/local/go && tar -C /usr/local -xzf go1.18.4.linux-amd64.tar.gz && \
	export PATH=$PATH:/usr/local/go/bin

