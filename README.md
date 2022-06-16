# FleetDash Web

FleetDash is a distributed log parsing tool for EVE Online that allows you to parse and visualize logs from your entire fleet.

See live counters of how much damage you've done, who is currently ECM'd, and more.

You can access the tool from https://fleetdash.space, or build and deploy it yourself.

## Overview

fleet-dash-web is the frontend for FleetDash. It is built with Vue3, and uses EVE ESI and fleet-dash-core as a backend.

Events are received via websockets, and emitted to all active components based on the event type.
Components can subscribe to only events they are interested in, and can do whatever they wish with these events.
This system allows for quick development of new components which only need to worry about their own business.


## Setup

fleet-dash-web is a single page app, and is quite simple to setup and deploy.

For local development, you can use `npm run dev` to build the frontend and run the server locally,
changes will be automatically reloaded.

For production, you can use `npm run build` to build the frontend.
This will create a production build in the `dist` directory.
You will then deploy this to your server as you would a static site.