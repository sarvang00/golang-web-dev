# Managing HTTP session using Redis
While working with the distributed applications, we probably have to implement stateless load balancing for frontend users. This is so we can persist session information in a database or a filesystem so that we can identify the user and retrieve their information if a server gets shut down or restarted.

**We will be solving this problem using Redis as the persistent store to save a session.**