// Code generated by csi-proxy-api-gen. DO NOT EDIT.

package v1beta2

import (
	"context"
	"net"

	"github.com/Microsoft/go-winio"
	"github.com/kubernetes-csi/csi-proxy/client"
	"github.com/kubernetes-csi/csi-proxy/client/api/disk/v1beta2"
	"github.com/kubernetes-csi/csi-proxy/client/apiversion"
	"google.golang.org/grpc"
)

const groupName = "disk"

var version = apiversion.NewVersionOrPanic("v1beta2")

type Client struct {
	client     v1beta2.DiskClient
	connection *grpc.ClientConn
}

// NewClient returns a client to make calls to the disk API group version v1beta2.
// It's the caller's responsibility to Close the client when done.
func NewClient() (*Client, error) {
	pipePath := client.PipePath(groupName, version)

	connection, err := grpc.Dial(pipePath,
		grpc.WithContextDialer(func(context context.Context, s string) (net.Conn, error) {
			return winio.DialPipeContext(context, s)
		}),
		grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	client := v1beta2.NewDiskClient(connection)
	return &Client{
		client:     client,
		connection: connection,
	}, nil
}

// Close closes the client. It must be called before the client gets GC-ed.
func (w *Client) Close() error {
	return w.connection.Close()
}

// ensures we implement all the required methods
var _ v1beta2.DiskClient = &Client{}

func (w *Client) DiskStats(context context.Context, request *v1beta2.DiskStatsRequest, opts ...grpc.CallOption) (*v1beta2.DiskStatsResponse, error) {
	return w.client.DiskStats(context, request, opts...)
}

func (w *Client) GetAttachState(context context.Context, request *v1beta2.GetAttachStateRequest, opts ...grpc.CallOption) (*v1beta2.GetAttachStateResponse, error) {
	return w.client.GetAttachState(context, request, opts...)
}

func (w *Client) ListDiskIDs(context context.Context, request *v1beta2.ListDiskIDsRequest, opts ...grpc.CallOption) (*v1beta2.ListDiskIDsResponse, error) {
	return w.client.ListDiskIDs(context, request, opts...)
}

func (w *Client) ListDiskLocations(context context.Context, request *v1beta2.ListDiskLocationsRequest, opts ...grpc.CallOption) (*v1beta2.ListDiskLocationsResponse, error) {
	return w.client.ListDiskLocations(context, request, opts...)
}

func (w *Client) PartitionDisk(context context.Context, request *v1beta2.PartitionDiskRequest, opts ...grpc.CallOption) (*v1beta2.PartitionDiskResponse, error) {
	return w.client.PartitionDisk(context, request, opts...)
}

func (w *Client) Rescan(context context.Context, request *v1beta2.RescanRequest, opts ...grpc.CallOption) (*v1beta2.RescanResponse, error) {
	return w.client.Rescan(context, request, opts...)
}

func (w *Client) SetAttachState(context context.Context, request *v1beta2.SetAttachStateRequest, opts ...grpc.CallOption) (*v1beta2.SetAttachStateResponse, error) {
	return w.client.SetAttachState(context, request, opts...)
}
