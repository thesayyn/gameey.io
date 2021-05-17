FROM ubuntu

RUN apt-get update
RUN apt-get install curl gnupg2 -y 

RUN curl -fsSL https://download.docker.com/linux/debian/gpg | apt-key add -
RUN echo "deb [arch=amd64] https://download.docker.com/linux/debian buster stable" | tee /etc/apt/sources.list.d/docker.list

RUN apt-get update -y 
RUN apt-get install containerd socat -y

COPY config.toml /etc/containerd/config.toml

CMD [ "containerd" ]