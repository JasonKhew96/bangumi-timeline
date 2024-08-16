// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package anigamer

import (
	json "encoding/json"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjson923a13afDecodeScraperEntityAnigamer(in *jlexer.Lexer, out *Lists) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(Lists, 0, 0)
			} else {
				*out = Lists{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v1 List
			(v1).UnmarshalEasyJSON(in)
			*out = append(*out, v1)
			in.WantComma()
		}
		in.Delim(']')
	}
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson923a13afEncodeScraperEntityAnigamer(out *jwriter.Writer, in Lists) {
	if in == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
		out.RawString("null")
	} else {
		out.RawByte('[')
		for v2, v3 := range in {
			if v2 > 0 {
				out.RawByte(',')
			}
			(v3).MarshalEasyJSON(out)
		}
		out.RawByte(']')
	}
}

// MarshalJSON supports json.Marshaler interface
func (v Lists) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson923a13afEncodeScraperEntityAnigamer(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Lists) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson923a13afEncodeScraperEntityAnigamer(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Lists) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson923a13afDecodeScraperEntityAnigamer(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Lists) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson923a13afDecodeScraperEntityAnigamer(l, v)
}
func easyjson923a13afDecodeScraperEntityAnigamer1(in *jlexer.Lexer, out *List) {
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
		case "acg_sn":
			out.AcgSn = int(in.Int())
		case "anime_sn":
			out.AnimeSn = int(in.Int())
		case "title":
			out.Title = string(in.String())
		case "dc_c1":
			out.DcC1 = int(in.Int())
		case "dc_c2":
			out.DcC2 = int(in.Int())
		case "favorite":
			out.Favorite = bool(in.Bool())
		case "flag":
			out.Flag = int(in.Int())
		case "cover":
			out.Cover = string(in.String())
		case "info":
			out.Info = string(in.String())
		case "popular":
			out.Popular = string(in.String())
		case "highlightTag":
			(out.HighlightTag).UnmarshalEasyJSON(in)
		case "score":
			out.Score = float64(in.Float64())
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
func easyjson923a13afEncodeScraperEntityAnigamer1(out *jwriter.Writer, in List) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"acg_sn\":"
		out.RawString(prefix[1:])
		out.Int(int(in.AcgSn))
	}
	{
		const prefix string = ",\"anime_sn\":"
		out.RawString(prefix)
		out.Int(int(in.AnimeSn))
	}
	{
		const prefix string = ",\"title\":"
		out.RawString(prefix)
		out.String(string(in.Title))
	}
	{
		const prefix string = ",\"dc_c1\":"
		out.RawString(prefix)
		out.Int(int(in.DcC1))
	}
	{
		const prefix string = ",\"dc_c2\":"
		out.RawString(prefix)
		out.Int(int(in.DcC2))
	}
	{
		const prefix string = ",\"favorite\":"
		out.RawString(prefix)
		out.Bool(bool(in.Favorite))
	}
	{
		const prefix string = ",\"flag\":"
		out.RawString(prefix)
		out.Int(int(in.Flag))
	}
	{
		const prefix string = ",\"cover\":"
		out.RawString(prefix)
		out.String(string(in.Cover))
	}
	{
		const prefix string = ",\"info\":"
		out.RawString(prefix)
		out.String(string(in.Info))
	}
	{
		const prefix string = ",\"popular\":"
		out.RawString(prefix)
		out.String(string(in.Popular))
	}
	{
		const prefix string = ",\"highlightTag\":"
		out.RawString(prefix)
		(in.HighlightTag).MarshalEasyJSON(out)
	}
	{
		const prefix string = ",\"score\":"
		out.RawString(prefix)
		out.Float64(float64(in.Score))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v List) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson923a13afEncodeScraperEntityAnigamer1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v List) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson923a13afEncodeScraperEntityAnigamer1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *List) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson923a13afDecodeScraperEntityAnigamer1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *List) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson923a13afDecodeScraperEntityAnigamer1(l, v)
}
func easyjson923a13afDecodeScraperEntityAnigamer2(in *jlexer.Lexer, out *HighlightTag) {
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
		case "bilingual":
			out.Bilingual = bool(in.Bool())
		case "edition":
			out.Edition = string(in.String())
		case "vipTime":
			out.VipTime = string(in.String())
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
func easyjson923a13afEncodeScraperEntityAnigamer2(out *jwriter.Writer, in HighlightTag) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"bilingual\":"
		out.RawString(prefix[1:])
		out.Bool(bool(in.Bilingual))
	}
	{
		const prefix string = ",\"edition\":"
		out.RawString(prefix)
		out.String(string(in.Edition))
	}
	{
		const prefix string = ",\"vipTime\":"
		out.RawString(prefix)
		out.String(string(in.VipTime))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v HighlightTag) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson923a13afEncodeScraperEntityAnigamer2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v HighlightTag) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson923a13afEncodeScraperEntityAnigamer2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *HighlightTag) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson923a13afDecodeScraperEntityAnigamer2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *HighlightTag) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson923a13afDecodeScraperEntityAnigamer2(l, v)
}
