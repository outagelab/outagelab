<script setup lang="ts">
import { ref, computed, watch, onErrorCaptured } from 'vue'
import { RouterView } from 'vue-router'
import AccountService, { ApiAuthErrror } from './services/AccountService';
import logoImg from "@/assets/images/Logo.svg"


const apiService = new AccountService()

const token = ref<string>(localStorage.token)

watch(token, (newValue) => {
  localStorage.token = newValue
})

const tokenExpired = computed((): boolean => {
  const { exp } = JSON.parse(atob(token.value?.split('.')[1] || '') || '{}') as { exp: number }
  return !exp || new Date(exp * 1000) < new Date()
})

const callback = async (response: any) => {
    token.value = await apiService.login(response.credential)
}

onErrorCaptured((error) => {
  if (error.constructor === ApiAuthErrror) {
    token.value = ""
    return false // stop error propagation
  }
})
</script>

<template>
  <v-app>
    <template v-if="tokenExpired">
      <v-main>
        <v-overlay :model-value="true" :persistent="true" class="fill-height d-flex justify-center align-center mb-10">
          <v-container>
            <v-sheet elevation="6" rounded="xl" min-width="750px">
              <v-banner elevation="1" class="rounded-t-xl pb-7">
                <v-row>
                  <v-col class="d-flex justify-center flex-column align-center">
                      <h1 class="text-h1 kanit-light mb-5">OutageLab</h1>
                      <logoImg width="170px"/>
                  </v-col>
                </v-row>
              </v-banner>
              <v-row class="mt-6">
                <v-col class="d-flex justify-center">
                  <h2 class="text-h4">Sign in</h2>
                </v-col>
              </v-row>
              <v-row class="pt-1 pb-4">
                <v-col class="d-flex justify-center">
                  <GoogleLogin :callback="callback" popup-type="TOKEN" />
                </v-col>
              </v-row>
            </v-sheet>
          </v-container>
        </v-overlay>
      </v-main>
    </template>
    <template v-else>
      <v-navigation-drawer :width="235" permanent class="bg-grey-lighten-3" elevation="2">
        <v-banner density="compact" class="d-flex justify-items py-2 bg-grey-lighten-3" height="65px" >
          <span class="mr-2 kanit-light" style="font-size: 28px">OutageLab</span>
          <logoImg width="50px"/>
        </v-banner>
        <v-list-item link to="/applications" title="Applications" prepend-icon="mdi-application-brackets"/>
        <v-list-item link to="/api-keys" title="API Keys" prepend-icon="mdi-key-chain-variant" />
      </v-navigation-drawer>
      <v-main>
        <v-container class="ml-0 pa-6">
          <RouterView />
        </v-container>
      </v-main>
    </template>
  </v-app>
</template>

<style scoped></style>
