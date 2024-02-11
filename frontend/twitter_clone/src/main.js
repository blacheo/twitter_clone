import './assets/main.css'

import { createApp } from 'vue'
import App from './App.vue'
import axios from './plugins/axios'

const app = createApp(App)

app.use(axios, {
	baseUrl: 'http://localhost:5000/'
})

app.mount('#app')
