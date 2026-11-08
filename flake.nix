{
  description = "fang dev shell";

  inputs = {
    nixpkgs.url = "github:nixos/nixpkgs/nixos-unstable";
    flake-utils.url = "github:numtide/flake-utils";
    go-overlay.url = "github:purpleclay/go-overlay";
  };

  outputs = {
    nixpkgs,
    flake-utils,
    go-overlay,
    ...
  }:
    flake-utils.lib.eachDefaultSystem (system: let
      pkgs = import nixpkgs {
        inherit system;
        overlays = [go-overlay.overlays.default];
      };
    in {
      formatter = pkgs.alejandra;
      devShells.default = pkgs.mkShell {
        buildInputs = with pkgs; [
          ((go-bin.fromGoMod ./go.mod).withTools ["gopls" "staticcheck" "golangci-lint"])
          go-task
          typos
        ];
      };
    });
}
