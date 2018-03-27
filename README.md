# react-static-serve

Simple library to serve react statically built react application.

All it does is serving index.html instead of 404 status page.

# Docker image

suppose `build` is folder where your React application is generated.

```
FROM jakubknejzlik/react-static-serve

COPY build .
```
