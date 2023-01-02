import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import VueGtag from "vue-gtag";
import './assets/main.css'

const app = createApp(App)

app.use(router)
app.use(VueGtag,  {
    config: { id: "G-K726D0Q3X5" }
}, router)

app.mount('#app')
