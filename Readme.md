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


| Edit mode | Published  |
| --- | --- |
| ![1st](https://sc-cdn.scaleengine.net/i/520e2f4a8ca107b0263936507120027e.png) | ![1st](https://sc-cdn.scaleengine.net/i/7ae0d31a40b0b9e7acc3f131754874cf.png) |
|![2nd](https://sc-cdn.scaleengine.net/i/5acba66070e24f76bc7f20224adc611e.png) | ![2nd](https://sc-cdn.scaleengine.net/i/fee3e1374cb13b1d8c292becb7f514ae.png) |

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

Prepare for HTML Publishing
----
```shell
slides build
```
Then hook up the `/public` directory to something like Netlify and *profit*


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

Getting Help
------------

For **feature requests** and **bug reports**  submit an issue
to the GitHub issue tracker
