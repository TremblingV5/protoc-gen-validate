package goshared

const constTpl = `{{ $r := .Rules }}
	{{ if $r.Const }}
		if m.maskHas(mask, "{{ .Field.Name }}") && {{ accessor . }} != {{ lit $r.GetConst }} {
			return {{ err . (t "<prefix>.const" "value must equal {{$1}}" $r.GetConst) }}
		}
	{{ end }}
`
