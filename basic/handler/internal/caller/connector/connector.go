// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.5
// Source: connector.proto

package connectorclient

import (
	"context"

	"github.com/nanachi-sh/susubot-code/basic/handler/pkg/protos/connector"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	BasicResponse   = connector.BasicResponse
	ConnectRequest  = connector.ConnectRequest
	ConnectResponse = connector.ConnectResponse
	Empty           = connector.Empty
	ReadResponse    = connector.ReadResponse
	WriteRequest    = connector.WriteRequest

	Connector interface {
		Connect(ctx context.Context, in *ConnectRequest, opts ...grpc.CallOption) (*ConnectResponse, error)
		Read(ctx context.Context, in *Empty, opts ...grpc.CallOption) (connector.Connector_ReadClient, error)
		Write(ctx context.Context, in *WriteRequest, opts ...grpc.CallOption) (*BasicResponse, error)
		Close(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*BasicResponse, error)
	}

	defaultConnector struct {
		cli zrpc.Client
	}
)

func NewConnector(cli zrpc.Client) Connector {
	return &defaultConnector{
		cli: cli,
	}
}

func (m *defaultConnector) Connect(ctx context.Context, in *ConnectRequest, opts ...grpc.CallOption) (*ConnectResponse, error) {
	client := connector.NewConnectorClient(m.cli.Conn())
	return client.Connect(ctx, in, opts...)
}

func (m *defaultConnector) Read(ctx context.Context, in *Empty, opts ...grpc.CallOption) (connector.Connector_ReadClient, error) {
	client := connector.NewConnectorClient(m.cli.Conn())
	return client.Read(ctx, in, opts...)
}

func (m *defaultConnector) Write(ctx context.Context, in *WriteRequest, opts ...grpc.CallOption) (*BasicResponse, error) {
	client := connector.NewConnectorClient(m.cli.Conn())
	return client.Write(ctx, in, opts...)
}

func (m *defaultConnector) Close(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*BasicResponse, error) {
	client := connector.NewConnectorClient(m.cli.Conn())
	return client.Close(ctx, in, opts...)
}
