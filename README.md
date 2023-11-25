# README

A go-demo app to be used in Anchor doc examples to show how to setup TLS with Anchor. 

The app has two endpoints:
* "/" mimics a web application and renders a simple HTML page. When the app is configured to ping a "backend" service over https the response body will be displayed in the html document.
* "/api" returns a simple json response to mimic an API endpoint. 


This app is referenced in the followling docs:

* https://docs.anchor.dev/examples/app-go
