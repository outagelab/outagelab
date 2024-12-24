import { render, type Directive } from 'vue'
import Prism from 'prismjs'
import 'prismjs/components/prism-go'
import 'prismjs/components/prism-bash'
import 'prismjs/components/prism-javascript'
import 'prismjs/components/prism-python'
import 'prismjs/themes/prism.css'

export default {
  async updated(el, binding) {
    let code = el.innerText
    code = code.replace(/^\s*\n/, '')
    const leadingWhitespace = code.match(/^\s*/)[0]
    code = code.replace(new RegExp('^' + leadingWhitespace, 'gm'), '')
    code = code.replace(new RegExp('\\n\\s*$', 'g'), '')

    const language = binding.arg!
    const highlightedCode = Prism.highlight(code, Prism.languages[language], language)

    let renderEl = null
    if (el.hasAttribute('hidden')) {
      renderEl = el.nextElementSibling
    } else {
      el.setAttribute('hidden', 'hidden')
      renderEl = document.createElement('pre')
      el.insertAdjacentElement('afterend', renderEl)
    }

    renderEl.innerHTML = highlightedCode
  }
} as Directive
