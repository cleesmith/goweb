# GoWeb

A replacement for "python -m SimpleHTTPServer" in Go.

***

### Features

* no huge python installation and packages required -- in other words:
  - there's no fooling around with installing stuff
  - just plop/scp/wget/curl a goweb release, unzip, and off you go
* its executable binary footprint is 4 times smaller than SimpleHTTPServer
* there are no external dependencies -- everything is from Go's standard libraries
* it's fast enough for my <a href="http://cleesmith.github.io/" target="_blank">blog</a>
  - currently at github.io but with goweb it can be hosted anywhere
* pre-built binary <a href="https://github.com/cleesmith/goweb/releases" target="_blank">releases</a> for all OS/architecture's that Go provides

### ToDos

* input params
  - host
  - port
  - path to folder for serving html and static assets
    * instead of just defaulting to the current folder
  - log file name -- to use ```/var/log/goweb/whatever.log```
* upstart or init.d script -- to avoid doing ```nohup ./goweb &```
* use <a href="https://github.com/mpalmer/giddyup" target="_blank">giddyup</a> to update a website/blog via ```git push```
  - similar to ```github.io``` or ```heroku```
  - this isn't exactly related to goweb -- just a thought about easy updating

***
