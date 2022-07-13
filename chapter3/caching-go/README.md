# Implementing caching in Go
Caching data in a web application is sometimes necessary to avoid requesting static data from a database or external service again and again. Go does not provide any built-in package to cache responses, but it does support it through external packages.

There are a number of packages, such as ```https://github.com/coocood/freecache``` and ```https://github.com/patrickmn/go-cache```, which can help in implementing caching and, in this recipe, we will be using the ```https://github.com/patrickmn/go-cache``` to implement it.