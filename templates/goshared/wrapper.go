package goshared

const wrapperTpl = `
	{{ $f := .Field }}{{ $r := .Rules }}

	if wrapper := {{ accessor . }}; wrapper != nil {
		{{ render (unwrap . "wrapper") }}
	} {{ if .MessageRules.GetRequired }} else if m.maskHas(mask, "{{ $f.Name }}") {
		return {{ err . (t "wrapper.required" "value is required and must not be nil.") }}
	} {{ end }}
`
