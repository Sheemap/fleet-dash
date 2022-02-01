import Home from '../pages/Home.vue'
import NotFound from '../pages/NotFound.vue'
import EveSSOCallback from '../pages/EveSSOCallback.vue'
import '../index.css'

import { createApp, h } from 'vue'
import axios from 'axios'
import VueAxios from 'vue-axios'

import { createPinia } from 'pinia'

const routes = {
    '/': Home,
    '/eve-sso/callback': EveSSOCallback,
}

const SimpleRouter = {
    data: () => ({
        currentRoute: window.location.pathname
    }),

    computed: {
        CurrentComponent() {
            return routes[this.currentRoute] || NotFound
        }
    },

    render() {
        let queryParams = window.location.search
            .substring(1)
            .split('&')
            .map(q => q.split('='))
            .map(q => ({ [q[0]]: q[1] }))
        return h(this.CurrentComponent, { queryParams })
    }
}

const app = createApp(SimpleRouter)

app.use(createPinia())
app.use(VueAxios, axios)

app.provide('axios', app.config.globalProperties.axios)

app.mount('#app')
