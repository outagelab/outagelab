<script setup lang="ts">
import type { Application, Account, Environment, Rule } from '@/models/Account'
import AccountService from '@/services/AccountService'
import { onBeforeMount, ref, computed } from 'vue'
import { useRoute } from 'vue-router'

const accountService = new AccountService()


const route = useRoute()

let account = ref<Account>()
const application = computed((): Application | undefined => {
  return account.value?.applications?.find(x => x.id === route.params.id)
})
const environments = computed((): Environment[] | undefined => {
  return application.value?.environments
})
const rules = computed((): Rule[] | undefined => {
  return application.value?.rules
})

let confirmation = ref<string>()

onBeforeMount(refreshPage)

async function refreshPage() {
  account.value = await accountService.getAccount()
}

async function updatePage() {
  await accountService.updateAccount(account.value!)
  confirmation.value = "Changes saved"
  await refreshPage()
}

function addEnvironment() {
  environments.value?.push({ id: window.prompt('Environment Name') })
}

const headers = [
  { key: 'enabled', width: '5%', align: 'left' },
  { key: 'type', width: '15%' },
  { key: 'host', width: '25%' },
  { key: 'status', width: '15%' },
  { key: 'duration', width: '15%' },
  { key: 'delete', width: '5%' }
]
</script>

<template>
  <div class="d-flex align-center">
    <h1>{{ application?.id }}</h1>
    <v-btn @click="updatePage" class="ml-auto" color="primary">Save</v-btn>
  </div>

  <div class="d-flex align-center ga-3 mt-8">
    <h2>Environments</h2>
    <v-btn size="small" @click="addEnvironment()">Add</v-btn>
  </div>
  <v-row class="mt-1">
    <v-data-table
      :items="environments"
      hide-default-header
      hide-default-footer
      id="outagerules"
      :headers="[{ key: 'id', width: '5%' }, { key: 'enabled' }]"
      density="compact"
    >
      <template v-slot:item.id="props">
        <div>{{ props.item.id }}</div>
      </template>
      <template v-slot:item.enabled="props">
        <v-switch
          :label="props.item.enabled ? 'Live' : 'Disabled'"
          v-model="props.item.enabled"
          color="primary"
        />
      </template>
    </v-data-table>
  </v-row>

  <div class="d-flex align-center ga-3 mt-12">
    <h2>Outage Rules</h2>
    <v-btn size="small" @click="() => rules?.push({ type: 'send-http' })">Add</v-btn>
  </div>
  <v-row class="mt-1">
    <v-data-table
      :items="rules"
      hide-default-header
      hide-default-footer
      id="outagerules"
      :headers="headers"
    >
      <template v-slot:item.enabled="props">
        <v-switch v-model="props.item.enabled" color="primary" />
      </template>
      <template v-slot:item.type="props">
        {{
          {
            'send-http': 'Sending HTTP requests'
          }[props.item.type]
        }}
      </template>
      <template v-slot:item.host="props">
        <v-text-field label="Host" variant="underlined" v-model="props.item.host" />
      </template>
      <template v-slot:item.status="props">
        <v-text-field
          label="Response Status"
          variant="underlined"
          type="number"
          v-model.number="props.item.status"
        />
      </template>
      <template v-slot:item.duration="props">
        <v-text-field
          label="Response Delay"
          variant="underlined"
          type="number"
          v-model.number="props.item.duration"
        />
      </template>
      <template v-slot:item.delete="props">
        <v-btn size="small" @click="() => rules?.splice(props.index, 1)">Delete</v-btn>
      </template>
    </v-data-table>
  </v-row>
  <v-snackbar v-model="confirmation" :top="true" color="green">
      {{ confirmation }}
  </v-snackbar>
</template>

<style>
#outagerules td {
  border-bottom: none;
}

#outagerules .v-switch .v-input__details {
  display: none;
}

#outagerules .v-switch .v-selection-control {
  --v-input-control-height: 30px;
}
</style>
