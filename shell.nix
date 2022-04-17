let
  pkgs = import (builtins.fetchGit {
    name = "nixos-unstable-2022-04-17";
    url = "https://github.com/nixos/nixpkgs";
    ref = "refs/heads/nixos-unstable";
    rev = "75ad56bdc927f3a9f9e05e3c3614c4c1fcd99fcb";
  }) { };
in pkgs.mkShell { buildInputs = [ pkgs.go pkgs.gotools pkgs.upx ]; }
