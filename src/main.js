import Home from './pages/Home.vue'
import NotFound from './pages/NotFound.vue'
import './index.css'

import { createApp, h } from 'vue'

const routes = {
    '/': Home,
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
        return h(this.CurrentComponent)
    }
}

createApp(SimpleRouter).mount('#app')
