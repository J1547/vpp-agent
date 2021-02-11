// Code generated by GoVPP's binapi-generator. DO NOT EDIT.

package ip

import (
	"context"
	"fmt"
	"io"

	api "git.fd.io/govpp.git/api"
	vpe "go.ligato.io/vpp-agent/v3/plugins/vpp/binapi/vpp2101/vpe"
)

// RPCService defines RPC service  ip.
type RPCService interface {
	IoamDisable(ctx context.Context, in *IoamDisable) (*IoamDisableReply, error)
	IoamEnable(ctx context.Context, in *IoamEnable) (*IoamEnableReply, error)
	IPAddressDump(ctx context.Context, in *IPAddressDump) (RPCService_IPAddressDumpClient, error)
	IPContainerProxyAddDel(ctx context.Context, in *IPContainerProxyAddDel) (*IPContainerProxyAddDelReply, error)
	IPContainerProxyDump(ctx context.Context, in *IPContainerProxyDump) (RPCService_IPContainerProxyDumpClient, error)
	IPDump(ctx context.Context, in *IPDump) (RPCService_IPDumpClient, error)
	IPMrouteAddDel(ctx context.Context, in *IPMrouteAddDel) (*IPMrouteAddDelReply, error)
	IPMrouteDump(ctx context.Context, in *IPMrouteDump) (RPCService_IPMrouteDumpClient, error)
	IPMtableDump(ctx context.Context, in *IPMtableDump) (RPCService_IPMtableDumpClient, error)
	IPPuntPolice(ctx context.Context, in *IPPuntPolice) (*IPPuntPoliceReply, error)
	IPPuntRedirect(ctx context.Context, in *IPPuntRedirect) (*IPPuntRedirectReply, error)
	IPPuntRedirectDump(ctx context.Context, in *IPPuntRedirectDump) (RPCService_IPPuntRedirectDumpClient, error)
	IPReassemblyEnableDisable(ctx context.Context, in *IPReassemblyEnableDisable) (*IPReassemblyEnableDisableReply, error)
	IPReassemblyGet(ctx context.Context, in *IPReassemblyGet) (*IPReassemblyGetReply, error)
	IPReassemblySet(ctx context.Context, in *IPReassemblySet) (*IPReassemblySetReply, error)
	IPRouteAddDel(ctx context.Context, in *IPRouteAddDel) (*IPRouteAddDelReply, error)
	IPRouteDump(ctx context.Context, in *IPRouteDump) (RPCService_IPRouteDumpClient, error)
	IPRouteLookup(ctx context.Context, in *IPRouteLookup) (*IPRouteLookupReply, error)
	IPSourceAndPortRangeCheckAddDel(ctx context.Context, in *IPSourceAndPortRangeCheckAddDel) (*IPSourceAndPortRangeCheckAddDelReply, error)
	IPSourceAndPortRangeCheckInterfaceAddDel(ctx context.Context, in *IPSourceAndPortRangeCheckInterfaceAddDel) (*IPSourceAndPortRangeCheckInterfaceAddDelReply, error)
	IPTableAddDel(ctx context.Context, in *IPTableAddDel) (*IPTableAddDelReply, error)
	IPTableDump(ctx context.Context, in *IPTableDump) (RPCService_IPTableDumpClient, error)
	IPTableFlush(ctx context.Context, in *IPTableFlush) (*IPTableFlushReply, error)
	IPTableReplaceBegin(ctx context.Context, in *IPTableReplaceBegin) (*IPTableReplaceBeginReply, error)
	IPTableReplaceEnd(ctx context.Context, in *IPTableReplaceEnd) (*IPTableReplaceEndReply, error)
	IPUnnumberedDump(ctx context.Context, in *IPUnnumberedDump) (RPCService_IPUnnumberedDumpClient, error)
	MfibSignalDump(ctx context.Context, in *MfibSignalDump) (RPCService_MfibSignalDumpClient, error)
	SetIPFlowHash(ctx context.Context, in *SetIPFlowHash) (*SetIPFlowHashReply, error)
	SwInterfaceIP6EnableDisable(ctx context.Context, in *SwInterfaceIP6EnableDisable) (*SwInterfaceIP6EnableDisableReply, error)
	SwInterfaceIP6SetLinkLocalAddress(ctx context.Context, in *SwInterfaceIP6SetLinkLocalAddress) (*SwInterfaceIP6SetLinkLocalAddressReply, error)
}

type serviceClient struct {
	conn api.Connection
}

func NewServiceClient(conn api.Connection) RPCService {
	return &serviceClient{conn}
}

func (c *serviceClient) IoamDisable(ctx context.Context, in *IoamDisable) (*IoamDisableReply, error) {
	out := new(IoamDisableReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) IoamEnable(ctx context.Context, in *IoamEnable) (*IoamEnableReply, error) {
	out := new(IoamEnableReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) IPAddressDump(ctx context.Context, in *IPAddressDump) (RPCService_IPAddressDumpClient, error) {
	stream, err := c.conn.NewStream(ctx)
	if err != nil {
		return nil, err
	}
	x := &serviceClient_IPAddressDumpClient{stream}
	if err := x.Stream.SendMsg(in); err != nil {
		return nil, err
	}
	if err = x.Stream.SendMsg(&vpe.ControlPing{}); err != nil {
		return nil, err
	}
	return x, nil
}

type RPCService_IPAddressDumpClient interface {
	Recv() (*IPAddressDetails, error)
	api.Stream
}

type serviceClient_IPAddressDumpClient struct {
	api.Stream
}

func (c *serviceClient_IPAddressDumpClient) Recv() (*IPAddressDetails, error) {
	msg, err := c.Stream.RecvMsg()
	if err != nil {
		return nil, err
	}
	switch m := msg.(type) {
	case *IPAddressDetails:
		return m, nil
	case *vpe.ControlPingReply:
		return nil, io.EOF
	default:
		return nil, fmt.Errorf("unexpected message: %T %v", m, m)
	}
}

func (c *serviceClient) IPContainerProxyAddDel(ctx context.Context, in *IPContainerProxyAddDel) (*IPContainerProxyAddDelReply, error) {
	out := new(IPContainerProxyAddDelReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) IPContainerProxyDump(ctx context.Context, in *IPContainerProxyDump) (RPCService_IPContainerProxyDumpClient, error) {
	stream, err := c.conn.NewStream(ctx)
	if err != nil {
		return nil, err
	}
	x := &serviceClient_IPContainerProxyDumpClient{stream}
	if err := x.Stream.SendMsg(in); err != nil {
		return nil, err
	}
	if err = x.Stream.SendMsg(&vpe.ControlPing{}); err != nil {
		return nil, err
	}
	return x, nil
}

type RPCService_IPContainerProxyDumpClient interface {
	Recv() (*IPContainerProxyDetails, error)
	api.Stream
}

type serviceClient_IPContainerProxyDumpClient struct {
	api.Stream
}

func (c *serviceClient_IPContainerProxyDumpClient) Recv() (*IPContainerProxyDetails, error) {
	msg, err := c.Stream.RecvMsg()
	if err != nil {
		return nil, err
	}
	switch m := msg.(type) {
	case *IPContainerProxyDetails:
		return m, nil
	case *vpe.ControlPingReply:
		return nil, io.EOF
	default:
		return nil, fmt.Errorf("unexpected message: %T %v", m, m)
	}
}

func (c *serviceClient) IPDump(ctx context.Context, in *IPDump) (RPCService_IPDumpClient, error) {
	stream, err := c.conn.NewStream(ctx)
	if err != nil {
		return nil, err
	}
	x := &serviceClient_IPDumpClient{stream}
	if err := x.Stream.SendMsg(in); err != nil {
		return nil, err
	}
	if err = x.Stream.SendMsg(&vpe.ControlPing{}); err != nil {
		return nil, err
	}
	return x, nil
}

type RPCService_IPDumpClient interface {
	Recv() (*IPDetails, error)
	api.Stream
}

type serviceClient_IPDumpClient struct {
	api.Stream
}

func (c *serviceClient_IPDumpClient) Recv() (*IPDetails, error) {
	msg, err := c.Stream.RecvMsg()
	if err != nil {
		return nil, err
	}
	switch m := msg.(type) {
	case *IPDetails:
		return m, nil
	case *vpe.ControlPingReply:
		return nil, io.EOF
	default:
		return nil, fmt.Errorf("unexpected message: %T %v", m, m)
	}
}

func (c *serviceClient) IPMrouteAddDel(ctx context.Context, in *IPMrouteAddDel) (*IPMrouteAddDelReply, error) {
	out := new(IPMrouteAddDelReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) IPMrouteDump(ctx context.Context, in *IPMrouteDump) (RPCService_IPMrouteDumpClient, error) {
	stream, err := c.conn.NewStream(ctx)
	if err != nil {
		return nil, err
	}
	x := &serviceClient_IPMrouteDumpClient{stream}
	if err := x.Stream.SendMsg(in); err != nil {
		return nil, err
	}
	if err = x.Stream.SendMsg(&vpe.ControlPing{}); err != nil {
		return nil, err
	}
	return x, nil
}

type RPCService_IPMrouteDumpClient interface {
	Recv() (*IPMrouteDetails, error)
	api.Stream
}

type serviceClient_IPMrouteDumpClient struct {
	api.Stream
}

func (c *serviceClient_IPMrouteDumpClient) Recv() (*IPMrouteDetails, error) {
	msg, err := c.Stream.RecvMsg()
	if err != nil {
		return nil, err
	}
	switch m := msg.(type) {
	case *IPMrouteDetails:
		return m, nil
	case *vpe.ControlPingReply:
		return nil, io.EOF
	default:
		return nil, fmt.Errorf("unexpected message: %T %v", m, m)
	}
}

func (c *serviceClient) IPMtableDump(ctx context.Context, in *IPMtableDump) (RPCService_IPMtableDumpClient, error) {
	stream, err := c.conn.NewStream(ctx)
	if err != nil {
		return nil, err
	}
	x := &serviceClient_IPMtableDumpClient{stream}
	if err := x.Stream.SendMsg(in); err != nil {
		return nil, err
	}
	if err = x.Stream.SendMsg(&vpe.ControlPing{}); err != nil {
		return nil, err
	}
	return x, nil
}

type RPCService_IPMtableDumpClient interface {
	Recv() (*IPMtableDetails, error)
	api.Stream
}

type serviceClient_IPMtableDumpClient struct {
	api.Stream
}

func (c *serviceClient_IPMtableDumpClient) Recv() (*IPMtableDetails, error) {
	msg, err := c.Stream.RecvMsg()
	if err != nil {
		return nil, err
	}
	switch m := msg.(type) {
	case *IPMtableDetails:
		return m, nil
	case *vpe.ControlPingReply:
		return nil, io.EOF
	default:
		return nil, fmt.Errorf("unexpected message: %T %v", m, m)
	}
}

func (c *serviceClient) IPPuntPolice(ctx context.Context, in *IPPuntPolice) (*IPPuntPoliceReply, error) {
	out := new(IPPuntPoliceReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) IPPuntRedirect(ctx context.Context, in *IPPuntRedirect) (*IPPuntRedirectReply, error) {
	out := new(IPPuntRedirectReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) IPPuntRedirectDump(ctx context.Context, in *IPPuntRedirectDump) (RPCService_IPPuntRedirectDumpClient, error) {
	stream, err := c.conn.NewStream(ctx)
	if err != nil {
		return nil, err
	}
	x := &serviceClient_IPPuntRedirectDumpClient{stream}
	if err := x.Stream.SendMsg(in); err != nil {
		return nil, err
	}
	if err = x.Stream.SendMsg(&vpe.ControlPing{}); err != nil {
		return nil, err
	}
	return x, nil
}

type RPCService_IPPuntRedirectDumpClient interface {
	Recv() (*IPPuntRedirectDetails, error)
	api.Stream
}

type serviceClient_IPPuntRedirectDumpClient struct {
	api.Stream
}

func (c *serviceClient_IPPuntRedirectDumpClient) Recv() (*IPPuntRedirectDetails, error) {
	msg, err := c.Stream.RecvMsg()
	if err != nil {
		return nil, err
	}
	switch m := msg.(type) {
	case *IPPuntRedirectDetails:
		return m, nil
	case *vpe.ControlPingReply:
		return nil, io.EOF
	default:
		return nil, fmt.Errorf("unexpected message: %T %v", m, m)
	}
}

func (c *serviceClient) IPReassemblyEnableDisable(ctx context.Context, in *IPReassemblyEnableDisable) (*IPReassemblyEnableDisableReply, error) {
	out := new(IPReassemblyEnableDisableReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) IPReassemblyGet(ctx context.Context, in *IPReassemblyGet) (*IPReassemblyGetReply, error) {
	out := new(IPReassemblyGetReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) IPReassemblySet(ctx context.Context, in *IPReassemblySet) (*IPReassemblySetReply, error) {
	out := new(IPReassemblySetReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) IPRouteAddDel(ctx context.Context, in *IPRouteAddDel) (*IPRouteAddDelReply, error) {
	out := new(IPRouteAddDelReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) IPRouteDump(ctx context.Context, in *IPRouteDump) (RPCService_IPRouteDumpClient, error) {
	stream, err := c.conn.NewStream(ctx)
	if err != nil {
		return nil, err
	}
	x := &serviceClient_IPRouteDumpClient{stream}
	if err := x.Stream.SendMsg(in); err != nil {
		return nil, err
	}
	if err = x.Stream.SendMsg(&vpe.ControlPing{}); err != nil {
		return nil, err
	}
	return x, nil
}

type RPCService_IPRouteDumpClient interface {
	Recv() (*IPRouteDetails, error)
	api.Stream
}

type serviceClient_IPRouteDumpClient struct {
	api.Stream
}

func (c *serviceClient_IPRouteDumpClient) Recv() (*IPRouteDetails, error) {
	msg, err := c.Stream.RecvMsg()
	if err != nil {
		return nil, err
	}
	switch m := msg.(type) {
	case *IPRouteDetails:
		return m, nil
	case *vpe.ControlPingReply:
		return nil, io.EOF
	default:
		return nil, fmt.Errorf("unexpected message: %T %v", m, m)
	}
}

func (c *serviceClient) IPRouteLookup(ctx context.Context, in *IPRouteLookup) (*IPRouteLookupReply, error) {
	out := new(IPRouteLookupReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) IPSourceAndPortRangeCheckAddDel(ctx context.Context, in *IPSourceAndPortRangeCheckAddDel) (*IPSourceAndPortRangeCheckAddDelReply, error) {
	out := new(IPSourceAndPortRangeCheckAddDelReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) IPSourceAndPortRangeCheckInterfaceAddDel(ctx context.Context, in *IPSourceAndPortRangeCheckInterfaceAddDel) (*IPSourceAndPortRangeCheckInterfaceAddDelReply, error) {
	out := new(IPSourceAndPortRangeCheckInterfaceAddDelReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) IPTableAddDel(ctx context.Context, in *IPTableAddDel) (*IPTableAddDelReply, error) {
	out := new(IPTableAddDelReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) IPTableDump(ctx context.Context, in *IPTableDump) (RPCService_IPTableDumpClient, error) {
	stream, err := c.conn.NewStream(ctx)
	if err != nil {
		return nil, err
	}
	x := &serviceClient_IPTableDumpClient{stream}
	if err := x.Stream.SendMsg(in); err != nil {
		return nil, err
	}
	if err = x.Stream.SendMsg(&vpe.ControlPing{}); err != nil {
		return nil, err
	}
	return x, nil
}

type RPCService_IPTableDumpClient interface {
	Recv() (*IPTableDetails, error)
	api.Stream
}

type serviceClient_IPTableDumpClient struct {
	api.Stream
}

func (c *serviceClient_IPTableDumpClient) Recv() (*IPTableDetails, error) {
	msg, err := c.Stream.RecvMsg()
	if err != nil {
		return nil, err
	}
	switch m := msg.(type) {
	case *IPTableDetails:
		return m, nil
	case *vpe.ControlPingReply:
		return nil, io.EOF
	default:
		return nil, fmt.Errorf("unexpected message: %T %v", m, m)
	}
}

func (c *serviceClient) IPTableFlush(ctx context.Context, in *IPTableFlush) (*IPTableFlushReply, error) {
	out := new(IPTableFlushReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) IPTableReplaceBegin(ctx context.Context, in *IPTableReplaceBegin) (*IPTableReplaceBeginReply, error) {
	out := new(IPTableReplaceBeginReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) IPTableReplaceEnd(ctx context.Context, in *IPTableReplaceEnd) (*IPTableReplaceEndReply, error) {
	out := new(IPTableReplaceEndReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) IPUnnumberedDump(ctx context.Context, in *IPUnnumberedDump) (RPCService_IPUnnumberedDumpClient, error) {
	stream, err := c.conn.NewStream(ctx)
	if err != nil {
		return nil, err
	}
	x := &serviceClient_IPUnnumberedDumpClient{stream}
	if err := x.Stream.SendMsg(in); err != nil {
		return nil, err
	}
	if err = x.Stream.SendMsg(&vpe.ControlPing{}); err != nil {
		return nil, err
	}
	return x, nil
}

type RPCService_IPUnnumberedDumpClient interface {
	Recv() (*IPUnnumberedDetails, error)
	api.Stream
}

type serviceClient_IPUnnumberedDumpClient struct {
	api.Stream
}

func (c *serviceClient_IPUnnumberedDumpClient) Recv() (*IPUnnumberedDetails, error) {
	msg, err := c.Stream.RecvMsg()
	if err != nil {
		return nil, err
	}
	switch m := msg.(type) {
	case *IPUnnumberedDetails:
		return m, nil
	case *vpe.ControlPingReply:
		return nil, io.EOF
	default:
		return nil, fmt.Errorf("unexpected message: %T %v", m, m)
	}
}

func (c *serviceClient) MfibSignalDump(ctx context.Context, in *MfibSignalDump) (RPCService_MfibSignalDumpClient, error) {
	stream, err := c.conn.NewStream(ctx)
	if err != nil {
		return nil, err
	}
	x := &serviceClient_MfibSignalDumpClient{stream}
	if err := x.Stream.SendMsg(in); err != nil {
		return nil, err
	}
	if err = x.Stream.SendMsg(&vpe.ControlPing{}); err != nil {
		return nil, err
	}
	return x, nil
}

type RPCService_MfibSignalDumpClient interface {
	Recv() (*MfibSignalDetails, error)
	api.Stream
}

type serviceClient_MfibSignalDumpClient struct {
	api.Stream
}

func (c *serviceClient_MfibSignalDumpClient) Recv() (*MfibSignalDetails, error) {
	msg, err := c.Stream.RecvMsg()
	if err != nil {
		return nil, err
	}
	switch m := msg.(type) {
	case *MfibSignalDetails:
		return m, nil
	case *vpe.ControlPingReply:
		return nil, io.EOF
	default:
		return nil, fmt.Errorf("unexpected message: %T %v", m, m)
	}
}

func (c *serviceClient) SetIPFlowHash(ctx context.Context, in *SetIPFlowHash) (*SetIPFlowHashReply, error) {
	out := new(SetIPFlowHashReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) SwInterfaceIP6EnableDisable(ctx context.Context, in *SwInterfaceIP6EnableDisable) (*SwInterfaceIP6EnableDisableReply, error) {
	out := new(SwInterfaceIP6EnableDisableReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) SwInterfaceIP6SetLinkLocalAddress(ctx context.Context, in *SwInterfaceIP6SetLinkLocalAddress) (*SwInterfaceIP6SetLinkLocalAddressReply, error) {
	out := new(SwInterfaceIP6SetLinkLocalAddressReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}
