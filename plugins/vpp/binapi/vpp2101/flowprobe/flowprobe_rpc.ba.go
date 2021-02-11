// Code generated by GoVPP's binapi-generator. DO NOT EDIT.

package flowprobe

import (
	"context"

	api "git.fd.io/govpp.git/api"
)

// RPCService defines RPC service  flowprobe.
type RPCService interface {
	FlowprobeParams(ctx context.Context, in *FlowprobeParams) (*FlowprobeParamsReply, error)
	FlowprobeTxInterfaceAddDel(ctx context.Context, in *FlowprobeTxInterfaceAddDel) (*FlowprobeTxInterfaceAddDelReply, error)
}

type serviceClient struct {
	conn api.Connection
}

func NewServiceClient(conn api.Connection) RPCService {
	return &serviceClient{conn}
}

func (c *serviceClient) FlowprobeParams(ctx context.Context, in *FlowprobeParams) (*FlowprobeParamsReply, error) {
	out := new(FlowprobeParamsReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) FlowprobeTxInterfaceAddDel(ctx context.Context, in *FlowprobeTxInterfaceAddDel) (*FlowprobeTxInterfaceAddDelReply, error) {
	out := new(FlowprobeTxInterfaceAddDelReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}
