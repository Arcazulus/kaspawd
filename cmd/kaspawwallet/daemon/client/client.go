package client

import (
	"context"
	"github.com/Arcazulus/kaspawd/cmd/kaspawwallet/daemon/server"
	"time"

	"github.com/pkg/errors"

	"github.com/Arcazulus/kaspawd/cmd/kaspawwallet/daemon/pb"
	"google.golang.org/grpc"
)

// Connect connects to the kaspawwalletd server, and returns the client instance
func Connect(address string) (pb.KaspawwalletdClient, func(), error) {
	// Connection is local, so 1 second timeout is sufficient
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	conn, err := grpc.DialContext(ctx, address, grpc.WithInsecure(), grpc.WithBlock(), grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(server.MaxDaemonSendMsgSize)))
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			return nil, nil, errors.New("kaspawwallet daemon is not running, start it with `kaspawwallet start-daemon`")
		}
		return nil, nil, err
	}

	return pb.NewKaspawwalletdClient(conn), func() {
		conn.Close()
	}, nil
}
