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

const showDialog = ref(false)

</script>

<template>
  <v-app-bar flat border="b">
    <v-app-bar-title text="Applications" class="text-h5" />
    <template #append>
        <v-btn
          variant="flat"
          color="primary"
          class="mr-3"
          append-icon="mdi-plus-circle"
          @click="showDialog = true"
          :disabled="!account"
        >
          Add
        </v-btn>
      </template>
  </v-app-bar>
  <v-dialog v-model="showDialog" max-width="800px" overflowed>
    <template #default="{ isActive }">
      <v-card>
        <v-card-title class="text-h5">Application Setup</v-card-title>
        <v-divider></v-divider>
        <v-card-text class="ma-n4 ml-n7">
          <ApplicationSetup :initial-apps="applications!" :api-key="account!.apiKeys[0]!.key" @close="isActive.value = false"/>
        </v-card-text>
      </v-card>
    </template>
  </v-dialog>
  <v-row>
    <v-col>
      <template v-if="account?.applications?.length === 0">
        <v-card
          title="Welcome to OutageLab!"
          text="Let's get started by adding OutageLab to one of your applications. Within minutes you'll be running your first outage experiment. Click the button below to begin."
          max-width="700px"
          elevation="2"
        >
          <template #actions>
            <v-btn
              variant="flat"
              color="primary"
              append-icon="mdi-rocket-launch"
              size="small"
              class="ml-2 mb-1"
              @click="showDialog = true"
            >
              Let's go!
            </v-btn>
          </template>
        </v-card>
      </template>
      <v-sheet elevation="2" rounded v-else>
        <v-data-table
          :items="applications"
          :loading="!account"
          hide-default-footer
          hover
          id="applications"
          :headers="[
            { title: 'Name', key: 'id', width: '5%' },
            { title: 'Experiments', key: 'experiments', width: '25%' },
          ]"
          density="compact"
        >
          <template #headers>
            <v-data-table-headers
              color="primary"
              sticky
              :header-props="{
                class: 'bg-grey-lighten-3 font-weight-bold',
              }"
            />
          </template>
          <template #item="{ item }">
            <v-data-table-row :item="item" @click="router.push(`/applications/details/${item.id}`)">
              <template #item.id>
                <div>{{ item.id }}</div>
              </template>
              <template #item.experiments>
                <span v-if="appLiveEnvs[item.id]?.length">
                  <v-icon icon="mdi-circle" class="mr-1" color="green"></v-icon>
                  Active in {{ appLiveEnvs[item.id]?.length }} environment{{ appLiveEnvs[item.id]?.length !== 1 ? "s" : ""}}
                </span>
                <span v-else>
                  <v-icon icon="mdi-circle-outline" class="mr-1" color="grey"></v-icon>
                  <span>Inactive</span>
                </span>
              </template>
            </v-data-table-row>
          </template>
        </v-data-table>
      </v-sheet>
    </v-col>
  </v-row>
</template>
