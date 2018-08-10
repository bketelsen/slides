function isPreview() {
    return !!window.location.search.match(/preview/gi);
}

function initializeReveal() {
    // Full list of configuration options available at:
    // https://github.com/hakimel/reveal.js#configuration

    Reveal.initialize({
        controls: true,
        progress: true,
        history: true,
        center: true,
        transition: 'slide', // none/fade/slide/convex/concave/zoom
        transitionSpeed: isPreview() ? 'fast' : 'default',
        embedded: isPreview() ? true : false,

        // Optional reveal.js plugins
        dependencies: [{
            src: '/static/reveal.js/lib/js/classList.js',
            condition: function () {
                return !document.body.classList;
            }
        },

        // Interpret Markdown in <section> elements
        {
            src: '/static/reveal.js/plugin/markdown/marked.js',
            condition: function () {
                return !!document.querySelector('[data-markdown]');
            }
        }, {
            src: '/static/reveal.js/plugin/markdown/markdown.js',
            condition: function () {
                return !!document.querySelector('[data-markdown]');
            }
        },

        // Syntax highlight for <code> elements
        {
            src: '/static/reveal.js/plugin/highlight/highlight.js',
            async: true,
            callback: function () {
                hljs.initHighlightingOnLoad();
            }
        },

        // Zoom in and out with Alt+click
        {
            src: '/static/reveal.js/plugin/zoom-js/zoom.js',
            async: true
        },

        // Speaker notes
        {
            src: '/static/reveal.js/plugin/notes/notes.js',
            async: true
        },

        // MathJax
        {
            src: '/static/reveal.js/plugin/math/math.js',
            async: true
        },

        {
            src: '/static/reveal.js/lib/js/classList.js',
            condition: function () {
                return !document.body.classList;
            }
        }
        ]
    });

    themesCtrl();
}

function highlightAnyCodeBlocks() {
    $(document).ready(function () {
        $('pre code').each(function (i, block) {
            hljs.highlightBlock(block);
        });
    });
}

function insertMarkdownReference() {
    var markdownReference = $('<section/>', {
        'data-markdown': "/slides.md",
        'data-separator': "^---\n",
        'data-separator-vertical': "^--\n",
        'data-separator-notes': "^Note:",
        'data-charset': "utf-8"
    });

    $('.slides').html(markdownReference);
}

function scrollToCurrentSlide() {
    var i = Reveal.getIndices();
    Reveal.slide(i.h, i.v, i.f);
}

function reloadMarkdown() {
    insertMarkdownReference();
    RevealMarkdown.initialize();
    highlightAnyCodeBlocks();
    scrollToCurrentSlide();
}

function externalLinksInNewWindow() {
    $(document.links).filter(function () {
        return this.hostname != window.location.hostname;
    }).attr('target', '_blank');
}

insertMarkdownReference();
initializeReveal();

function themesCtrl() {
    var defaultTheme = "black.css",
        currentTheme = localStorage.getItem('theme?') ||
            defaultTheme;

    function setTheme(theme) {
        cssEl = $("#theme");
        cssEl.attr("href", "/static/reveal.js/css/theme/" + theme);
        localStorage.setItem('theme?', theme);
    }
    setTheme(currentTheme);

    if (!isPreview()) {
        return
    }
    var availableThemes = [
        "black.css",
        "beige.css",
        "blood.css",
        "league.css",
        "moon.css",
        "night.css",
        "serif.css",
        "simple.css",
        "sky.css",
        "solarized.css",
        "white.css",
    ];
    themeEl = $("#themes");
    availableThemes.forEach(function (theme) {
        elem = $("<option value=" + theme + ">" + theme + "</option>");
        themeEl.append(elem);
    })
    themeEl.val(currentTheme);
    themeEl.change(function () {
        val = themeEl.val()
        setTheme(val);
    });
    themeEl.attr("hidden", false);
}

// Monkey patch Reveal so we can reload markdown through an
// inter window message (using the reveal rpc api)
// (yes, reveal has an rpc api!)
// see save.js
Reveal.reloadMarkdown = reloadMarkdown;
