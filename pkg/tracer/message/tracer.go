package message

import (
	"fmt"

	"go.opentelemetry.io/otel/attribute"
	trace1 "go.opentelemetry.io/otel/trace"

	npool "github.com/NpoolPlatform/message/npool/g11n/mgr/v1/message"
)

func trace(span trace1.Span, in *npool.MessageReq, index int) trace1.Span {
	span.SetAttributes(
		attribute.String(fmt.Sprintf("ID.%v", index), in.GetID()),
		attribute.String(fmt.Sprintf("AppID.%v", index), in.GetAppID()),
		attribute.String(fmt.Sprintf("LangID.%v", index), in.GetLangID()),
		attribute.String(fmt.Sprintf("MessageID.%v", index), in.GetMessageID()),
		attribute.String(fmt.Sprintf("Message.%v", index), in.GetMessage()),
		attribute.Int(fmt.Sprintf("GetIndex.%v", index), int(in.GetGetIndex())),
		attribute.Bool(fmt.Sprintf("Disabled.%v", index), in.GetDisabled()),
	)
	return span
}

func Trace(span trace1.Span, in *npool.MessageReq) trace1.Span {
	return trace(span, in, 0)
}

func TraceConds(span trace1.Span, in *npool.Conds) trace1.Span {
	span.SetAttributes(
		attribute.String("ID.Op", in.GetID().GetOp()),
		attribute.String("ID.Value", in.GetID().GetValue()),
		attribute.String("AppID.Op", in.GetAppID().GetOp()),
		attribute.String("AppID.Value", in.GetAppID().GetValue()),
		attribute.String("LangID.Op", in.GetLangID().GetOp()),
		attribute.String("LangID.Value", in.GetLangID().GetValue()),
		attribute.String("MessageID.Op", in.GetMessageID().GetOp()),
		attribute.String("MessageID.Value", in.GetMessageID().GetValue()),
	)
	return span
}

func TraceMany(span trace1.Span, infos []*npool.MessageReq) trace1.Span {
	for index, info := range infos {
		span = trace(span, info, index)
	}
	return span
}
