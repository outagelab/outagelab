<script setup lang="ts">
import type { Account, AccountApiKey } from '@/models/Account'
import AccountService from '@/services/AccountService'
import { onBeforeMount, ref, computed } from 'vue'

const accountService = new AccountService()

let account = ref<Account>()

const apiKeys = computed((): AccountApiKey[] | undefined => {
  return account.value?.apiKeys
})

onBeforeMount(refreshPage)

async function refreshPage() {
  account.value = await accountService.getAccount()
}
</script>

<template>
  <v-app-bar elevation="2">
      <v-app-bar-title text="API Keys" class="text-h5" />
  </v-app-bar>

  <v-row class="mt-1">
    <v-data-table
      :items="apiKeys"
      hide-default-header
      hide-default-footer
      id="apikeys"
      :headers="[{ key: 'key' }]"
      density="compact"
    >
      <template v-slot:item.key="props">
        <div>{{ props.item.key }}</div>
      </template>
    </v-data-table>
  </v-row>
</template>

<style>
#apikeys td {
  border-bottom: none;
}
</style>
