## Slides

[![MIT licensed](https://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/bketelsen/slides/blob/master/LICENSE)

This repo is a reworked version of [Sandstorm Hacker Slides](https://github.com/jacksingleton/hacker-slides) which features easy set up run outside of Sandstorm and without vagrant-spk. Likewise you can publish and edit your previous markdown slides which is not supported in the original version.


#### Features:

- Reach ui editor
- Markdown markup
- Live reload
- Color schemes
- Pdf print
- [Live version](https://talks.bjk.fyi)

Getting Started
----
Install from releases: [releases](https://github.com/bketelsen/slides/releases)
*coming soon* : brew, apt, and friends

Install from Source:

`go get github.com/bketelsen/slides`


Initialize a New Slide Repository
----
`slides init [reponame]`

```shell
slides init mytalks
cd mytalks
git init
git add --all
```
Then edit anything you want to change in `publish.tmpl` for individual slides and `root.tmpl` for the talk listing page.

`slides init` clones https://github.com/bketelsen/slides-template into the `mytalks` directory as a base for your decks.  The web assets in this directory are used to build the HTML files for your slides.

Directory Structure Of a Slide Repository
------
If you ran `slides init mytalks` your directory structure should look like this:

``` 
/mytalks --> repo root
    /public --> output files from `slides build`, published HTML
    /slides --> your slide decks, in Markdown format
    /static --> files used for `slides dev` local server
    /templates --> Go template files for `/public` and `/static`
    /initial-slides.md --> the template file used for `slides new {name}`
```


Create New Slide Deck
----
```shell
slides new mydeckname
```

Run Development Server (With Live Editing!)
----
```shell
slides dev
```

Visit [localhost](http://localhost:8080) to see your slides and make live edits.

Prepare for HTML Publishing
----
```shell
slides build
```

For maximum awesome, run `slides build` and setup Netlify or another static host to publish your `/public` directory.

Use local images
----
Store pictures you want to use in the images subfolder, slides/images/ and reference them in the editor as Markdown:
```
![demoPicture](/images/demo.png)
```
or as HTML:
```
<img src="/images/demo.png">
```

Screenshots
----

| Edit mode | Published  |
| --- | --- |
| ![1st](https://sc-cdn.scaleengine.net/i/520e2f4a8ca107b0263936507120027e.png) | ![1st](https://sc-cdn.scaleengine.net/i/7ae0d31a40b0b9e7acc3f131754874cf.png) |
|![2nd](https://sc-cdn.scaleengine.net/i/5acba66070e24f76bc7f20224adc611e.png) | ![2nd](https://sc-cdn.scaleengine.net/i/fee3e1374cb13b1d8c292becb7f514ae.png) |


Built on Open Source
----
This project is a heavily modified fork of [hacker-slides](https://github.com/msoedov/hacker-slides) and is built on the [Shoulders of Giants](/SHOULDERS.md)


Getting Help
------------

For **feature requests** and **bug reports**  submit an issue
to the GitHub issue tracker
