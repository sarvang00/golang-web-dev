# Creating an HTTP session
HTTP is a stateless protocol, which means each time a client retrieves a web page, the client opens a separate connection to the server and the server responds to it without keeping any record of the previous client request. So, if we want to implement a mechanism where the server knows about a request that the client has sent to it, then we can implement it using a session.

When we are working with sessions, clients just need to send an ID and the data is loaded from the server for the corresponding ID. There are three ways that we can implement this in a web application:
- Cookies
- Hidden form fields
- URL rewriting

**We will be using gorilla's sessions package to implement session.**