# mqtt-restapi-adapter
adapter between REST API and MQTT implemented on golang.

mqtt-restapi-adapter act as MQTT client which subscribes to certain topic in MQTT server. This project scopes are:
1. As an adapter, the script relays message **from** MQTT message in certain topic **to** associated REST API. In this case, everytime adapter receive message(s) from MQTT (which it already subscribed to), it then perform certain REST API call to REST API server (web server).
2. As an adapter, the script relays message **from** associated REST API **to** MQTT message in certain topic. In this case, the adapter act as REST API server (web server). Everytime adapter receive REST API call, it then publish the message to certain topic in MQTT server.
