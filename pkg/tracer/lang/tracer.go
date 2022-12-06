package lang

import (
	"fmt"

	"go.opentelemetry.io/otel/attribute"
	trace1 "go.opentelemetry.io/otel/trace"

	npool "github.com/NpoolPlatform/message/npool/g11n/mgr/v1/lang"
)

func trace(span trace1.Span, in *npool.LangReq, index int) trace1.Span {
	span.SetAttributes(
		attribute.String(fmt.Sprintf("ID.%v", index), in.GetID()),
		attribute.String(fmt.Sprintf("Lang.%v", index), in.GetLang()),
		attribute.String(fmt.Sprintf("Logo.%v", index), in.GetLogo()),
		attribute.String(fmt.Sprintf("Name.%v", index), in.GetName()),
		attribute.String(fmt.Sprintf("Short.%v", index), in.GetShort()),
	)
	return span
}

func Trace(span trace1.Span, in *npool.LangReq) trace1.Span {
	return trace(span, in, 0)
}

func TraceConds(span trace1.Span, in *npool.Conds) trace1.Span {
	span.SetAttributes(
		attribute.String("ID.Op", in.GetID().GetOp()),
		attribute.String("ID.Value", in.GetID().GetValue()),
		attribute.String("Lang.Op", in.GetLang().GetOp()),
		attribute.String("Lang.Value", in.GetLang().GetValue()),
		attribute.String("Name.Op", in.GetName().GetOp()),
		attribute.String("Name.Value", in.GetName().GetValue()),
		attribute.String("Short.Op", in.GetShort().GetOp()),
		attribute.String("Short.Value", in.GetShort().GetValue()),
	)
	return span
}

func TraceMany(span trace1.Span, infos []*npool.LangReq) trace1.Span {
	for index, info := range infos {
		span = trace(span, info, index)
	}
	return span
}
