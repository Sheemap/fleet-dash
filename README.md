# FleetDash

FleetDash is a realtime combat dashboard for EVE Online. FleetDash can parse the entire fleet's combat logs, and visualize them with customizable widgets.

---

## fleet-dash-web

This is the actual dashboard, the visualization hub. It opens a websocket to fleet-dash-core, and listens for relevant events. Itll injest em, and update any dashboard widgets that care. Due to the messaging system, it's very easy to build new widgets.

### Technical info

fleet-dash-web is built with Vue3 and Vite. Typescript is used for all component logic, and Tailwind for styling.

-- add more :)

---

### TODO

* Get a favicon that isnt the vue3 logo.

* Change how widgets are customized. Open a modal or something to show a much larger window, thats easier to digest.

* Add home screen text.