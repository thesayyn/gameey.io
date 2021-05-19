package main

import (
	"context"
	"log"
	"os"
	"path/filepath"

	"github.com/containerd/containerd"
	"github.com/containerd/containerd/namespaces"
	"github.com/containerd/containerd/oci"
	"github.com/opencontainers/runtime-spec/specs-go"
	"github.com/urfave/cli/v2"
)

func run(c *cli.Context) error {

	containerdPath, err := filepath.Abs(c.String("containerd"))

	if err != nil {
		return err
	}

	client, err := containerd.New(containerdPath)

	if err != nil {
		return err
	}

	defer client.Close()

	log.Println("Successfully connected to the containerd")

	ctx := namespaces.WithNamespace(context.Background(), "game1")

	imageName := "docker.io/itzg/minecraft-server:latest"

	log.Printf("Pulling image %s", imageName)

	image, err := client.Pull(ctx, imageName, []containerd.RemoteOpt{
		containerd.WithSchema1Conversion,
		containerd.WithPullUnpack,
	}...)

	if err != nil {
		return err
	}

	size, err := image.Size(ctx)

	if err != nil {
		return err
	}

	log.Printf("Successfully pulled the image %d with digest %s", size, image.Target().Digest.String())

	id := "minecraft"

	var s specs.Spec

	spec := containerd.WithSpec(&s,
		oci.WithDefaultSpec(),
		oci.WithDefaultUnixDevices,
		oci.WithImageConfig(image),
	)

	container, err := client.NewContainer(
		ctx,
		id,
		containerd.WithNewSnapshot(id, image),
		containerd.WithImage(image),
		containerd.WithImageStopSignal(image, "SIGTERM"),
		spec,
	)

	if err != nil {
		return err
	}

	log.Printf("Successfully created the container %s", container.ID())

	defer container.Delete(ctx, containerd.WithSnapshotCleanup)

	return nil
}

func main() {
	app := &cli.App{
		Name:        "let",
		Description: "Gameey let agent",
		Flags: []cli.Flag{
			&cli.PathFlag{
				Name:       "containerd",
				Value:      "/var/run/containerd/containerd.sock",
				HasBeenSet: true,
				TakesFile:  true,
				Usage:      "Relative or absolute path to containerd.sock",
			},
		},
		Action: run,
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
