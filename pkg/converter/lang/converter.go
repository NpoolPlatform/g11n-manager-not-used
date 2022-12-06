package lang

import (
	"github.com/NpoolPlatform/g11n-manager/pkg/db/ent"
	npool "github.com/NpoolPlatform/message/npool/g11n/mgr/v1/lang"
)

func Ent2Grpc(row *ent.Lang) *npool.Lang {
	if row == nil {
		return nil
	}

	return &npool.Lang{
		ID:    row.ID.String(),
		Lang:  row.Lang,
		Logo:  row.Logo,
		Name:  row.Name,
		Short: row.Short,
	}
}

func Ent2GrpcMany(rows []*ent.Lang) []*npool.Lang {
	infos := []*npool.Lang{}
	for _, row := range rows {
		infos = append(infos, Ent2Grpc(row))
	}
	return infos
}
