import type { Directive } from 'vue'
import Prism from 'prismjs'
import 'prismjs/components/prism-go'
import 'prismjs/components/prism-bash'
import 'prismjs/components/prism-javascript'
import 'prismjs/components/prism-python'
import 'prismjs/themes/prism.css'

export default {
  async mounted(el, binding) {
    let code = el.innerText
    const leadingWhitespace = code.match(/^\s*/)[0]
    code = code.replace(new RegExp('^' + leadingWhitespace, 'gm'), '')
    code = code.replace(new RegExp('\\n\\s*$', 'g'), '')

    const language = binding.arg!
    const highlightedCode = Prism.highlight(code, Prism.languages[language], language)
    el.innerHTML = highlightedCode
  }
} as Directive
