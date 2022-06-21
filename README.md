## rng [![go report card](https://goreportcard.com/badge/github.com/jfcg/rng)](https://goreportcard.com/report/github.com/jfcg/rng) [![go.dev ref](https://pkg.go.dev/static/frontend/badge/badge.svg)](https://pkg.go.dev/github.com/jfcg/rng#pkg-overview)

Package `rng` is a compact, fast, [sponge](https://en.wikipedia.org/wiki/Sponge_function)-based,
lockless and hard-to-predict random number generator. See `Green tick > Go / Tests > Details` for
some statistical tests and benchmarks. It is compared with standard library's
[math/rand](https://pkg.go.dev/math/rand) and an alternative implementation
[exp/rand](https://pkg.go.dev/golang.org/x/exp/rand) below:

Library|Effective Entropy<br>(hidden information, in bits)
:---|---:
rng|128
std| 31
alt| 64

Library|Used memory (in bytes)
:---|---:
rng|  24
std|4920
alt|  48

`rng` API adheres to [semantic](https://semver.org) versioning. 
`rng` is not suitable for cryptographic applications.

See also [Contributing](./.github/CONTRIBUTING.md), [Security](./.github/SECURITY.md) and [Support](./.github/SUPPORT.md) guides.