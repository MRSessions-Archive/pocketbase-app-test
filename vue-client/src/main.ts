/**
 * main.ts
 *
 * Bootstraps Vuetify and other plugins then mounts the App`
 */

// Components
import App from './App.vue'

// Composables
import { createApp } from 'vue'

// Plugins
import { registerPlugins } from '@/plugins'
import pbServer from './pocketbase'
import { pocketBaseSymbol } from './symbols/injectionSymbols'

const app = createApp(App)

registerPlugins(app)
app.provide(pocketBaseSymbol, pbServer)

app.mount('#app')
