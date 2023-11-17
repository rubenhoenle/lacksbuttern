{
  description = "Lachsbuttern wie ganz normale Lackaffen";

  inputs = {
    nixpkgs.url = "github:nixos/nixpkgs/nixos-unstable";
    treefmt-nix = {
      url = "github:numtide/treefmt-nix";
      inputs.nixpkgs.follows = "nixpkgs";
    };
  };

  outputs = { self, nixpkgs, treefmt-nix }:
    let
      system = "x86_64-linux";
      pkgs = import nixpkgs { inherit system; };
      treefmtEval = treefmt-nix.lib.evalModule pkgs ./treefmt.nix;
      lacksbuttern = pkgs.buildGoModule {
        pname = "lacksbuttern";
        version = "1.0.0";
        vendorHash = "sha256-8wYERVt3PIsKkarkwPu8Zy/Sdx43P6g2lz2xRfvTZ2E=";
        src = ./.;
      };
    in
    {

      formatter.${system} = treefmtEval.config.build.wrapper;
      checks.${system}.formatter = treefmtEval.config.build.check self;
      devShells.${system}.default = pkgs.mkShell {
        packages = with pkgs;[
          go
        ];
      };
      packages.${system} = {
        default = lacksbuttern;
      };
    };
}
