# Hacker Slides

[twitter]: # (@you)
[event]: # (Your Event Name)
[eventurl]: # (https://www.event.com/)
[title]: # (Deck Title)
[image]: # (/images/demo.png)
[imagealt]: # (A Dog Grinning)
[date]: # (March 17, 2018)
[videourl]: # (https://www.event.com/myvideo/url/)

### Hack together simple slides

<!-- .slide: data-transition="zoom" -->

---

## The Basics

- Separate slides using '`---`' on a blank line
- For vertical slides use '`--`'  
- Write github flavored markdown
- Click 'Present' (top right) when you're ready to talk

---

## Quick tips

- There is also a speaker view, with notes - press '`s`'
- Press '`?`' with focus on the presentation for shortcuts
- <em>You can use html when necessary</em>
- Share the 'Present' URL with anyone you like!

Note:
- Anything after `Note:` will only appear here

---

## More markdown (fragments)

* static text
* fragment <!-- .element: class="fragment" -->
* fragment grow <!-- .element: class="fragment grow" -->
* fragment highlight-red <!-- .element: class="fragment highlight-red" -->
* press key down <!-- .element: class="fragment fade-up" -->

--

## More markdown (tables)

****

|h1|h2|h3|
|-|-|-|
|a|b|c|

****

--

## More markdown (code)

```
version: '2'
services:
  slides:
    image: msoedov/hacker-slides

    ports:
      - 8080:8080
    volumes:
      - ./slides:/app/slides
    restart: always

    environment:
     - USER=bob
     - PASSWORD=pa55

```

--

## Local images

![demoPicture](/images/demo.png)

Copy images into slides/images/ & include with MD:

```
![demoPicture](/images/demo.png)

```
or HTML:

```
<img src="/images/demo.png">

```


---

## Learn more

- [RevealJS Demo/Manual](http://lab.hakim.se/reveal-js)
- [RevealJS Project/README](https://github.com/hakimel/reveal.js)
- [GitHub Flavored Markdown](https://help.github.com/articles/github-flavored-markdown)
