<script setup lang="ts">
import type { Application, Account, Environment, Rule } from '@/models/Account'
import AccountService from '@/services/AccountService'
import { onBeforeMount, ref, computed } from 'vue'
import { useRoute } from 'vue-router'

const accountService = new AccountService()

const route = useRoute()

let initAccountJSON = ""
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

const dirty = computed((): boolean => {
  return initAccountJSON !== JSON.stringify(account.value)
})

let showConfirmation = ref<boolean>()
let showError = ref<string>()

onBeforeMount(refreshPage)

async function refreshPage() {
  account.value = await accountService.getAccount()
  initAccountJSON = JSON.stringify(account.value)
}

async function updatePage() {
  const {valid} = await form.value.validate()
  if (!valid) {
    showError.value = 'validation'
    return
  }
  try {
    for (let app of account.value?.applications || []) {
      for (let rule of app.rules || []) {
        if (!rule.status) {
          rule.status = null
        }
        if (!rule.duration) {
          rule.duration = null
        }
      }
    }
    await accountService.updateAccount(account.value!)
    showConfirmation.value = true
  } catch {
    showError.value = 'internal'
  }
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

const hostFieldRules = [
  (value: string): string|true => {
    if (!value) {
      return "Required"
    }
    if (value.includes("/") || value.includes("?") || value.includes("#")) {
      return "Exclude scheme, path, and other URL segments"
    }
    return true
  }
]

const responseStatusFieldRules = [
  (value: number): string|true => {
    if (!value) {
      return true
    }
    if (value < 400 || value >= 600) {
      console.log({hey: value})
      return "Must be between 400-599"
    }
    return true
  }
]

const responseDelayFieldRules = [
  (value: number): string|true => {
    if (!value) {
      return true
    }
    if (value < 0) {
      return "Must not be negative"
    }
    return true
  }
]

const form = ref<any>(null)
</script>

<template>
  <div class="d-flex align-center">
    <h1>{{ application?.id }}</h1>
    <v-btn @click="updatePage" class="ml-auto" color="primary" :disabled="!dirty">Update</v-btn>
  </div>
  <v-form ref="form">
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
      <v-btn size="small" @click="() => rules?.push({ type: 'send-http', enabled: true })">Add</v-btn>
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
              'send-http': 'HTTP Client Requests'
            }[props.item.type]
          }}
        </template>
        <template v-slot:item.host="props">
          <v-text-field
            label="Request Host"
            variant="underlined"
            v-model="props.item.host"
            required
            :rules="hostFieldRules"
          />
        </template>
        <template v-slot:item.status="props">
          <v-text-field
            label="Set Error Status"
            variant="underlined"
            type="number"
            v-model.number="props.item.status"
            :rules="responseStatusFieldRules"
          />
        </template>
        <template v-slot:item.duration="props">
          <v-text-field
            label="Add Latency (Seconds)"
            variant="underlined"
            type="number"
            v-model.number="props.item.duration"
            :rules="responseDelayFieldRules"
          />
        </template>
        <template v-slot:item.delete="props">
          <v-btn size="small" @click="() => rules?.splice(props.index, 1)">Delete</v-btn>
        </template>
      </v-data-table>
    </v-row>
  </v-form>
  <v-snackbar v-model="showConfirmation" :top="true" color="green">
    <h4>Changes saved</h4>
    Allow up to 30 seconds for changes to take effect
  </v-snackbar>
  <v-snackbar v-model="showError" :top="true" color="red">
    <template v-if="showError == 'validation'">
      <h4>Validation error</h4>
      Resolve validation errors and try again
    </template>
    <template v-if="showError == 'internal'">
      <h4>Something went wrong</h4>
      Try reloading the page
    </template>
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
