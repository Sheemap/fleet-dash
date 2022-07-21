## Fleet Dash Client

The client is what watches a player's log directory, parses logs as they are written, and ships them off to the core server.

### TODO

* Login using the PKCE flow, so we don't need to store any client secret
* Fix the old failing tests

### Random disorganized thoughts

This is written using Electron.NET and Blazor. This was an experiment for me as I was exploring Blazor, and wanted to make the client cross-platform. I'm not really happy with the UI architecture. It would be nice to rewrite it in something like Tauri or MAUI. Or honestly even a headless daemon with a web ui for configuring would be satisfactory. This part of the code is just not a pleasure to develop in.
