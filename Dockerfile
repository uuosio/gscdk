FROM rocker/binder:3.6.3

## Declares build arguments
ARG NB_USER
ARG NB_UID

## Copies your repo files into the Docker Container
USER root
RUN python3 --version
## Install eosio.cdt 1.7.0
RUN apt update
RUN apt install -y libncurses5
RUN apt install -y ninja-build
RUN apt install -y cmake
RUN apt install -y clang
RUN apt install -y git
RUN wget https://golang.org/dl/go1.16.7.linux-amd64.tar.gz
RUN rm -rf /usr/local/go && tar -C /usr/local -xzf go1.16.7.linux-amd64.tar.gz
RUN ln -sf /usr/local/go/bin/go /usr/local/bin/go
RUN rm go1.16.7.linux-amd64.tar.gz

RUN wget https://github.com/EOSIO/eosio.cdt/releases/download/v1.7.0/eosio.cdt_1.7.0-1-ubuntu-18.04_amd64.deb
RUN apt install -y ./eosio.cdt_1.7.0-1-ubuntu-18.04_amd64.deb
RUN rm ./eosio.cdt_1.7.0-1-ubuntu-18.04_amd64.deb

RUN wget https://github.com/uuosio/uuosio.gscdk/releases/download/v0.1.0/release.tar.gz
RUN tar -C /usr/local -xzf release.tar.gz
RUN ln -sf /usr/local/uuosio.gscdk/bin/tinygo /usr/local/bin/tinygo
RUN ln -sf /usr/local/uuosio.gscdk/bin/eosio-go /usr/local/bin/eosio-go
RUN ln -sf /usr/local/uuosio.gscdk/bin/eosio-strip /usr/local/bin/eosio-strip

COPY . ${HOME}
## Enable this to copy files from the binder subdirectory
## to the home, overriding any existing files.
## Useful to create a setup on binder that is different from a
## clone of your repository
## COPY binder ${HOME}
RUN chown -R ${NB_USER} ${HOME}

## Become normal user again
USER ${NB_USER}

## Run an install.R script, if it exists.
RUN if [ -f install.R ]; then R --quiet -f install.R; fi
RUN python3 -m pip install https://github.com/uuosio/UUOSKit/releases/download/v0.8.4/uuoskit-0.8.4-cp37-cp37m-linux_x86_64.whl
