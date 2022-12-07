package message

import (
	"github.com/NpoolPlatform/g11n-manager/pkg/db/ent"
	npool "github.com/NpoolPlatform/message/npool/g11n/mgr/v1/message"
)

func Ent2Grpc(row *ent.Message) *npool.Message {
	if row == nil {
		return nil
	}

	return &npool.Message{
		ID:        row.ID.String(),
		AppID:     row.AppID.String(),
		LangID:    row.LangID.String(),
		MessageID: row.MessageID,
		Message:   row.Message,
		GetIndex:  row.GetIndex,
		Disabled:  row.Disabled,
		CreatedAt: row.CreatedAt,
		UpdatedAt: row.UpdatedAt,
	}
}

func Ent2GrpcMany(rows []*ent.Message) []*npool.Message {
	infos := []*npool.Message{}
	for _, row := range rows {
		infos = append(infos, Ent2Grpc(row))
	}
	return infos
}
