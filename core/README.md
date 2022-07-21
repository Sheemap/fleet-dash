## Fleet Dash Core

This is the backbone of the system, but is pretty simple in concept. Core simply injests logs that clients will send, and spits them back out to any currently subscribed websockets. It handles authorization using the Eve ESI, so no one is able to access fleet data for fleets they arent a part of.

### TODO

* Improve the method for gathering events to send to websockets (Currently just continuously polls the DB)
* Add support for more than just CockroachDB

### Random disorganized thoughts

This was my first "real" app written in Golang, and first time workign with Gokit period. There are certainly some improvements that could be made simply just because I didn't (and still don't) fully understand the best way to architect this solution.

The polling for new events bugs me. It was a quick and easy way to get it done while still remaining stateless. But it likely wont scale well, and just isnt a pretty solution. I would love to move to something like RabbitMQ or some other messaging platform. Though this certainly adds complexity.
