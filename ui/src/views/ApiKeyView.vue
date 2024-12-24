<script setup lang="ts">
import type { Account, AccountApiKey } from '@/models/Account'
import AccountService from '@/services/AccountService'
import { onBeforeMount, ref, computed } from 'vue'

const accountService = new AccountService()

let account = ref<Account>()
let showConfirmation = ref<boolean>(false)

const apiKeys = computed((): AccountApiKey[] | undefined => {
  return account.value?.apiKeys
})

onBeforeMount(refreshPage)

async function refreshPage() {
  account.value = await accountService.getAccount()
}

function copyText(txt) {
    const clipboardItem = new ClipboardItem({
        'text/html': new Blob([txt], { type: 'text/html' }),
        'text/plain': new Blob([txt], { type: 'text/plain' })
    });

    navigator.clipboard.write([clipboardItem])

    showConfirmation.value = true
}

</script>

<template>
  <v-app-bar flat border="b">
      <v-app-bar-title text="API Keys" class="text-h5" />
  </v-app-bar>

  <v-row class="mt-1" v-for="{key} of apiKeys" :key="key">
      <v-card class="pa-3" elevation="2">
        <span class="text-body-1">
          {{ key.slice(0, 4) }}
          <span class="masked-text">{{ '•'.repeat(key.length - 4) }}</span>
          <v-btn class="ml-2" size="x-small" variant="tonal" icon="mdi-content-copy" @click="copyText(key)"></v-btn>
        </span>
      </v-card>
  </v-row>
  <v-snackbar v-model="showConfirmation" location="top right" :timeout="1500">
    API Key copied to clipboard
  </v-snackbar>
</template>

<style>
.masked-text {
  margin-left: -2px;
  letter-spacing: 0.1em;
}
</style>
