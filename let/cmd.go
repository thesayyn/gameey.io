package main

import (
	"context"
	"fmt"
	"log"
	"path/filepath"

	"github.com/containerd/containerd"
	"github.com/containerd/containerd/namespaces"
)

func main() {
	path, _ := filepath.Abs("./containerd/containerd.sock")
	client, err := containerd.New(path)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	ctx := namespaces.WithNamespace(context.Background(), "example")

	cntrs, _ := client.Containers(ctx)

	img, err := client.Pull(ctx, "docker.io/itzg/minecraft-server:latest", containerd.WithPullUnpack)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(cntrs)
	fmt.Println(img.Target())
}
