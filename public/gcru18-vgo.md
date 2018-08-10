# Repeatable Builds with vgo


---

## About Me

- Microsoft Azure - Cloud Developer Advocate
- Using Go since 2010
- Twitter: @bketelsen
- Github: bketelsen
- Podcast: gotime.fm
- Email: bjk@microsoft.com

---

## This Talk

[https://cda.ms/jD](https://cda.ms/jD)

---

## History of Building Go

- 2009 6g, 6l - Makefiles |
- 2010 goinstall |
- 2011 go command |
- 2014 gopkg.in |
- 2014 semver proposal |
- 2015 go 1.5 vendor experiment |
- 2013-2018 godep,glide,gb,others |

---


## 2018

<br>

@fa[star gp-tip](vgo)

Documentation [https://research.swtch.com/vgo](https://research.swtch.com/vgo)
Code [https://github.com/golang/vgo](https://github.com/golang/vgo)
---

## vgo

<br/>
<i class='fas fa-check'></i> Defines Go Modules


---
## vgo

<br/>
<i class='fas fa-check'></i> Proposed Modification to "go" Tool


---

## vgo

- The Import Compatibility Rule |
- Minimal Version Selection |
- Semantic Import Versioning |
- Graceful API Updates |
- Reproducible Builds |
- Verifiable Builds |
- Verified Builds |
- Proxy Server Support |


---

## The Import Compatibility Rule 

> If an old package and a new package have the same import path, the new package must be backwards-compatible with the old package.

---

## Minimal Version Selection

> Minimal version selection allows modules to specify only minimum requirements for their dependency modules.

---

## Semantic Import Versioning



---?image=assets/image/impver.png&size=auto 60%&color=white


---

## Graceful API Updates

---

@title[Graceful API Updates]

<p><span class="slide-title">Graceful API Updates</span></p>

```go
package p // v1
import v2 "p/v2"

func EnableFoo() {
	//go:fix
	v2.SetFoo(true)
}
func DisableFoo() {
	//go:fix
	v2.SetFoo(false)
}
```

@[1,2](v1 package imports v2, allows you to set compatibility defaults.)
@[4-7](Future "go fix" commands will inline the changes.)

---

## Reproducible Builds

@fa[clipboard-list] Dependencies specified in go.mod

---

## Reproducible Builds

```text
module "github.com/bketelsen/captainhook"

require (
	"github.com/gorilla/context" v0.0.0-20160226214623-1ea25387ff6f
	"github.com/gorilla/mux" v1.6.0
)
```
@[1](Module Declaration)
@[3-6](Dependencies, with version specified)

---

## Verifiable Builds

@fa[clipboard-list] Inspect dependencies in compiled binary

---

## Verifiable Builds

```text
goversion -m captainhook 
captainhook go1.10
	path  github.com/bketelsen/captainhook
	mod   github.com/bketelsen/captainhook  (devel)
	dep   github.com/gorilla/mux            v1.6.0
```
@[1](rsc.io/goversion)
@[2](Binary name and Go version used to build it)
@[3](Import Path)
@[4](Module Declaration)
@[5](Dependencies With Version)

---

## Verified Builds

@fa[clipboard-list] Version hashes listed in go.modverify

---

## Verified Builds

```text
github.com/gorilla/mux v1.6.0 h1:UykbtMB/w5No2LmE16gINgLj+r/vbziTgaoERQv6U+0=
```

---

## Verified Builds

```bash
$ vgo verify
all modules verified
```

---
## Verified Builds

```bash
$ vgo verify
github.com/gorilla/mux v1.6.0: dir has been modified (...)
```

---

## Demo



---

## How?

### Change "go get" to make packages self-contained

---

## Change Build Sources

### $GOPATH/src/v/$PACKAGE/@v

---

## Demo



---

## Getting Started

- Add VCS tags to your releases
- Add go.mod file to your repository
- There is no step three

---


## Example with Makefile

[Captain Hook](https://github.com/bketelsen/captainhook)

### https://github.com/bketelsen/captainhook


---

## Proxy Server Support

```bash
$ GOPROXY=https://gomods.io vgo build
...
```

---

## Proxy Server 

- Keeps copies of modules for faster delivery |
- Set with environment variable |
  + GOPROXY=https://1.5.12.10/proxy vgo build
  + export GOPROXY=https://1.5.12.10/proxy
- Support for whitelists and/or blacklist |
- Fixed set of "official" proxies |
- But you can run your own! |

---



## Proxy

- Athens! |
- https://github.com/gomods/athens |
- Open Source |
- Created At Microsoft |
- In cooperation with the Go team |

---

## Registry

- Athens! |
- https://github.com/gomods/athens |
- Open Source |
- Created At Microsoft |
- In cooperation with the Go team |

---

## Status

<i class='fas fa-check'></i> Open Source Proxy Server
<i class='fas fa-check'></i> Multiple Storage Options
<i class='fas fa-check'></i> Well Defined Go Interface

---

## Coming Soon (tm)

<i class='fas fa-star'></i> Globally Distributed Proxies
<i class='fas fa-star'></i> Hosted Registry

---

## Hosted Registry

- Combines Proxy with Vanity Domains
- Automatically updates from VCS
  - Push tag, registry updates
- Adds extra metadata
  - scorecards
  - version history
  - import counts

---

## Hosted Registry

- Separate module from VCS
- Insulate package from Github control

---

## Hosted Registry 

- libgo.io
- gomods.io
- gopherpacks.io
- goreg.io
- goprox.io
- gomodul.es

---

## Hosted Registry

[gopherpacks.io](https://gopherpacks.io) *

*Coming Soon*
---

## Global Proxies

- Athens running on multiple clouds
- Multiple Hosting Companies
- In geographically varied regions
- Synchronized content
- Global reach, Fast Builds!!

---

## Foundation

- Long Term Plan for Athens |
- Multiple Companies Participating |
- Lets Athens live without SPOF |
- Microsoft |
- Google |
- Yandex ? |

---

## Open Source

- We <i class='fas fa-heart'></i> Contributions
- https://github.com/gomods/athens
- Lots of work left to do

---

### Questions?

<br>

<i class='fas fa-twitter'></i> @bketelsen

@fa[github gp-contact](bketelsen)

@fa[github gp-contact](gomods)

[GopherSnacks.com](https://www.gophersnacks.com)


---
<!-- .slide: data-background-image="/images/gitpitch-audience.jpg" -->

## Thank You
