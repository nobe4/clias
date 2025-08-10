package generators

import "github.com/nobe4/dent.go"

//nolint:gochecknoglobals // This is fine.
var tmpls = map[string]string{
	"alias": dent.DedentString(`
		{{- range $b, $a := . -}}
			alias {{ $b }}='clias {{ $b }}'
		{{ end -}}
	`),

	"comp-bash": dent.DedentString(`
		TODO
	`),

	"comp-zsh": dent.DedentString(`
		TODO
	`),

	"comp-clias-bash": dent.DedentString(`
	_complete_clias() {
		local out=()

		_select_choices() {
			local choice="${1}"; shift
			local choices=("${@}")

			for elem in "${choices[@]}"; do
				if [[ $elem == "${choice}"* ]]; then
					out+=("$elem")
				fi
			done
		}

		local idx="$COMP_CWORD"
		local current_word=${COMP_WORDS[idx]}

		local binary_choices=({{- range $b, $a := . -}}'{{- $b }}' {{ end -}})
		{{- range $b, $a := . }}
		local {{ $b }}_choices=({{ range $n , $x := $a -}}'{{- $n }}' {{ end }})
		{{- end }}

		if [[ $idx == "1" ]]; then
			_select_choices "$current_word" "${binary_choices[@]}"
		elif [[ $idx == "2" ]]; then
			case "${COMP_WORDS[1]}" in
				{{- range $b, $a := . }}
				"{{ $b }}") _select_choices "$current_word" "${ {{- $b }}_choices[@]}" ;;
				{{- end }}
			esac
		fi


		COMPREPLY=("${out[@]}")
	}

	complete -F _complete_clias clias
	`),

	"comp-clias-zsh": dent.DedentString(`
		TODO
	`),
}
