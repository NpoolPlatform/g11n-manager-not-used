//nolint:nolintlint,dupl
package country

import (
	"context"

	converter "github.com/NpoolPlatform/g11n-manager/pkg/converter/country"
	crud "github.com/NpoolPlatform/g11n-manager/pkg/crud/country"
	commontracer "github.com/NpoolPlatform/g11n-manager/pkg/tracer"
	tracer "github.com/NpoolPlatform/g11n-manager/pkg/tracer/country"

	constant "github.com/NpoolPlatform/g11n-manager/pkg/message/const"

	"go.opentelemetry.io/otel"
	scodes "go.opentelemetry.io/otel/codes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/g11n/mgr/v1/country"

	"github.com/google/uuid"
)

func (s *Server) CreateDetail(ctx context.Context, in *npool.CreateDetailRequest) (*npool.CreateDetailResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CreateDetail")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.Trace(span, in.GetInfo())

	err = validate(in.GetInfo())
	if err != nil {
		return &npool.CreateDetailResponse{}, err
	}

	span = commontracer.TraceInvoker(span, "country", "crud", "Create")

	info, err := crud.Create(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorf("fail create country: %v", err.Error())
		return &npool.CreateDetailResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateDetailResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) CreateDetails(ctx context.Context, in *npool.CreateDetailsRequest) (*npool.CreateDetailsResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CreateDetails")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	if len(in.GetInfos()) == 0 {
		return &npool.CreateDetailsResponse{}, status.Error(codes.InvalidArgument, "Infos is empty")
	}

	err = duplicate(in.GetInfos())
	if err != nil {
		return &npool.CreateDetailsResponse{}, err
	}

	span = tracer.TraceMany(span, in.GetInfos())
	span = commontracer.TraceInvoker(span, "country", "crud", "CreateBulk")

	rows, err := crud.CreateBulk(ctx, in.GetInfos())
	if err != nil {
		logger.Sugar().Errorf("fail create countrys: %v", err)
		return &npool.CreateDetailsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateDetailsResponse{
		Infos: converter.Ent2GrpcMany(rows),
	}, nil
}

func (s *Server) GetDetail(ctx context.Context, in *npool.GetDetailRequest) (*npool.GetDetailResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetDetail")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = commontracer.TraceID(span, in.GetID())

	id, err := uuid.Parse(in.GetID())
	if err != nil {
		return &npool.GetDetailResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "country", "crud", "Row")

	info, err := crud.Row(ctx, id)
	if err != nil {
		logger.Sugar().Errorf("fail get country: %v", err)
		return &npool.GetDetailResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetDetailResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) GetDetailOnly(ctx context.Context, in *npool.GetDetailOnlyRequest) (*npool.GetDetailOnlyResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetDetailOnly")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, in.GetConds())
	span = commontracer.TraceInvoker(span, "country", "crud", "RowOnly")

	info, err := crud.RowOnly(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorf("fail get countrys: %v", err)
		return &npool.GetDetailOnlyResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetDetailOnlyResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) GetDetails(ctx context.Context, in *npool.GetDetailsRequest) (*npool.GetDetailsResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetDetails")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, in.GetConds())
	span = commontracer.TraceOffsetLimit(span, int(in.GetOffset()), int(in.GetLimit()))
	span = commontracer.TraceInvoker(span, "country", "crud", "Rows")

	rows, total, err := crud.Rows(ctx, in.GetConds(), int(in.GetOffset()), int(in.GetLimit()))
	if err != nil {
		logger.Sugar().Errorf("fail get countrys: %v", err)
		return &npool.GetDetailsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetDetailsResponse{
		Infos: converter.Ent2GrpcMany(rows),
		Total: uint32(total),
	}, nil
}

func (s *Server) ExistDetail(ctx context.Context, in *npool.ExistDetailRequest) (*npool.ExistDetailResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "ExistDetail")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = commontracer.TraceID(span, in.GetID())

	id, err := uuid.Parse(in.GetID())
	if err != nil {
		return &npool.ExistDetailResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "country", "crud", "Exist")

	exist, err := crud.Exist(ctx, id)
	if err != nil {
		logger.Sugar().Errorf("fail check country: %v", err)
		return &npool.ExistDetailResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistDetailResponse{
		Info: exist,
	}, nil
}

func (s *Server) ExistDetailConds(ctx context.Context,
	in *npool.ExistDetailCondsRequest) (*npool.ExistDetailCondsResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "ExistDetailConds")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, in.GetConds())
	span = commontracer.TraceInvoker(span, "country", "crud", "ExistConds")

	exist, err := crud.ExistConds(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorf("fail check country: %v", err)
		return &npool.ExistDetailCondsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistDetailCondsResponse{
		Info: exist,
	}, nil
}

func (s *Server) CountDetails(ctx context.Context, in *npool.CountDetailsRequest) (*npool.CountDetailsResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CountDetails")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, in.GetConds())
	span = commontracer.TraceInvoker(span, "country", "crud", "Count")

	total, err := crud.Count(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorf("fail count countrys: %v", err)
		return &npool.CountDetailsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CountDetailsResponse{
		Info: total,
	}, nil
}
