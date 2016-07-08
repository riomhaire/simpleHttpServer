# simpleHttpServer

This is a very simple http file server which only currently supports GET requests
and allows you to define a mapping of a URI to a file via a mapping file. Generally I use it for serving up json or xml during the development process.

The approach taken is slightly different from most simple go http servers which publish
from a specified root directory in the file system.

An example mapping.yml is given which shows how the very basic configuration
works:

~~~yml
port: 8080

mappings:
    - uri: /api/v1/test
      file: /tmp/a.json

    - uri: /api/v1/echo
      file: /tmp/b.json

    - uri: /api/v1/echo/c
      file: /tmp/c.json
~~~

You can put the config in a file called 'mapping.yml' and then run the server since its looks for this file in the current directory; or you can place is elsewhere and use the '-m' flag along with the filename and location.
 
