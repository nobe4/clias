{
  pkgs,
  ...
}:

{
  enterShell = ''
    source ~/.bashrc
    export PATH=$(find "$(pwd)" -type d -name 'clias_${pkgs.stdenv.hostPlatform.go.GOOS}_${pkgs.stdenv.hostPlatform.go.GOARCH}_*'):$PATH
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
    goreleaser --snapshot --clean
  '';

  enterTest = ''
    test
    lint
  '';
}
