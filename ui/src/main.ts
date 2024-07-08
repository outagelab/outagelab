import './assets/main.css'

import { createApp } from 'vue'
import App from './App.vue'
import router from './router'

// Vuetify
import 'vuetify/styles'
import { createVuetify } from 'vuetify'
import * as components from 'vuetify/components'
import * as directives from 'vuetify/directives'

import vue3GoogleLogin from 'vue3-google-login'

const app = createApp(App)

app.use(vue3GoogleLogin, {
  clientId: '587099227224-7ah3jhomkin3ok6b1f7s20tav9qu70n9.apps.googleusercontent.com'
})

// TODO: installs everything without tree shaking, see docs to optimize
app.use(
  createVuetify({
    components,
    directives,
    theme: {
      //   defaultTheme: 'dark'
    }
  })
)

app.use(router)

app.mount('#app')
