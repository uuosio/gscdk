FROM ubuntu:18.04


RUN apt update
RUN apt install -y wget
RUN apt install -y python3.7-dev
RUN apt install -y python3-pip
RUN python3.7 -m pip install --upgrade pip
ARG DEBIAN_FRONTEND=noninteractive
RUN apt install -y --no-install-recommends postgresql
RUN python3.7 -m pip install pytest
RUN python3.7 -m pip install uuoskit
RUN python3.7 -m pip install gscdk
RUN python3.7 -m pip install https://github.com/learnforpractice/binder-test/releases/download/0.1/uuosio-0.1.0-cp37-cp37m-linux_x86_64.whl

#install golang
RUN wget https://golang.org/dl/go1.16.7.linux-amd64.tar.gz
RUN rm -rf /usr/local/go && tar -C /usr/local -xzf go1.16.7.linux-amd64.tar.gz
RUN ln -sf /usr/local/go/bin/go /usr/local/bin/go

