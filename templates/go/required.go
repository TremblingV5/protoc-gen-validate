package golang

const requiredTpl = `
	{{ if .Rules.GetRequired }}
		if m.maskHas(mask, "{{ .Field.Name }}") && {{ accessor . }} == nil {
			return {{ err . (t "<prefix>.required" "value is required") }}
		}
	{{ end }}
`
