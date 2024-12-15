<script setup lang="ts">
    import { ref } from "vue"

    let apiKey = "325345234523452345"

    let tab = ref<string|null>("go")
</script>

<template>

<v-stepper-vertical hide-actions multiple :model-value="[0, 1, 2, 3]">
    <v-stepper-vertical-item value="1" title="Install the package">
        <v-card>
            <v-tabs v-model="tab" bg-color="primary" density="compact">
                <v-tab model-value="go">Go</v-tab>
                <v-tab model-value="nodejs">JavaScript</v-tab>
                <v-tab model-value="python">Python</v-tab>
            </v-tabs>
            <v-card-text>
                <v-tabs-window v-model="tab">
                    <v-tabs-window-item
                        value="go"
                        reverse-transition="false"
                        transition="false"
                    >
                        <pre v-highlight:sh>
                            go get github.com/outagelab/go-sdk
                        </pre>
                    </v-tabs-window-item>
                    <v-tabs-window-item
                        value="nodejs"
                        reverse-transition="false"
                        transition="false"
                    >
                        <pre v-highlight:sh>
                            npm install -s outagelab
                        </pre>
                    </v-tabs-window-item>
                    <v-tabs-window-item
                        value="python"
                        reverse-transition="false"
                        transition="false"
                    >
                        <pre v-highlight:sh>
                            pip install outagelab
                        </pre>
                    </v-tabs-window-item>
                </v-tabs-window>
            </v-card-text>
        </v-card>
    </v-stepper-vertical-item>

    <v-stepper-vertical-item value="2" title="Add code">
        <v-card>
            <v-tabs v-model="tab" bg-color="primary" density="compact">
                <v-tab model-value="go">Go</v-tab>
                <v-tab model-value="nodejs">JavaScript</v-tab>
                <v-tab model-value="python">Python</v-tab>
            </v-tabs>
            <v-card-text>
                <v-tabs-window v-model="tab">
                    <v-tabs-window-item
                        value="go"
                        reverse-transition="false"
                        transition="false"
                    >
                        <pre v-highlight:go>
                            import "github.com/outagelab/go-sdk/outagelab"

                            func main() {
                                outagelab.Start(outagelab.Options{
                                    Application: "my-service",
                                    Environment: "local",
                                    ApiKey:      "{{apiKey}}",
                                    Host:        "https://app.outagelab.com",
                                })
                            }
                        </pre>
                    </v-tabs-window-item>
                    <v-tabs-window-item
                        value="nodejs"
                        reverse-transition="false"
                        transition="false"
                    >
                        <pre v-highlight:js>
                            const outagelab = require("outagelab");

                            outagelab.start({
                                application: "my-service",
                                environment: "local",
                                apiKey: process.env.OUTAGELAB_API_KEY,
                                host: "https://app.outagelab.com",
                            });
                        </pre>
                    </v-tabs-window-item>
                    <v-tabs-window-item
                        value="python"
                        reverse-transition="false"
                        transition="false"
                    >
                        <pre v-highlight:py>
                            import outagelab

                            outagelab.start(
                                application="my-service",
                                environment="local",
                                api_key=os.environ.get("OUTAGELAB_API_KEY"),
                                host="https://app.outagelab.com"
                            )
                        </pre>
                    </v-tabs-window-item>
                </v-tabs-window>
            </v-card-text>
        </v-card>
    </v-stepper-vertical-item>
    <v-stepper-vertical-item value="3" title="Run and verify">
        TODO
    </v-stepper-vertical-item>
</v-stepper-vertical>
</template>
