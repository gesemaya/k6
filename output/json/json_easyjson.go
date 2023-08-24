// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package json

import (
	json "encoding/json"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
	metrics "github.com/gesemaya/k6/metrics"
	time "time"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjson42239ddeDecodeGoK6IoK6OutputJson(in *jlexer.Lexer, out *sampleEnvelope) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "metric":
			out.Metric = string(in.String())
		case "type":
			out.Type = string(in.String())
		case "data":
			easyjson42239ddeDecode(in, &out.Data)
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson42239ddeEncodeGoK6IoK6OutputJson(out *jwriter.Writer, in sampleEnvelope) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"metric\":"
		out.RawString(prefix[1:])
		out.String(string(in.Metric))
	}
	{
		const prefix string = ",\"type\":"
		out.RawString(prefix)
		out.String(string(in.Type))
	}
	{
		const prefix string = ",\"data\":"
		out.RawString(prefix)
		easyjson42239ddeEncode(out, in.Data)
	}
	out.RawByte('}')
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v sampleEnvelope) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson42239ddeEncodeGoK6IoK6OutputJson(w, v)
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *sampleEnvelope) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson42239ddeDecodeGoK6IoK6OutputJson(l, v)
}
func easyjson42239ddeDecode(in *jlexer.Lexer, out *struct {
	Time     time.Time         `json:"time"`
	Value    float64           `json:"value"`
	Tags     *metrics.TagSet   `json:"tags"`
	Metadata map[string]string `json:"metadata,omitempty"`
}) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "time":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.Time).UnmarshalJSON(data))
			}
		case "value":
			out.Value = float64(in.Float64())
		case "tags":
			if in.IsNull() {
				in.Skip()
				out.Tags = nil
			} else {
				if out.Tags == nil {
					out.Tags = new(metrics.TagSet)
				}
				(*out.Tags).UnmarshalEasyJSON(in)
			}
		case "metadata":
			if in.IsNull() {
				in.Skip()
			} else {
				in.Delim('{')
				if !in.IsDelim('}') {
					out.Metadata = make(map[string]string)
				} else {
					out.Metadata = nil
				}
				for !in.IsDelim('}') {
					key := string(in.String())
					in.WantColon()
					var v1 string
					v1 = string(in.String())
					(out.Metadata)[key] = v1
					in.WantComma()
				}
				in.Delim('}')
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson42239ddeEncode(out *jwriter.Writer, in struct {
	Time     time.Time         `json:"time"`
	Value    float64           `json:"value"`
	Tags     *metrics.TagSet   `json:"tags"`
	Metadata map[string]string `json:"metadata,omitempty"`
}) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"time\":"
		out.RawString(prefix[1:])
		out.Raw((in.Time).MarshalJSON())
	}
	{
		const prefix string = ",\"value\":"
		out.RawString(prefix)
		out.Float64(float64(in.Value))
	}
	{
		const prefix string = ",\"tags\":"
		out.RawString(prefix)
		if in.Tags == nil {
			out.RawString("null")
		} else {
			(*in.Tags).MarshalEasyJSON(out)
		}
	}
	if len(in.Metadata) != 0 {
		const prefix string = ",\"metadata\":"
		out.RawString(prefix)
		{
			out.RawByte('{')
			v2First := true
			for v2Name, v2Value := range in.Metadata {
				if v2First {
					v2First = false
				} else {
					out.RawByte(',')
				}
				out.String(string(v2Name))
				out.RawByte(':')
				out.String(string(v2Value))
			}
			out.RawByte('}')
		}
	}
	out.RawByte('}')
}
func easyjson42239ddeDecodeGoK6IoK6OutputJson1(in *jlexer.Lexer, out *metricEnvelope) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "type":
			out.Type = string(in.String())
		case "data":
			easyjson42239ddeDecode1(in, &out.Data)
		case "metric":
			out.Metric = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson42239ddeEncodeGoK6IoK6OutputJson1(out *jwriter.Writer, in metricEnvelope) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"type\":"
		out.RawString(prefix[1:])
		out.String(string(in.Type))
	}
	{
		const prefix string = ",\"data\":"
		out.RawString(prefix)
		easyjson42239ddeEncode1(out, in.Data)
	}
	{
		const prefix string = ",\"metric\":"
		out.RawString(prefix)
		out.String(string(in.Metric))
	}
	out.RawByte('}')
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v metricEnvelope) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson42239ddeEncodeGoK6IoK6OutputJson1(w, v)
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *metricEnvelope) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson42239ddeDecodeGoK6IoK6OutputJson1(l, v)
}
func easyjson42239ddeDecode1(in *jlexer.Lexer, out *struct {
	Name       string               `json:"name"`
	Type       metrics.MetricType   `json:"type"`
	Contains   metrics.ValueType    `json:"contains"`
	Thresholds metrics.Thresholds   `json:"thresholds"`
	Submetrics []*metrics.Submetric `json:"submetrics"`
}) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "name":
			out.Name = string(in.String())
		case "type":
			if data := in.UnsafeBytes(); in.Ok() {
				in.AddError((out.Type).UnmarshalText(data))
			}
		case "contains":
			if data := in.UnsafeBytes(); in.Ok() {
				in.AddError((out.Contains).UnmarshalText(data))
			}
		case "thresholds":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.Thresholds).UnmarshalJSON(data))
			}
		case "submetrics":
			if in.IsNull() {
				in.Skip()
				out.Submetrics = nil
			} else {
				in.Delim('[')
				if out.Submetrics == nil {
					if !in.IsDelim(']') {
						out.Submetrics = make([]*metrics.Submetric, 0, 8)
					} else {
						out.Submetrics = []*metrics.Submetric{}
					}
				} else {
					out.Submetrics = (out.Submetrics)[:0]
				}
				for !in.IsDelim(']') {
					var v3 *metrics.Submetric
					if in.IsNull() {
						in.Skip()
						v3 = nil
					} else {
						if v3 == nil {
							v3 = new(metrics.Submetric)
						}
						easyjson42239ddeDecodeGoK6IoK6Metrics(in, v3)
					}
					out.Submetrics = append(out.Submetrics, v3)
					in.WantComma()
				}
				in.Delim(']')
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson42239ddeEncode1(out *jwriter.Writer, in struct {
	Name       string               `json:"name"`
	Type       metrics.MetricType   `json:"type"`
	Contains   metrics.ValueType    `json:"contains"`
	Thresholds metrics.Thresholds   `json:"thresholds"`
	Submetrics []*metrics.Submetric `json:"submetrics"`
}) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"name\":"
		out.RawString(prefix[1:])
		out.String(string(in.Name))
	}
	{
		const prefix string = ",\"type\":"
		out.RawString(prefix)
		out.Raw((in.Type).MarshalJSON())
	}
	{
		const prefix string = ",\"contains\":"
		out.RawString(prefix)
		out.Raw((in.Contains).MarshalJSON())
	}
	{
		const prefix string = ",\"thresholds\":"
		out.RawString(prefix)
		out.Raw((in.Thresholds).MarshalJSON())
	}
	{
		const prefix string = ",\"submetrics\":"
		out.RawString(prefix)
		if in.Submetrics == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v4, v5 := range in.Submetrics {
				if v4 > 0 {
					out.RawByte(',')
				}
				if v5 == nil {
					out.RawString("null")
				} else {
					easyjson42239ddeEncodeGoK6IoK6Metrics(out, *v5)
				}
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}
func easyjson42239ddeDecodeGoK6IoK6Metrics(in *jlexer.Lexer, out *metrics.Submetric) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "name":
			out.Name = string(in.String())
		case "suffix":
			out.Suffix = string(in.String())
		case "tags":
			if in.IsNull() {
				in.Skip()
				out.Tags = nil
			} else {
				if out.Tags == nil {
					out.Tags = new(metrics.TagSet)
				}
				(*out.Tags).UnmarshalEasyJSON(in)
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson42239ddeEncodeGoK6IoK6Metrics(out *jwriter.Writer, in metrics.Submetric) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"name\":"
		out.RawString(prefix[1:])
		out.String(string(in.Name))
	}
	{
		const prefix string = ",\"suffix\":"
		out.RawString(prefix)
		out.String(string(in.Suffix))
	}
	{
		const prefix string = ",\"tags\":"
		out.RawString(prefix)
		if in.Tags == nil {
			out.RawString("null")
		} else {
			(*in.Tags).MarshalEasyJSON(out)
		}
	}
	out.RawByte('}')
}
