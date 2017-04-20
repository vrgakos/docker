package client

import (
	"github.com/docker/docker/api/types"
	"golang.org/x/net/context"
	"encoding/json"
)

// CheckpointCreate creates a checkpoint from the given container with the given name
func (cli *Client) CheckpointCreate(ctx context.Context, container string, options types.CheckpointCreateOptions) (types.CheckpointStat, error) {
	var stat types.CheckpointStat

	resp, err := cli.post(ctx, "/containers/"+container+"/checkpoints", nil, options, nil)
	if err != nil {
		return stat, err
	}

	err = json.NewDecoder(resp.body).Decode(&stat)
	ensureReaderClosed(resp)
	return stat, err
}
