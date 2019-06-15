# hello-world

as minimal as I could make it, using golang and docker.

# assuming you have aws

1. ssh into your aws instance
2. install docker
3. ```docker run --rm -p 8080:8080 isakbm/hello-world-server```
4. make sure to allow inbound traffic on port 8080

# cloning and building

1. ```git clone https://github.com/isakbm/hello-world```
2. ```./build.sh```

running ``./build.sh`` will prompt you for sudo, and subsequently for login to isakbm dockerhub.
you should edit the ``./build.sh`` to use your dockerhub username if you have one... its free ;)
