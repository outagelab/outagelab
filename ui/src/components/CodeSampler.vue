<script setup lang="ts">
import { defineProps, defineModel } from 'vue'
import Prism from 'prismjs'

let { tabs } = defineProps<{
    tabs: { label: string, value: string, lang: string }[]
}>()

let model = defineModel<string>()

const dst = {} as { [x: string]: HTMLElement }
function onSrcUpdate(srcEl: HTMLElement, id: string, lang: string) {
    const dstEl = dst[id]
    if (!srcEl || !dstEl) {
        return
    }

    let code = srcEl.innerText
    code = code.replace(/^\s*\n/, '')
    const leadingWhitespace = code.match(/^\s*/)?.[0] || ''
    code = code.replace(new RegExp('^' + leadingWhitespace, 'gm'), '')
    code = code.replace(new RegExp('\\n\\s*$', 'g'), '')

    const highlightedCode = Prism.highlight(code, Prism.languages[lang], lang)

    dstEl.innerHTML = highlightedCode
}

function copyHTML() {
    const html = dst[model.value!]!.innerText;
    const clipboardItem = new ClipboardItem({
        'text/html': new Blob([html], { type: 'text/html' }),
        'text/plain': new Blob([html], { type: 'text/plain' })
    });

    navigator.clipboard.write([clipboardItem])
}
</script>

<template>

<v-card>
    <v-tabs v-model="model" bg-color="primary" density="compact">
        <v-tab v-for="tab of tabs" :key="tab.value" :value="tab.value">{{ tab.label }}</v-tab>
    </v-tabs>
    <v-card-text>
        <v-fab icon="mdi-content-copy" size="x-small" absolute location="top end" @click="copyHTML"></v-fab>
        <v-tabs-window v-model="model">
            <v-tabs-window-item
                v-for="tab of tabs"
                :key="tab.value"
                :value="tab.value"
                reverse-transition="false"
                transition="false"
                ref="windows"
            >
                <pre :ref="(el) => dst[tab.value] = el!"></pre>
                <pre hidden :ref="(el) => onSrcUpdate(el, tab.value, tab.lang)">
                    <slot :name="tab.value"></slot>
                </pre>
            </v-tabs-window-item>
        </v-tabs-window>
    </v-card-text>
</v-card>

</template>

<style>
.v-fab > .v-fab__container {
    transform: translateY(-18%) translateX(18%);
}
</style>
