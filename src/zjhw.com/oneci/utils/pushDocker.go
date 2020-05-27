package utils

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"io"
	"os"
)

func PushDocker(app, version, project, env, ty string, timestamp int64) {
	const DOCKER_HOST = "tcp://127.0.0.1:2376"
	os.Setenv("DOCKER_HOST", DOCKER_HOST)
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}
	authConfig := types.AuthConfig{
		Username: "admin",
		Password: "dd@2019",
	}
	encodedJSON, err := json.Marshal(authConfig)
	if err != nil {
		panic(err)
	}
	authStr := base64.URLEncoding.EncodeToString(encodedJSON)

	out, err := cli.ImagePull(ctx, fmt.Sprintf("harbor.zjhw.com/%s-%s/%s:%s.%d.%s",
		project, env, app, version, timestamp, ty), types.ImagePullOptions{RegistryAuth: authStr})

	if err != nil {
		panic(err)
	}

	defer out.Close()
	io.Copy(os.Stdout, out)
}
