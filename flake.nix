{
  description = "Fusion RSS reader development environment";

  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";
  };

  outputs = { self, nixpkgs }:
    let
      systems = [ "x86_64-linux" "aarch64-linux" "x86_64-darwin" "aarch64-darwin" ];
      forEachSystem = f: nixpkgs.lib.genAttrs systems (system: f system);
    in
    {
      devShells = forEachSystem (system:
        let
          pkgs = nixpkgs.legacyPackages.${system};
        in
        {
          default = pkgs.mkShell {
            packages = [
              pkgs.go_1_26
              pkgs.gotools  # provides goimports
              pkgs.nodejs_24
              pkgs.pnpm
              pkgs.sqlite
            ];

            FUSION_ALLOW_EMPTY_PASSWORD=true;
            FUSION_PORT=8081;

            shellHook = ''
              echo "Fusion dev environment"
              echo "  Go:   $(go version)"
              echo "  Node: $(node --version)"
              echo "  pnpm: $(pnpm --version)"
            '';
          };
        });
    };
}
