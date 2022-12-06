//nolint:dupl
package message

import (
	"context"
	"fmt"
	"time"

	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/g11n/mgr/v1/message"

	constant "github.com/NpoolPlatform/g11n-manager/pkg/message/const"
)

var timeout = 10 * time.Second

type handler func(context.Context, npool.ManagerClient) (cruder.Any, error)

func withCRUD(ctx context.Context, handler handler) (cruder.Any, error) {
	_ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	conn, err := grpc2.GetGRPCConn(constant.ServiceName, grpc2.GRPCTAG)
	if err != nil {
		return nil, fmt.Errorf("fail get message connection: %v", err)
	}

	defer conn.Close()

	cli := npool.NewManagerClient(conn)

	return handler(_ctx, cli)
}

func CreateMessage(ctx context.Context, in *npool.MessageReq) (*npool.Message, error) {
	info, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.CreateMessage(ctx, &npool.CreateMessageRequest{
			Info: in,
		})
		if err != nil {
			return nil, fmt.Errorf("fail create message: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail create message: %v", err)
	}
	return info.(*npool.Message), nil
}

func CreateMessages(ctx context.Context, in []*npool.MessageReq) ([]*npool.Message, error) {
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.CreateMessages(ctx, &npool.CreateMessagesRequest{
			Infos: in,
		})
		if err != nil {
			return nil, fmt.Errorf("fail create messages: %v", err)
		}
		return resp.GetInfos(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail create messages: %v", err)
	}
	return infos.([]*npool.Message), nil
}

func GetMessage(ctx context.Context, id string) (*npool.Message, error) {
	info, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.GetMessage(ctx, &npool.GetMessageRequest{
			ID: id,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get message: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail get message: %v", err)
	}
	return info.(*npool.Message), nil
}

func GetMessageOnly(ctx context.Context, conds *npool.Conds) (*npool.Message, error) {
	info, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.GetMessageOnly(ctx, &npool.GetMessageOnlyRequest{
			Conds: conds,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get message: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail get message: %v", err)
	}
	return info.(*npool.Message), nil
}

func GetMessages(ctx context.Context, conds *npool.Conds, limit, offset int32) ([]*npool.Message, uint32, error) {
	var total uint32
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.GetMessages(ctx, &npool.GetMessagesRequest{
			Conds:  conds,
			Limit:  limit,
			Offset: offset,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get messages: %v", err)
		}
		total = resp.GetTotal()
		return resp.GetInfos(), nil
	})
	if err != nil {
		return nil, 0, fmt.Errorf("fail get messages: %v", err)
	}
	return infos.([]*npool.Message), total, nil
}

func ExistMessage(ctx context.Context, id string) (bool, error) {
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.ExistMessage(ctx, &npool.ExistMessageRequest{
			ID: id,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get message: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return false, fmt.Errorf("fail get message: %v", err)
	}
	return infos.(bool), nil
}

func ExistMessageConds(ctx context.Context, conds *npool.Conds) (bool, error) {
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.ExistMessageConds(ctx, &npool.ExistMessageCondsRequest{
			Conds: conds,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get message: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return false, fmt.Errorf("fail get message: %v", err)
	}
	return infos.(bool), nil
}

func CountMessages(ctx context.Context, conds *npool.Conds) (uint32, error) {
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.CountMessages(ctx, &npool.CountMessagesRequest{
			Conds: conds,
		})
		if err != nil {
			return nil, fmt.Errorf("fail count message: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return 0, fmt.Errorf("fail count message: %v", err)
	}
	return infos.(uint32), nil
}
