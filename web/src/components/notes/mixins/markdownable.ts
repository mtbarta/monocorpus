import Vue from 'vue'
import marked from 'marked';
import sanitize from 'sanitize-html'
import katex from 'katex'

const renderer = new marked.Renderer();

/**
 * I'm taking a note from how Gitlab has done rendering.
 * 
 * https://github.com/gitlabhq/gitlabhq/blob/master/app/assets/javascripts/notebook/cells/markdown.vue
 * 
 * to render latex inside of markdown, they regex for latex and
 * then render both accordingly.
 */

const katexRegexString = `(
  ^\\\\begin{[a-zA-Z]+}\\s
  |
  ^\\$\\$
  |
  \\s\\$(?!\\$)
)
  ((.|\\n)+?)
(
  \\s\\\\end{[a-zA-Z]+}$
  |
  \\$\\$$
  |
  \\$
)
`.replace(/\s/g, '').trim();

renderer.paragraph = (t) => {
  let text = t;
  let inline = false;
  if (typeof katex !== 'undefined') {
    const katexString = text.replace(/&amp;/g, '&')
      .replace(/&=&/g, '\\space=\\space')
      .replace(/<(\/?)em>/g, '_');
    const regex = new RegExp(katexRegexString, 'gi');
    const matchLocation = katexString.search(regex);
    const numberOfMatches = katexString.match(regex);
    if (numberOfMatches && numberOfMatches.length !== 0) {
      if (matchLocation > 0) {
        let matches = regex.exec(katexString);
        inline = true;
        while (matches !== null) {
          const renderedKatex = katex.renderToString(matches[0].replace(/\$/g, ''));
          text = `${text.replace(matches[0], ` ${renderedKatex}`)}`;
          matches = regex.exec(katexString);
        }
      } else {
        const matches = regex.exec(katexString);
        text = katex.renderToString(matches[2]);
      }
    }
  }

  return `<p class="${inline ? 'inline-katex' : ''}">${text}</p>`;
};

renderer.listitem = function(text) {
  if (/^\s*\[[x ]\]\s*/.test(text)) {
  text = text
    .replace(/^\s*\[ \]\s*/, '<i class="icon material-icons">check_box_outline_blank</i> ')
    .replace(/^\s*\[x\]\s*/, '<i class="icon material-icons">check_box</i> ');
      return '<li style="list-style: none">' + text + '</li>';
    } else {
      return '<li>' + text + '</li>';
    }
  };

marked.setOptions({
  gfm: true,
  smartLists: true,
  breaks: true,
  sanitize: true,
  renderer,
  highlight: function(code) {
    return require('highlight.js').highlightAuto(code).value;
  }
})

export default Vue.extend({
  methods: {
    renderCode(code: string): string {
      const out = sanitize(marked(code.replace(/\\/g, '\\\\')), {
        allowedTags: false,
        allowedAttributes: {
          '*': ['class', 'style']
        }
      })
      return out
    }
  }
})