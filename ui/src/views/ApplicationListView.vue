<script setup lang="ts">
import type { Application, Account } from '@/models/Account'
import AccountService from '@/services/AccountService'
import { onBeforeMount, ref, computed } from 'vue'
import router from '@/router'

const accountService = new AccountService()

let account = ref<Account>()

const applications = computed((): Application[] | undefined => {
  return account.value?.applications
})

function addApplication() {
  applications.value?.push({ id: window.prompt("Application Name"), environments: [], rules: [] })
  accountService.updateAccount(account.value!)
}

onBeforeMount(refreshPage)

async function refreshPage() {
  account.value = await accountService.getAccount()
}

const appLiveEnvs = computed((): { [x: string]: string[] } => {
  const result = {} as { [x: string]: string[] }
  for (let app of account.value?.applications || []) {
    result[app.id] = app.environments.filter((x) => x.enabled).map((x) => x.id)
  }
  return result
})


</script>

<template>
  <div class="d-flex align-center ga-3">
    <h1>Applications</h1>
    <v-btn size="small" @click="addApplication">Add</v-btn>
  </div>
  <v-row class="mt-1">
    <v-data-table
      :items="applications"
      hide-default-header
      hide-default-footer
      id="applications"
      :headers="[
        { key: 'id', width: '5%' },
        { key: 'environments', width: '25%' },
        { key: 'edit', width: '25%' }
      ]"
      density="compact"
    >
      <template v-slot:item.id="props">
        <div>{{ props.item.id }}</div>
      </template>

      <template v-slot:item.environments="props">
        <template v-if="appLiveEnvs[props.item.id]?.length">
          Live in {{ appLiveEnvs[props.item.id]?.length }} environment<template
            v-if="appLiveEnvs[props.item.id]?.length != 1"
            >s</template
          >
        </template>
      </template>

      <template v-slot:item.edit="props">
        <v-btn size="small" @click="() => router.push(`/applications/${props.item.id}`)">Edit</v-btn>      </template>
    </v-data-table>
  </v-row>
</template>

<style scoped></style>
