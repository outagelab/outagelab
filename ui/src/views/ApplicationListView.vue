<script setup lang="ts">
import type { Application, Account } from '@/models/Account'
import AccountService from '@/services/AccountService'
import { onBeforeMount, ref, computed } from 'vue'
import router from '@/router'
import ApplicationSetup from '@/components/ApplicationSetup.vue'

const accountService = new AccountService()

let account = ref<Account>()

const applications = computed((): Application[] | undefined => {
  return account.value?.applications
})

function addApplication() {
  let applicationName: string|null = ""
  while (applicationName === "") {
    applicationName = window.prompt("Application Name")
  }
  if (!applicationName) {
    return
  }
  applications.value?.push({ id: applicationName, environments: [], rules: [] })
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

<template v-if="account">
  <v-app-bar elevation="2">
    <v-app-bar-title text="Applications" class="text-h5" />
    <template #append>
        <v-btn
          variant="flat"
          color="primary"
          @click="addApplication"
          class="mr-3"
        >
          Add
        </v-btn>
      </template>
  </v-app-bar>
  <v-row>
    <v-col>
      <template v-if="!applications?.length">
        <ApplicationSetup />
      </template>
      <v-sheet v-else>
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
            <v-btn size="small" @click="() => router.push(`/applications/${props.item.id}`)">Edit</v-btn>
          </template>
        </v-data-table>
      </v-sheet>
    </v-col>
  </v-row>
</template>
