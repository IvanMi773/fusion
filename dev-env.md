flake.nix specifies tools needed for development. to enter nix shell use:
```
nix develop
```

if flakes are not enabled in your system, use this command:
```
nix --extra-experimental-features 'nix-command flakes' develop
```

when inside a shell:
```
~/fusion/backend$ go run cmd/fusion/main.go
```

```
~/fusion/frontend$ pnpm run dev
```
