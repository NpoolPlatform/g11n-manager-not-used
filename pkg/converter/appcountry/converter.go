package appcountry

import (
	"github.com/NpoolPlatform/g11n-manager/pkg/db/ent"
	npool "github.com/NpoolPlatform/message/npool/g11n/mgr/v1/appcountry"
)

func Ent2Grpc(row *ent.AppCountry) *npool.Country {
	if row == nil {
		return nil
	}

	return &npool.Country{
		ID:        row.ID.String(),
		AppID:     row.ID.String(),
		CountryID: row.ID.String(),
		CreatedAt: row.CreatedAt,
		UpdatedAt: row.UpdatedAt,
	}
}

func Ent2GrpcMany(rows []*ent.AppCountry) []*npool.Country {
	infos := []*npool.Country{}
	for _, row := range rows {
		infos = append(infos, Ent2Grpc(row))
	}
	return infos
}
