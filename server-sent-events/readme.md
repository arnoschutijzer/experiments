```bash
$ npm install && npm start

# in a different terminal
$ curl -H Accept:text/event-stream http://localhost:3123/events

# in yet a different terminal
$ curl -X POST \
     -H "Content-Type: application/json" \
     -d '{"info": "Water is wet."}'\
     -s http://localhost:3123/fact
```

https://developer.mozilla.org/en-US/docs/Web/API/Server-sent_events/Using_server-sent_events

https://www.digitalocean.com/community/tutorials/nodejs-server-sent-events-build-realtime-app

Requires sticky sessions to be enabled & [redis in the backend?](https://redis.io/docs/interact/pubsub/).
