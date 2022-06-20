# FleetDash

FleetDash is a realtime combat dashboard for EVE Online. FleetDash can parse the entire fleet's combat logs, and visualize them with customizable widgets.

---

## fleet-dash-web

This is the actual dashboard, the visualization hub. It opens a websocket to fleet-dash-core, and listens for relevant events. Itll injest em, and update any dashboard widgets that care. Due to the messaging system, it's very easy to build new widgets.

### Technical info

fleet-dash-web is built with Vue3 and Vite. Typescript is used for all component logic, and Tailwind for styling.

-- add more :)

This template should help get you started developing with Vue 3 in Vite. The template uses Vue 3 `<script setup>` SFCs, check out the [script setup docs](https://v3.vuejs.org/api/sfc-script-setup.html#sfc-script-setup) to learn more.

## Recommended IDE Setup

- [VSCode](https://code.visualstudio.com/) + [Volar](https://marketplace.visualstudio.com/items?itemName=johnsoncodehk.volar)
