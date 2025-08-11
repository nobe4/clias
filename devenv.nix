{
  pkgs,
  ...
}:

{
  enterShell = ''
    source ~/.bashrc
    export PATH=$(pwd)/dist:$PATH
  '';

  packages = with pkgs; [
    golangci-lint
    golangci-lint-langserver
    goreleaser
    gopls
  ];

  languages.go.enable = true;

  scripts.test.exec = ''
    go test -race -count 10 -cover -coverprofile=coverage ./...
    go tool cover -html=coverage -o coverage.html
    rm coverage
  '';

  scripts.run.exec = ''
    go run cmd/clias/main.go $@
  '';

  scripts.lint.exec = ''
    golangci-lint run --fix ./...
  '';

  scripts.build.exec = ''
    goreleaser build --auto-snapshot --clean --single-target --output ./dist/clias
  '';

  enterTest = ''
    test
    lint
  '';
}
