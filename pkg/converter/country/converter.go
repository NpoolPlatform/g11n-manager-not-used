package country

import (
	"github.com/NpoolPlatform/g11n-manager/pkg/db/ent"
	npool "github.com/NpoolPlatform/message/npool/g11n/mgr/v1/country"
)

func Ent2Grpc(row *ent.Country) *npool.Country {
	if row == nil {
		return nil
	}

	return &npool.Country{
		ID:      row.ID.String(),
		Country: row.Country,
		Flag:    row.Flag,
		Code:    row.Code,
		Short:   row.Short,
	}
}

func Ent2GrpcMany(rows []*ent.Country) []*npool.Country {
	infos := []*npool.Country{}
	for _, row := range rows {
		infos = append(infos, Ent2Grpc(row))
	}
	return infos
}
