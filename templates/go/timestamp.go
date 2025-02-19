package golang

const timestampTpl = `{{ $f := .Field }}{{ $r := .Rules }}
	{{ template "required" . }}

	{{ if or $r.Lt $r.Lte $r.Gt $r.Gte $r.LtNow $r.GtNow $r.Within $r.Const }}
		if t := {{ accessor . }}; m.maskHas(mask, "{{ $f.Name }}") && t != nil {
			ts, err := ptypes.Timestamp(t)
			if err != nil { return {{ errCause . "err" (t "timestamp.valid" "value is not a valid timestamp") }} }

			{{ template "timestampcmp" . }}
		}
	{{ end }}
`
