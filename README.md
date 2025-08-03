## rng [![go.dev ref](https://pkg.go.dev/static/frontend/badge/badge.svg)](https://pkg.go.dev/github.com/jfcg/rng#pkg-overview) [![report card](https://goreportcard.com/badge/github.com/jfcg/rng)](https://goreportcard.com/report/github.com/jfcg/rng) [![coverage](./.github/cover.svg)](https://github.com/jfcg/rng/actions/workflows/QA.yml) [![OpenSSF badge](https://www.bestpractices.dev/projects/8318/badge)](https://www.bestpractices.dev/projects/8318)

Package `rng` is a compact, fast, [sponge](https://en.wikipedia.org/wiki/Sponge_function)-based,
lockless and hard-to-predict random number generator. See `Green tick > QA / Tests > Details` for
some statistical tests and benchmarks. It is compared with standard library's
[math/rand](https://pkg.go.dev/math/rand) and [math/rand/v2](https://pkg.go.dev/math/rand/v2) below:

Library|Effective entropy in bits|Used memory in bytes
:---|---:|---:
rng      |128|  24
std      | 31|4920
v2.PCG   | 64|  32
v2.ChaCha|256| 336

`rng` API adheres to [semantic](https://semver.org) versioning. 
`rng` is not suitable for cryptographic applications because uses 128 bits capacity.

### Support
See [Contributing](./.github/CONTRIBUTING.md), [Security](./.github/SECURITY.md) and [Support](./.github/SUPPORT.md) guides. Also if you use `rng` and like it, please support via [Github Sponsors](https://github.com/sponsors/jfcg) or:
- BTC:`bc1qr8m7n0w3xes6ckmau02s47a23e84umujej822e`
- ETH:`0x3a844321042D8f7c5BB2f7AB17e20273CA6277f6`
