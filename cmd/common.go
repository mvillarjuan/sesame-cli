package cmd

import "context"

type Data struct {
}

type External interface {
	FetchData(ctx context.Context, id string) (*Data, error)
}

func getSesionId(sesameUrl, sesameUsername, sesamePassword string) (string, error) {
	return "", nil
}
