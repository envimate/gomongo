# gomongo

Simple example implementation of an HTTP service connected to mongodb.

Used in the blog posts:

- https://blog.envimate.me/2016/03/10/gomongo-part-1/
- https://blog.envimate.me/2016/04/21/gomongo-part-2/


# Startup

## requirements

To run this example you need to have 

- docker 1.9+
- docker-compose 1.6.+


To see it in action, check it out

`git clone git@github.com:envimate/gomongo.git`

`cd gomongo`

Then run using `docker-compose` 

`docker-compose -f docker/docker-compose.yml up -d`

By default port 80 is mapped so you should see:

`curl http://localhost:80/test

got request /test`
