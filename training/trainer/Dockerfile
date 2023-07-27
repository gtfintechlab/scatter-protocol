# Use a base image with Docker installed (e.g., Docker official image)
FROM docker:dind

ARG NODE_ID
ARG TOPIC_NAME
ENV NODE_ID=${NODE_ID} TOPIC_NAME=${TOPIC_NAME}

# Copy the inner Dockerfile into the outer container
COPY /jobs/$NODE_ID/$TOPIC_NAME/Dockerfile /tmp/Dockerfile

RUN mkdir -p /tmp/$NODE_ID/$TOPIC_NAME/output/
RUN mkdir -p /tmp/inner-output/

# Build and run inner Dockerfile inside the outer container
CMD docker build -t inner-image /tmp/ --security-opt=no-new-privileges --security-opt=seccomp:unconfined && \
    docker run --name inner-container -v inner-output:/tmp/inner-output/ -it inner-image && \
    docker cp inner-container:/tmp/inner-output/. /tmp/$NODE_ID/$TOPIC_NAME/output/ && \
    docker image rm inner-image -f && \
    docker rm inner-container

VOLUME inner-output


# Sample Command To Build:
# docker build --build-arg NODE_ID=test --build-arg TOPIC_NAME=test -t outer-image .

# Sample Command To Run:
# docker run --privileged -v /var/run/docker.sock:/var/run/docker.sock -v /home/samrat/Documents/scatterchain/training/trainer/jobs/test/test/output:/tmp/test/test/output/ -it outer-image