# Athens: The Center of Knowledge


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

[https://cda.ms/B6](https://cda.ms/B6)


---
## Athens
<!-- .slide: data-background-image="/images/athens.jpg" -->


---


## Athens Is

<br/>
@fa[check] Umbrella Project Name

---
## Athens Is

<br/>
@fa[check] Caching Proxy Server

---

## Athens Is

<br/>
@fa[check] That YOU Can Run

---

## Athens Is Also

<br/>
@fa[check] A Specification

---

## Athens Is Also

<br/>
@fa[check] About Trust

---

## Athens Is Also

<br/>
@fa[check] And Decentralized Verification

---

## Athens Is Also

<br/>
@fa[check] Just a project codename

---

## What's In It For You?

- Repeatable Builds 
- Every Time
- Even when Github/Gitlab/Bitbucket is down
- Local or closer caching proxies

---


## Specification

- API for the `go` command 
- Specification for module validation
- Specification for trust protocol


---

## Go Command

- Proxy support already included! 
- Set the GOPROXY environment variable 
- GOPROXY=https://some.address 

---

## Module Validation

- Hash of module contents stored in `go.sum` file
- Future-proof design to allow multiple hash algorithms
- Modules downloaded, validated and signed by `Notaries`

---

## Notaries

- Notaries download, verify hash, then sign modules 
- Notaries are completely independent 
- A `signed` module will contain a `certificate` 

---

## Publishers

- Publishers will collect certificates from Notaries 
- Certificates will be published at a `/log` API endpoint 
- Interested `Subscribers` may query the `/log` API 
- The `/log` API will support `after=` and `match=` for filtering 

---

## Notaries + Publishers

<br/>
<h3> Discoverability </h3>

---

## Authenticated Proxies

- Signed modules verified by public keys YOU choose to trust

---

## Authenticated Proxies

- The `go` command will support authenticated proxies 
- Trusted keys will be stored locally 
- `$GOPATH/go.key` and/or `$GOROOT/go.key` 
- Keys will be weighted, with a `1.0` score required for acceptance 

---

## Example go.key file

```
	server http://goproxy.microsoft.com/pkg
	server http://goproxy.google.com/pkg
	server http://goproxy.mycompany.com/pkg

	key 0.5 MyCompany s1:YXJaDOW...7IFlc=
	key 0.5 Google s1:pTFZ+webXa...f7SvSU=
	key 0.5 Microsoft s1:XEGD...eQQkkshI=
```
---

## Public Proxies

- Geographically distributed proxy servers
- Public servers, dependable infrastructure

---

## Run Your Own Proxy

- Run proxy locally, inside your firewall 
- Include/Exclude listing for modules 
- Prevent undesired modules from being used 
- Like `exclude github.com/gobuffalo/*` 
- Or `exclude github.com/*` 

---

## Demo!

---


## Decentralized, Federated, Independent

- Services are decentralized, independent
- Signed packages mean no worries about MITM
- All Open Source 
- Protocols are open 
- External services are interfaces, use your favorites

---

## Protocols Will Be Published

- Proposal to the `golang` repo soon 
- Open specification means anyone can participate 
- What will YOU build on top of the protocol? 

---

## Protocol Is Important

- Building block for future tools 

---

## Future Ideas - Code Provenance

- Verified, signed commits are proven
- Ability to require provenance at the proxy server

---

## Your Ideas On Top

- Code Quality 
- Code Metrics 
- Vanity Domain Server + Proxy

---

## Foundation

- Long Term Vision 
- Multiple Companies Participating 
- Members will host Notaries 
- Certificate Servers 
- Public Proxies 

---

## Foundation

- Allows Notaries to live without SPOF 
- Microsoft, Google already committed
- Others coming soon

---

## What's Next?

- Foundation being formed now 
- Notaries coming online soon 
- Public proxies available soon 
- Proposal for spec submitted soon

---

## Open Source

- We @fa[heart] Contributions
- https://github.com/gomods/athens
- Lots of work left to do
- Documentation almost completely missing
- 21 contributors, v0.0.1 Release last week

---

## Gratitude and Credit

- Russ Cox
- Aaron Schlesinger
- Paul Jolly
- Go Dependency projects (dep/govendor/glide, etc)
- Athens project contributors
- Buffalo / Mark Bates

---

### Reach Out!

<br>

@fa[twitter gp-contact](@bketelsen)

@fa[github gp-contact](bketelsen)

@fa[github gp-contact](gomods)



---?image=assets/image/gitpitch-audience.jpg&opacity=100

@title[Thank You!]

## Thank You
