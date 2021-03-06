# GoWeb

A replacement for "python -m SimpleHTTPServer" in Go.

***

### Features

* no huge python installation and packages required -- in other words:
  - there's no fooling around with installing stuff
  - just plop/scp/wget/curl a goweb release, unzip, and off you go (so to speak)
* its executable binary footprint is 4 times smaller than SimpleHTTPServer
* there are no external dependencies -- everything is from Go's standard libraries
* it's fast enough for my <a href="http://cleesmith.github.io/" target="_blank">blog</a>
  - currently at github.io but with goweb it can be hosted anywhere
* pre-built binary <a href="https://github.com/cleesmith/goweb/releases" target="_blank">releases</a> for all OS/architecture's that Go covers

### ToDos

* input params
  - host (optional)
  - port (optional)
  - path to folder for serving html and static assets (optional)
    * instead of just defaulting to the current folder
  - log file name -- to use ```/var/log/goweb/whatever.log``` (optional)
* provide an upstart or init.d script -- as opposed to doing ```nohup ./goweb &```

### aSides

* use <a href="https://github.com/mpalmer/giddyup" target="_blank">giddyup</a> to update a website/blog via ```git push```
  - similar to ```github.io``` or ```heroku```
* yes, you're right, a simple ```git pull``` works too
  - assuming your website/blog uses git source control
* simple is always best -- KisS -- ranDom capitaliZation is underRated
* goweb is not web nirvana or even life changing ...
  - I have often needed a simple web server for various reasons
  - it is portable with no dependencies
  - it has chicken chasin' speed
  - always easy to find and deploy thanks to github

***
