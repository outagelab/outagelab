<script setup lang="ts">
import { ref, computed, watch } from 'vue'
import { RouterView } from 'vue-router'
import AccountService from './services/AccountService';

const apiService = new AccountService()

const token = ref<string>(localStorage.token)

watch(token, (newValue) => {
  localStorage.token = newValue
})

const tokenExpired = computed((): boolean => {
  const { exp } = JSON.parse(atob(token.value?.split('.')[1] || '') || '{}') as { exp: number }
  return !exp || new Date(exp * 1000) < new Date()
})

const callback = async (response) => {
  token.value = await apiService.login(response.credential)
}
</script>

<template>
  <v-app>
    <template v-if="tokenExpired">
      <v-container class="fill-height d-flex justify-center mb-10" fluid>
        <div class="mb-10">
          <v-row class="mb-16">
            <v-col class="d-flex justify-center">
              <h1 class="text-h1">OutageLab</h1>
            </v-col>
          </v-row>
          <v-row>
            <v-col class="d-flex justify-center">
              <h2 class="text-h4">Sign in</h2>
            </v-col>
          </v-row>
          <v-row class="mb-10">
            <v-col class="d-flex justify-center">
              <GoogleLogin :callback="callback" popup-type="TOKEN" />
            </v-col>
          </v-row>
        </div>
      </v-container>
    </template>
    <template v-else>
      <v-navigation-drawer :width="289">
        <v-list-item>
          <h1 class="a-10 text-h4 ma-3">OutageLab</h1>
        </v-list-item>
        <v-divider></v-divider>
        <v-list-item link to="/applications">
          <h4>Applications</h4>
        </v-list-item>
        <v-list-item link to="/api-keys">
          <h4>API Keys</h4>
        </v-list-item>
      </v-navigation-drawer>
      <v-main>
        <v-container> <RouterView /></v-container>
      </v-main>
    </template>
  </v-app>
</template>

<style scoped></style>
