import Home from '../pages/Home.vue'
import NotFound from '../pages/NotFound.vue'
import EveSSOCallback from '../pages/EveSSOCallback.vue'
import '../index.css'

import {Component, createApp, h, RendererElement, RendererNode, VNode} from 'vue'

import { createPinia } from 'pinia'
import mitt from "mitt";

const routes = new Map<string, Component>();
routes.set('/', Home);
routes.set('/eve-sso/callback', EveSSOCallback);

const SimpleRouter : Component = {
    data: () => ({
        currentRoute: window.location.pathname
    }),

    computed: {
        CurrentComponent() : Component {
            return routes.get(this.currentRoute) || NotFound
        }
    },

    render() : VNode<RendererNode, RendererElement, {}> {
        let queryParams = window.location.search
            .substring(1)
            .split('&')
            .map(q => q.split('='))
            .map(q => ({ [q[0]]: q[1] }))
        return h(this.CurrentComponent, { queryParams })
    }
}

const emitter = mitt()
const app = createApp(SimpleRouter)

app.use(createPinia())

app.provide('emitter', emitter)

app.mount('#app')
