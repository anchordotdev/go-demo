# README

A go-demo app to be used in Anchor doc examples to show how to setup TLS with Anchor. 

This example app has two endpoints:

* `/` mimics a web application by rendering a simple HTML page. When the app is configured to ping a "backend" service over https the response body will be displayed in the html document.
* `/api` mimics an API endpoint by returning a simple JSON response.

This app is referenced in the followling docs:

* https://anchor.dev/docs/examples/local-https/app-go 
* https://anchor.dev/docs/examples/service-to-service-tls/service-to-service-go

