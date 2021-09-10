Hello.
Firebase Cloud Messaging (FCM) is a cross-platform messaging solution.
Use it to send messages from the server -> to client devices.

This code (main.go) is the first steps, to connect FCM and send a message.

NOTE:
Do not use FCM if you need to send from client -> to server, because FCM uses legacy XMPP protocol (it's easier to use Rest API).
