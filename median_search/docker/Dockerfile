FROM ubuntu:16.04

WORKDIR /opt

ENV PROJ_DIR=/opt/cs133_hw1
COPY include ${PROJ_DIR}/include
COPY src ${PROJ_DIR}/src
COPY test ${PROJ_DIR}/test
COPY test_cases ${PROJ_DIR}/test_cases
COPY cmake ${PROJ_DIR}/cmake
COPY CMakeLists.txt ${PROJ_DIR}
COPY docker/apt_source.txt /etc/apt/sources.list

RUN apt-get update && \
    apt-get install --yes \
        build-essential \
        git \
        cmake \
        vim && \
    rm -rf /var/lib/apt/lists/*
