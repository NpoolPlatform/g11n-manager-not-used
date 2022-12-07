package applang

import (
	"github.com/NpoolPlatform/g11n-manager/pkg/db/ent"
	npool "github.com/NpoolPlatform/message/npool/g11n/mgr/v1/applang"
)

func Ent2Grpc(row *ent.AppLang) *npool.Lang {
	if row == nil {
		return nil
	}

	return &npool.Lang{
		ID:        row.ID.String(),
		AppID:     row.ID.String(),
		LangID:    row.ID.String(),
		Main:      row.Main,
		CreatedAt: row.CreatedAt,
		UpdatedAt: row.UpdatedAt,
	}
}

func Ent2GrpcMany(rows []*ent.AppLang) []*npool.Lang {
	infos := []*npool.Lang{}
	for _, row := range rows {
		infos = append(infos, Ent2Grpc(row))
	}
	return infos
}
