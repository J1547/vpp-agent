// Code generated by GoVPP's binapi-generator. DO NOT EDIT.

package memclnt

import (
	"context"

	api "go.fd.io/govpp/api"
)

// RPCService defines RPC service memclnt.
type RPCService interface {
	APIVersions(ctx context.Context, in *APIVersions) (*APIVersionsReply, error)
	ControlPing(ctx context.Context, in *ControlPing) (*ControlPingReply, error)
	GetFirstMsgID(ctx context.Context, in *GetFirstMsgID) (*GetFirstMsgIDReply, error)
	MemclntCreate(ctx context.Context, in *MemclntCreate) (*MemclntCreateReply, error)
	MemclntDelete(ctx context.Context, in *MemclntDelete) (*MemclntDeleteReply, error)
	MemclntKeepalive(ctx context.Context, in *MemclntKeepalive) (*MemclntKeepaliveReply, error)
	MemclntReadTimeout(ctx context.Context, in *MemclntReadTimeout) error
	MemclntRxThreadSuspend(ctx context.Context, in *MemclntRxThreadSuspend) error
	RPCCall(ctx context.Context, in *RPCCall) (*RPCCallReply, error)
	RxThreadExit(ctx context.Context, in *RxThreadExit) error
	SockInitShm(ctx context.Context, in *SockInitShm) (*SockInitShmReply, error)
	SockclntCreate(ctx context.Context, in *SockclntCreate) (*SockclntCreateReply, error)
	SockclntDelete(ctx context.Context, in *SockclntDelete) (*SockclntDeleteReply, error)
	TracePluginMsgIds(ctx context.Context, in *TracePluginMsgIds) error
}

type serviceClient struct {
	conn api.Connection
}

func NewServiceClient(conn api.Connection) RPCService {
	return &serviceClient{conn}
}

func (c *serviceClient) APIVersions(ctx context.Context, in *APIVersions) (*APIVersionsReply, error) {
	out := new(APIVersionsReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) ControlPing(ctx context.Context, in *ControlPing) (*ControlPingReply, error) {
	out := new(ControlPingReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) GetFirstMsgID(ctx context.Context, in *GetFirstMsgID) (*GetFirstMsgIDReply, error) {
	out := new(GetFirstMsgIDReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) MemclntCreate(ctx context.Context, in *MemclntCreate) (*MemclntCreateReply, error) {
	out := new(MemclntCreateReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) MemclntDelete(ctx context.Context, in *MemclntDelete) (*MemclntDeleteReply, error) {
	out := new(MemclntDeleteReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) MemclntKeepalive(ctx context.Context, in *MemclntKeepalive) (*MemclntKeepaliveReply, error) {
	out := new(MemclntKeepaliveReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) MemclntReadTimeout(ctx context.Context, in *MemclntReadTimeout) error {
	stream, err := c.conn.NewStream(ctx)
	if err != nil {
		return err
	}
	err = stream.SendMsg(in)
	if err != nil {
		return err
	}
	err = stream.Close()
	if err != nil {
		return err
	}
	return nil
}

func (c *serviceClient) MemclntRxThreadSuspend(ctx context.Context, in *MemclntRxThreadSuspend) error {
	stream, err := c.conn.NewStream(ctx)
	if err != nil {
		return err
	}
	err = stream.SendMsg(in)
	if err != nil {
		return err
	}
	err = stream.Close()
	if err != nil {
		return err
	}
	return nil
}

func (c *serviceClient) RPCCall(ctx context.Context, in *RPCCall) (*RPCCallReply, error) {
	out := new(RPCCallReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) RxThreadExit(ctx context.Context, in *RxThreadExit) error {
	stream, err := c.conn.NewStream(ctx)
	if err != nil {
		return err
	}
	err = stream.SendMsg(in)
	if err != nil {
		return err
	}
	err = stream.Close()
	if err != nil {
		return err
	}
	return nil
}

func (c *serviceClient) SockInitShm(ctx context.Context, in *SockInitShm) (*SockInitShmReply, error) {
	out := new(SockInitShmReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) SockclntCreate(ctx context.Context, in *SockclntCreate) (*SockclntCreateReply, error) {
	out := new(SockclntCreateReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) SockclntDelete(ctx context.Context, in *SockclntDelete) (*SockclntDeleteReply, error) {
	out := new(SockclntDeleteReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) TracePluginMsgIds(ctx context.Context, in *TracePluginMsgIds) error {
	stream, err := c.conn.NewStream(ctx)
	if err != nil {
		return err
	}
	err = stream.SendMsg(in)
	if err != nil {
		return err
	}
	err = stream.Close()
	if err != nil {
		return err
	}
	return nil
}
