# Welcome to FleetDash

A *real-time* dashboard for your Eve Online fleets.
 
----

### Overview

FleetDash works by parsing player's combat logs and sending to the core server. 
This server then sends those events to authorized fleet members logged into the website, where they can be visualized in a fun lil dashboard.

This project organizes those 3 components into folders, `client`, `core`, and `web`. Which contains the source code for each

### How to use

First off, since Eve only logs combat events client side, *every* fleet member needs to run the FleetDash client locally.
This will parse those logs into a common format, and ship it to the core server. 

*Note: Logs will only be shipped when the toon  is logged in, and in an active session. 
If there is no one actively using the dashboard, the client will sit idle doing nothing.*

1) Download the latest release from the [Releases](https://github.com/Sheemap/fleet-dash/releases) section
2) Extract the zip file, and run `FleetDashClient.exe`
3) Once the UI opens, log in to each character you want to show up on the FleetDash dashboard. Only logged in toons will have their logs watched
4) You're done! Leave the application open, and just let the magic happen

Second off, we don't visualize any data locally, so you have to visit the dashboard to do so.

1) Go to https://fleetdash.space
2) Login with a character that is a part of your fleet
3) Click `Join Session` *If there is already an active session, this step will be done for you*
4) Voila! Soon after you've joined, you should see the widgets start to show data (if your fleet is actively doing things)

You can customize your dashbaord by drag-n-dropping and resizing all the widgets. Just hold click to drag, and hold click the bottom right corner of the widget to resize.

Some widgets have customizable settings, for these, click the gear icon in the top right to modify those settings.

Your dashboard will be saved locally, but if you open a new browser, or clear browser data, you will have to rebuild. (until some sort of backup feature is built)
