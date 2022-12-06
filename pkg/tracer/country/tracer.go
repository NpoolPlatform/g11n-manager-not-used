package country

import (
	"fmt"

	"go.opentelemetry.io/otel/attribute"
	trace1 "go.opentelemetry.io/otel/trace"

	npool "github.com/NpoolPlatform/message/npool/g11n/mgr/v1/country"
)

func trace(span trace1.Span, in *npool.CountryReq, index int) trace1.Span {
	span.SetAttributes(
		attribute.String(fmt.Sprintf("ID.%v", index), in.GetID()),
		attribute.String(fmt.Sprintf("Country.%v", index), in.GetCountry()),
		attribute.String(fmt.Sprintf("Flag.%v", index), in.GetFlag()),
		attribute.String(fmt.Sprintf("Code.%v", index), in.GetCode()),
		attribute.String(fmt.Sprintf("Short.%v", index), in.GetShort()),
	)
	return span
}

func Trace(span trace1.Span, in *npool.CountryReq) trace1.Span {
	return trace(span, in, 0)
}

func TraceConds(span trace1.Span, in *npool.Conds) trace1.Span {
	span.SetAttributes(
		attribute.String("ID.Op", in.GetID().GetOp()),
		attribute.String("ID.Value", in.GetID().GetValue()),
		attribute.String("Country.Op", in.GetCountry().GetOp()),
		attribute.String("Country.Value", in.GetCountry().GetValue()),
		attribute.String("Code.Op", in.GetCode().GetOp()),
		attribute.String("Code.Value", in.GetCode().GetValue()),
	)
	return span
}

func TraceMany(span trace1.Span, infos []*npool.CountryReq) trace1.Span {
	for index, info := range infos {
		span = trace(span, info, index)
	}
	return span
}
