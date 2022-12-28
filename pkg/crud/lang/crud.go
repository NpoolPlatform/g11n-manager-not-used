package lang

import (
	"context"
	"fmt"
	"time"

	constant "github.com/NpoolPlatform/g11n-manager/pkg/message/const"
	commontracer "github.com/NpoolPlatform/g11n-manager/pkg/tracer"
	tracer "github.com/NpoolPlatform/g11n-manager/pkg/tracer/lang"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/codes"

	"github.com/NpoolPlatform/g11n-manager/pkg/db"
	"github.com/NpoolPlatform/g11n-manager/pkg/db/ent"
	"github.com/NpoolPlatform/g11n-manager/pkg/db/ent/lang"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/g11n/mgr/v1/lang"

	"github.com/google/uuid"
)

func CreateSet(c *ent.LangCreate, in *npool.LangReq) *ent.LangCreate {
	if in.ID != nil {
		c.SetID(uuid.MustParse(in.GetID()))
	}
	if in.Lang != nil {
		c.SetLang(in.GetLang())
	}
	if in.Logo != nil {
		c.SetLogo(in.GetLogo())
	}
	if in.Name != nil {
		c.SetName(in.GetName())
	}
	if in.Short != nil {
		c.SetShort(in.GetShort())
	}
	return c
}

func Create(ctx context.Context, in *npool.LangReq) (*ent.Lang, error) {
	var info *ent.Lang
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "Create")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(codes.Error, "db operation fail")
			span.RecordError(err)
		}
	}()

	span = tracer.Trace(span, in)

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		c := CreateSet(cli.Lang.Create(), in)
		info, err = c.Save(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}

func CreateBulk(ctx context.Context, in []*npool.LangReq) ([]*ent.Lang, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CreateBulk")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(codes.Error, "db operation fail")
			span.RecordError(err)
		}
	}()

	span = tracer.TraceMany(span, in)

	rows := []*ent.Lang{}
	err = db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		bulk := make([]*ent.LangCreate, len(in))
		for i, info := range in {
			bulk[i] = CreateSet(tx.Lang.Create(), info)
		}
		rows, err = tx.Lang.CreateBulk(bulk...).Save(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}
	return rows, nil
}

func UpdateSet(info *ent.Lang, in *npool.LangReq) *ent.LangUpdateOne {
	stm := info.Update()

	if in.Lang != nil {
		stm = stm.SetLang(in.GetLang())
	}
	if in.Logo != nil {
		stm = stm.SetLogo(in.GetLogo())
	}
	if in.Name != nil {
		stm = stm.SetName(in.GetName())
	}
	if in.Short != nil {
		stm = stm.SetShort(in.GetShort())
	}

	return stm
}

func Update(ctx context.Context, in *npool.LangReq) (*ent.Lang, error) {
	var info *ent.Lang
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "Create")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(codes.Error, "db operation fail")
			span.RecordError(err)
		}
	}()

	span = tracer.Trace(span, in)

	err = db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		info, err = tx.Lang.Query().Where(lang.ID(uuid.MustParse(in.GetID()))).ForUpdate().Only(_ctx)
		if err != nil {
			return fmt.Errorf("fail query lang: %v", err)
		}

		stm := UpdateSet(info, in)

		info, err = stm.Save(_ctx)
		if err != nil {
			return fmt.Errorf("fail update lang: %v", err)
		}

		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail update lang: %v", err)
	}

	return info, nil
}

func Row(ctx context.Context, id uuid.UUID) (*ent.Lang, error) {
	var info *ent.Lang
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "Row")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(codes.Error, "db operation fail")
			span.RecordError(err)
		}
	}()

	span = commontracer.TraceID(span, id.String())

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		info, err = cli.Lang.Query().Where(lang.ID(id)).Only(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}

func SetQueryConds(conds *npool.Conds, cli *ent.Client) (*ent.LangQuery, error) {
	stm := cli.Lang.Query()
	if conds.ID != nil {
		switch conds.GetID().GetOp() {
		case cruder.EQ:
			stm.Where(lang.ID(uuid.MustParse(conds.GetID().GetValue())))
		case cruder.NEQ:
			stm.Where(lang.IDNEQ(uuid.MustParse(conds.GetID().GetValue())))
		default:
			return nil, fmt.Errorf("invalid lang field")
		}
	}
	if conds.Lang != nil {
		switch conds.GetLang().GetOp() {
		case cruder.EQ:
			stm.Where(lang.Lang(conds.GetLang().GetValue()))
		case cruder.LIKE:
			stm.Where(lang.LangContains(conds.GetLang().GetValue()))
		default:
			return nil, fmt.Errorf("invalid lang field")
		}
	}
	if conds.Name != nil {
		switch conds.GetName().GetOp() {
		case cruder.LIKE:
			stm.Where(lang.Name(conds.GetName().GetValue()))
		default:
			return nil, fmt.Errorf("invalid lang field")
		}
	}
	if conds.Short != nil {
		switch conds.GetShort().GetOp() {
		case cruder.LIKE:
			stm.Where(lang.Short(conds.GetShort().GetValue()))
		default:
			return nil, fmt.Errorf("invalid lang field")
		}
	}
	if conds.Langs != nil {
		switch conds.GetLangs().GetOp() {
		case cruder.IN:
			stm.Where(lang.LangIn(conds.GetLangs().GetValue()...))
		default:
			return nil, fmt.Errorf("invalid lang field")
		}
	}
	return stm, nil
}

func Rows(ctx context.Context, conds *npool.Conds, offset, limit int) ([]*ent.Lang, int, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "Rows")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(codes.Error, "db operation fail")
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, conds)
	span = commontracer.TraceOffsetLimit(span, offset, limit)

	rows := []*ent.Lang{}
	var total int
	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		stm, err := SetQueryConds(conds, cli)
		if err != nil {
			return err
		}

		total, err = stm.Count(_ctx)
		if err != nil {
			return err
		}

		rows, err = stm.
			Offset(offset).
			Order(ent.Desc(lang.FieldUpdatedAt)).
			Limit(limit).
			All(_ctx)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return nil, 0, err
	}
	return rows, total, nil
}

func RowOnly(ctx context.Context, conds *npool.Conds) (*ent.Lang, error) {
	var info *ent.Lang
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "RowOnly")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(codes.Error, "db operation fail")
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, conds)

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		stm, err := SetQueryConds(conds, cli)
		if err != nil {
			return err
		}

		info, err = stm.Only(_ctx)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}

func Count(ctx context.Context, conds *npool.Conds) (uint32, error) {
	var err error
	var total int

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "Count")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(codes.Error, "db operation fail")
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, conds)

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		stm, err := SetQueryConds(conds, cli)
		if err != nil {
			return err
		}

		total, err = stm.Count(_ctx)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return 0, err
	}

	return uint32(total), nil
}

func Exist(ctx context.Context, id uuid.UUID) (bool, error) {
	var err error
	exist := false

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "Exist")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(codes.Error, "db operation fail")
			span.RecordError(err)
		}
	}()

	span = commontracer.TraceID(span, id.String())

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		exist, err = cli.Lang.Query().Where(lang.ID(id)).Exist(_ctx)
		return err
	})
	if err != nil {
		return false, err
	}

	return exist, nil
}

func ExistConds(ctx context.Context, conds *npool.Conds) (bool, error) {
	var err error
	exist := false

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "ExistConds")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(codes.Error, "db operation fail")
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, conds)

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		stm, err := SetQueryConds(conds, cli)
		if err != nil {
			return err
		}

		exist, err = stm.Exist(_ctx)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return false, err
	}

	return exist, nil
}

func Delete(ctx context.Context, id uuid.UUID) (*ent.Lang, error) {
	var info *ent.Lang
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "Delete")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(codes.Error, "db operation fail")
			span.RecordError(err)
		}
	}()

	span = commontracer.TraceID(span, id.String())

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		info, err = cli.Lang.UpdateOneID(id).
			SetDeletedAt(uint32(time.Now().Unix())).
			Save(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}
