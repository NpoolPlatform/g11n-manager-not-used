//nolint:dupl
package country

import (
	"context"
	"fmt"
	"time"

	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/g11n/mgr/v1/country"

	constant "github.com/NpoolPlatform/g11n-manager/pkg/message/const"
)

var timeout = 10 * time.Second

type handler func(context.Context, npool.ManagerClient) (cruder.Any, error)

func withCRUD(ctx context.Context, handler handler) (cruder.Any, error) {
	_ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	conn, err := grpc2.GetGRPCConn(constant.ServiceName, grpc2.GRPCTAG)
	if err != nil {
		return nil, fmt.Errorf("fail get country connection: %v", err)
	}

	defer conn.Close()

	cli := npool.NewManagerClient(conn)

	return handler(_ctx, cli)
}

func CreateCountry(ctx context.Context, in *npool.CountryReq) (*npool.Country, error) {
	info, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.CreateCountry(ctx, &npool.CreateCountryRequest{
			Info: in,
		})
		if err != nil {
			return nil, fmt.Errorf("fail create country: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail create country: %v", err)
	}
	return info.(*npool.Country), nil
}

func CreateCountries(ctx context.Context, in []*npool.CountryReq) ([]*npool.Country, error) {
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.CreateCountries(ctx, &npool.CreateCountriesRequest{
			Infos: in,
		})
		if err != nil {
			return nil, fmt.Errorf("fail create countrys: %v", err)
		}
		return resp.GetInfos(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail create countrys: %v", err)
	}
	return infos.([]*npool.Country), nil
}

func GetCountry(ctx context.Context, id string) (*npool.Country, error) {
	info, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.GetCountry(ctx, &npool.GetCountryRequest{
			ID: id,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get country: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail get country: %v", err)
	}
	return info.(*npool.Country), nil
}

func GetCountryOnly(ctx context.Context, conds *npool.Conds) (*npool.Country, error) {
	info, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.GetCountryOnly(ctx, &npool.GetCountryOnlyRequest{
			Conds: conds,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get country: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail get country: %v", err)
	}
	return info.(*npool.Country), nil
}

func GetCountries(ctx context.Context, conds *npool.Conds, limit, offset int32) ([]*npool.Country, uint32, error) {
	var total uint32
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.GetCountries(ctx, &npool.GetCountriesRequest{
			Conds:  conds,
			Limit:  limit,
			Offset: offset,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get countrys: %v", err)
		}
		total = resp.GetTotal()
		return resp.GetInfos(), nil
	})
	if err != nil {
		return nil, 0, fmt.Errorf("fail get countrys: %v", err)
	}
	return infos.([]*npool.Country), total, nil
}

func ExistCountry(ctx context.Context, id string) (bool, error) {
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.ExistCountry(ctx, &npool.ExistCountryRequest{
			ID: id,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get country: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return false, fmt.Errorf("fail get country: %v", err)
	}
	return infos.(bool), nil
}

func ExistCountryConds(ctx context.Context, conds *npool.Conds) (bool, error) {
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.ExistCountryConds(ctx, &npool.ExistCountryCondsRequest{
			Conds: conds,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get country: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return false, fmt.Errorf("fail get country: %v", err)
	}
	return infos.(bool), nil
}

func CountCountries(ctx context.Context, conds *npool.Conds) (uint32, error) {
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.CountCountries(ctx, &npool.CountCountriesRequest{
			Conds: conds,
		})
		if err != nil {
			return nil, fmt.Errorf("fail count country: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return 0, fmt.Errorf("fail count country: %v", err)
	}
	return infos.(uint32), nil
}
