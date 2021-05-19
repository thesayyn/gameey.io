#!/bin/bash
docker build -t containerd containerd
docker run -p 12375:12375/TCP containerd 

# docker exec -it fe771a22ae24 socat TCP-LISTEN:12375,reuseaddr,fork UNIX-CLIENT:/run/containerd/containerd.sock
# socat UNIX-LISTEN:./containerd/containerd.sock,fork,reuseaddr,unlink-early,mode=777 TCP:localhost:12375

# socat UNIX-LISTEN:./containerd/containerd.sock,fork,reuseaddr,unlink-early,mode=777 TCP:5.39.117.15:12375