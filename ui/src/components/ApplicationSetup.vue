<script setup lang="ts">
import useInterval from "@/hooks/useInterval";
import type { Application } from "@/models/Account";
import AccountService from "@/services/AccountService";
import { ref, defineProps } from "vue"
import router from "@/router";
import CodeSampler from "./CodeSampler.vue";

const accountService = new AccountService()

const { apiKey, initialApps } = defineProps<{
    apiKey: any,
    initialApps: Application[]
}>()

const newApps = ref<Application[]>()
let tab = ref<string>("go")

useInterval(async () => {
    const account = await accountService.getAccount()

    const initialAppIds = initialApps.map(app => app.id)
    newApps.value = account?.applications?.filter(app => !initialAppIds.includes(app.id))
}, 5000)

const listeningDotsCount = ref<number>(0)
useInterval(() => {
    listeningDotsCount.value = (listeningDotsCount.value + 1) % 4
}, 333)
</script>

<template>
<v-stepper-vertical hide-actions multiple :model-value="[0, 1, 2, 3]" flat>
    <v-stepper-vertical-item value="1" class="pt-2">
        <template #title>
            <span class="text-h7">Install the package</span>
        </template>
        <CodeSampler
            v-model="tab"
            :tabs="[
                { label: 'Go', value: 'go', lang: 'bash' },
                { label: 'Node.js', value: 'nodejs', lang: 'bash' },
                { label: 'Python', value: 'python', lang: 'bash' }
            ]"
        >
            <template #go>
                go get github.com/outagelab/go-sdk
            </template>
            <template #nodejs>
                npm install -s outagelab
            </template>
            <template #python>
                pip install outagelab
            </template>
        </CodeSampler>
    </v-stepper-vertical-item>

    <v-stepper-vertical-item value="2">
        <template #title>
            <span class="text-h7">Add setup code</span>
        </template>
        <CodeSampler
            v-model="tab"
            :tabs="[
                { label: 'Go', value: 'go', lang: 'go' },
                { label: 'Node.js', value: 'nodejs', lang: 'js' },
                { label: 'Python', value: 'python', lang: 'py' }
            ]"
        >
            <template #go>
                import "github.com/outagelab/go-sdk/outagelab"

                func main() {
                    outagelab.Start(outagelab.Options{
                        Application: "my-service",
                        Environment: "local",
                        ApiKey:      "{{apiKey}}",
                        Host:        "https://app.outagelab.com",
                    })
                }
            </template>
            <template #nodejs>
                const outagelab = require("outagelab");

                outagelab.start({
                    application: "my-service",
                    environment: "local",
                    apiKey: "{{apiKey}}",
                    host: "https://app.outagelab.com",
                });
            </template>
            <template #python>
                import outagelab

                outagelab.start(
                    application="my-service",
                    environment="local",
                    api_key="{{apiKey}}",
                    host="https://app.outagelab.com"
                )
            </template>
        </CodeSampler>
    </v-stepper-vertical-item>
    <v-stepper-vertical-item value="3" class="pb-2">
        <template #title>
            <span class="text-h7">Run and verify</span>
        </template>
        <div>Once your application is running and connected successfully, it will appear here.</div>
        <v-card v-if="newApps?.length" class="mt-4" elevation="3" width="400px">
            <v-card-title class="text-h7">
                Applications Discovered
            </v-card-title>
            <v-divider></v-divider>
            <v-list>
                <v-list-item v-for="app in newApps" :key="app.id">
                    <span class="">{{ app.id }} {{ app.environments.map(x => x.id ).join(", ") }}</span>
                    <v-btn
                        size="small"
                        class="ml-4"
                        color="primary"
                        @click="router.push(`/applications/details/${app.id}`)"
                        append-icon="mdi-open-in-new"
                    >
                        Open
                    </v-btn>
                </v-list-item>
            </v-list>
        </v-card>
        <v-card v-else class="mt-4" elevation="3" width="300px">
            <template #text>
                <span class="text-body-1">Listening for new applications{{" .".repeat(listeningDotsCount)}}</span>
            </template>
        </v-card>
    </v-stepper-vertical-item>
</v-stepper-vertical>
</template>
