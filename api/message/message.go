//nolint:nolintlint,dupl
package message

import (
	"context"
	"fmt"

	converter "github.com/NpoolPlatform/g11n-manager/pkg/converter/message"
	crud "github.com/NpoolPlatform/g11n-manager/pkg/crud/message"
	commontracer "github.com/NpoolPlatform/g11n-manager/pkg/tracer"
	tracer "github.com/NpoolPlatform/g11n-manager/pkg/tracer/message"

	constant "github.com/NpoolPlatform/g11n-manager/pkg/message/const"

	"go.opentelemetry.io/otel"
	scodes "go.opentelemetry.io/otel/codes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/g11n/mgr/v1/message"

	"github.com/google/uuid"
)

func (s *Server) CreateMessage(ctx context.Context, in *npool.CreateMessageRequest) (*npool.CreateMessageResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CreateMessage")
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
		return &npool.CreateMessageResponse{}, err
	}

	span = commontracer.TraceInvoker(span, "message", "crud", "Create")

	info, err := crud.Create(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorf("fail create message: %v", err.Error())
		return &npool.CreateMessageResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateMessageResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) CreateMessages(ctx context.Context, in *npool.CreateMessagesRequest) (*npool.CreateMessagesResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CreateMessages")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	if len(in.GetInfos()) == 0 {
		return &npool.CreateMessagesResponse{}, status.Error(codes.InvalidArgument, "Infos is empty")
	}

	err = duplicate(in.GetInfos())
	if err != nil {
		return &npool.CreateMessagesResponse{}, err
	}

	span = tracer.TraceMany(span, in.GetInfos())
	span = commontracer.TraceInvoker(span, "message", "crud", "CreateBulk")

	rows, err := crud.CreateBulk(ctx, in.GetInfos())
	if err != nil {
		logger.Sugar().Errorf("fail create messages: %v", err)
		return &npool.CreateMessagesResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateMessagesResponse{
		Infos: converter.Ent2GrpcMany(rows),
	}, nil
}

func (s *Server) UpdateMessage(ctx context.Context, in *npool.UpdateMessageRequest) (*npool.UpdateMessageResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "UpdateMessage")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.Trace(span, in.GetInfo())

	if _, err := uuid.Parse(in.GetInfo().GetID()); err != nil {
		return &npool.UpdateMessageResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}
	if in.GetInfo().MessageID != nil && in.GetInfo().GetMessageID() == "" {
		return &npool.UpdateMessageResponse{}, status.Error(codes.InvalidArgument, "invalid messageid")
	}
	if in.GetInfo().Message != nil && in.GetInfo().GetMessage() == "" {
		return &npool.UpdateMessageResponse{}, status.Error(codes.InvalidArgument, "invalid message")
	}

	span = commontracer.TraceInvoker(span, "message", "crud", "Update")

	info, err := crud.Update(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorf("fail update message: %v", err.Error())
		return &npool.UpdateMessageResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.UpdateMessageResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) GetMessage(ctx context.Context, in *npool.GetMessageRequest) (*npool.GetMessageResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetMessage")
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
		return &npool.GetMessageResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "message", "crud", "Row")

	info, err := crud.Row(ctx, id)
	if err != nil {
		logger.Sugar().Errorf("fail get message: %v", err)
		return &npool.GetMessageResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetMessageResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func ValidateConds(ctx context.Context, conds *npool.Conds) error {
	if conds.ID != nil {
		if _, err := uuid.Parse(conds.GetID().GetValue()); err != nil {
			return err
		}
	}
	if conds.AppID != nil {
		if _, err := uuid.Parse(conds.GetAppID().GetValue()); err != nil {
			return err
		}
	}
	if conds.LangID != nil {
		if _, err := uuid.Parse(conds.GetLangID().GetValue()); err != nil {
			return err
		}
	}
	if conds.MessageID != nil && conds.GetMessageID().GetValue() == "" {
		return fmt.Errorf("messageid is invalid")
	}
	for _, id := range conds.GetMessageIDs().GetValue() {
		if id == "" {
			return fmt.Errorf("messageid is invalid")
		}
	}
	return nil
}

func (s *Server) GetMessageOnly(ctx context.Context, in *npool.GetMessageOnlyRequest) (*npool.GetMessageOnlyResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetMessageOnly")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	if err := ValidateConds(ctx, in.GetConds()); err != nil {
		return &npool.GetMessageOnlyResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = tracer.TraceConds(span, in.GetConds())
	span = commontracer.TraceInvoker(span, "message", "crud", "RowOnly")

	info, err := crud.RowOnly(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorf("fail get messages: %v", err)
		return &npool.GetMessageOnlyResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetMessageOnlyResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) GetMessages(ctx context.Context, in *npool.GetMessagesRequest) (*npool.GetMessagesResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetMessages")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	if err := ValidateConds(ctx, in.GetConds()); err != nil {
		return &npool.GetMessagesResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = tracer.TraceConds(span, in.GetConds())
	span = commontracer.TraceOffsetLimit(span, int(in.GetOffset()), int(in.GetLimit()))
	span = commontracer.TraceInvoker(span, "message", "crud", "Rows")

	rows, total, err := crud.Rows(ctx, in.GetConds(), int(in.GetOffset()), int(in.GetLimit()))
	if err != nil {
		logger.Sugar().Errorf("fail get messages: %v", err)
		return &npool.GetMessagesResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetMessagesResponse{
		Infos: converter.Ent2GrpcMany(rows),
		Total: uint32(total),
	}, nil
}

func (s *Server) ExistMessage(ctx context.Context, in *npool.ExistMessageRequest) (*npool.ExistMessageResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "ExistMessage")
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
		return &npool.ExistMessageResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "message", "crud", "Exist")

	exist, err := crud.Exist(ctx, id)
	if err != nil {
		logger.Sugar().Errorf("fail check message: %v", err)
		return &npool.ExistMessageResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistMessageResponse{
		Info: exist,
	}, nil
}

func (s *Server) ExistMessageConds(ctx context.Context,
	in *npool.ExistMessageCondsRequest) (*npool.ExistMessageCondsResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "ExistMessageConds")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	if err := ValidateConds(ctx, in.GetConds()); err != nil {
		return &npool.ExistMessageCondsResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = tracer.TraceConds(span, in.GetConds())
	span = commontracer.TraceInvoker(span, "message", "crud", "ExistConds")

	exist, err := crud.ExistConds(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorf("fail check message: %v", err)
		return &npool.ExistMessageCondsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistMessageCondsResponse{
		Info: exist,
	}, nil
}

func (s *Server) CountMessages(ctx context.Context, in *npool.CountMessagesRequest) (*npool.CountMessagesResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CountMessages")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	if err := ValidateConds(ctx, in.GetConds()); err != nil {
		return &npool.CountMessagesResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = tracer.TraceConds(span, in.GetConds())
	span = commontracer.TraceInvoker(span, "message", "crud", "Count")

	total, err := crud.Count(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorf("fail count messages: %v", err)
		return &npool.CountMessagesResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CountMessagesResponse{
		Info: total,
	}, nil
}

func (s *Server) DeleteMessage(ctx context.Context, in *npool.DeleteMessageRequest) (*npool.DeleteMessageResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "DeleteMessage")
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
		return &npool.DeleteMessageResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "message", "crud", "Delete")

	info, err := crud.Delete(ctx, id)
	if err != nil {
		logger.Sugar().Errorf("fail delete message: %v", err)
		return &npool.DeleteMessageResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.DeleteMessageResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}
