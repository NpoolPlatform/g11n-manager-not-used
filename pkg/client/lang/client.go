//nolint:dupl
package lang

import (
	"context"
	"fmt"
	"time"

	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/g11n/mgr/v1/lang"

	constant "github.com/NpoolPlatform/g11n-manager/pkg/message/const"
)

var timeout = 10 * time.Second

type handler func(context.Context, npool.ManagerClient) (cruder.Any, error)

func withCRUD(ctx context.Context, handler handler) (cruder.Any, error) {
	_ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	conn, err := grpc2.GetGRPCConn(constant.ServiceName, grpc2.GRPCTAG)
	if err != nil {
		return nil, fmt.Errorf("fail get lang connection: %v", err)
	}

	defer conn.Close()

	cli := npool.NewManagerClient(conn)

	return handler(_ctx, cli)
}

func CreateLang(ctx context.Context, in *npool.LangReq) (*npool.Lang, error) {
	info, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.CreateLang(ctx, &npool.CreateLangRequest{
			Info: in,
		})
		if err != nil {
			return nil, fmt.Errorf("fail create lang: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail create lang: %v", err)
	}
	return info.(*npool.Lang), nil
}

func CreateLangs(ctx context.Context, in []*npool.LangReq) ([]*npool.Lang, error) {
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.CreateLangs(ctx, &npool.CreateLangsRequest{
			Infos: in,
		})
		if err != nil {
			return nil, fmt.Errorf("fail create langs: %v", err)
		}
		return resp.GetInfos(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail create langs: %v", err)
	}
	return infos.([]*npool.Lang), nil
}

func UpdateLang(ctx context.Context, in *npool.LangReq) (*npool.Lang, error) {
	info, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.UpdateLang(ctx, &npool.UpdateLangRequest{
			Info: in,
		})
		if err != nil {
			return nil, fmt.Errorf("fail update lang: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail update lang: %v", err)
	}
	return info.(*npool.Lang), nil
}

func GetLang(ctx context.Context, id string) (*npool.Lang, error) {
	info, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.GetLang(ctx, &npool.GetLangRequest{
			ID: id,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get lang: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail get lang: %v", err)
	}
	return info.(*npool.Lang), nil
}

func GetLangOnly(ctx context.Context, conds *npool.Conds) (*npool.Lang, error) {
	info, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.GetLangOnly(ctx, &npool.GetLangOnlyRequest{
			Conds: conds,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get lang: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail get lang: %v", err)
	}
	return info.(*npool.Lang), nil
}

func GetLangs(ctx context.Context, conds *npool.Conds, offset, limit int32) ([]*npool.Lang, uint32, error) {
	var total uint32
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.GetLangs(ctx, &npool.GetLangsRequest{
			Conds:  conds,
			Limit:  limit,
			Offset: offset,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get langs: %v", err)
		}
		total = resp.GetTotal()
		return resp.GetInfos(), nil
	})
	if err != nil {
		return nil, 0, fmt.Errorf("fail get langs: %v", err)
	}
	return infos.([]*npool.Lang), total, nil
}

func ExistLang(ctx context.Context, id string) (bool, error) {
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.ExistLang(ctx, &npool.ExistLangRequest{
			ID: id,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get lang: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return false, fmt.Errorf("fail get lang: %v", err)
	}
	return infos.(bool), nil
}

func ExistLangConds(ctx context.Context, conds *npool.Conds) (bool, error) {
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.ExistLangConds(ctx, &npool.ExistLangCondsRequest{
			Conds: conds,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get lang: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return false, fmt.Errorf("fail get lang: %v", err)
	}
	return infos.(bool), nil
}

func CountLangs(ctx context.Context, conds *npool.Conds) (uint32, error) {
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.CountLangs(ctx, &npool.CountLangsRequest{
			Conds: conds,
		})
		if err != nil {
			return nil, fmt.Errorf("fail count lang: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return 0, fmt.Errorf("fail count lang: %v", err)
	}
	return infos.(uint32), nil
}
