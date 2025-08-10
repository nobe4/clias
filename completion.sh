# file to test completion
# https://www.gnu.org/software/bash/manual/html_node/Programmable-Completion.html
# https://mill-build.org/blog/14-bash-zsh-completion.html
log(){
  echo "$@" >> /tmp/log
}

_complete_clias() {
  local out=()

  _select_choices() {
      local choice="${1}"; shift
      local choices=("${@}")

      log "select '${choice}' from '${choices[*]}'"

      for elem in "${choices[@]}"; do
        if [[ $elem == "${choice}"* ]]; then
            out+=("$elem")
        fi
      done
  }

  local idx="$COMP_CWORD"
  local current_word=${COMP_WORDS[idx]}

  local binary_choices=(ls go bash)
  local ls_choices=(a b c)
  local go_choices=(1 2 3)
  local bash_choices=(x y z)

  if [[ $idx == "1" ]]; then
      _select_choices "$current_word" "${binary_choices[@]}"
  elif [[ $idx == "2" ]]; then
      case "${COMP_WORDS[1]}" in
          "go") _select_choices "$current_word" "${go_choices[@]}" ;;
          "ls") _select_choices "$current_word" "${ls_choices[@]}" ;;
          "bash") _select_choices "$current_word" "${bash_choices[@]}" ;;
      esac
  fi


  COMPREPLY=("${out[@]}")
}

complete -F _complete_clias clias
