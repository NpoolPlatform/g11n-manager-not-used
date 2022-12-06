package appcountry

import (
	"fmt"

	"go.opentelemetry.io/otel/attribute"
	trace1 "go.opentelemetry.io/otel/trace"

	npool "github.com/NpoolPlatform/message/npool/g11n/mgr/v1/appcountry"
)

func trace(span trace1.Span, in *npool.CountryReq, index int) trace1.Span {
	span.SetAttributes(
		attribute.String(fmt.Sprintf("ID.%v", index), in.GetID()),
		attribute.String(fmt.Sprintf("AppID.%v", index), in.GetAppID()),
		attribute.String(fmt.Sprintf("CountryID.%v", index), in.GetCountryID()),
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
		attribute.String("AppID.Op", in.GetAppID().GetOp()),
		attribute.String("AppID.Value", in.GetAppID().GetValue()),
		attribute.String("CountryID.Op", in.GetCountryID().GetOp()),
		attribute.String("CountryID.Value", in.GetCountryID().GetValue()),
	)
	return span
}

func TraceMany(span trace1.Span, infos []*npool.CountryReq) trace1.Span {
	for index, info := range infos {
		span = trace(span, info, index)
	}
	return span
}
