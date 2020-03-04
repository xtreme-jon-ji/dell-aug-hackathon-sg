# PAS Cheatsheet

## Table of Contents

[Install the CF CLI](#install-the-cf-cli)

[Deploying the Sample Application](#deploying-the-sample-application)
- [Download examples](#download-examples)
- [Accessing the Platform](#accessing-the-platform)
- [Building and pushing applications to Cloud Foundry](#building-and-pushing-applications-to-cloud-foundry)
  - [ReactJS app](#reactjs-app)
  - [Python app](#python-app)
  - [Go app](#go-app)
  - [Spring app](#spring-app)
- [Once apps are running](#once-apps-are-running)
  - [View App status](#view-app-status)
  - [View Logs](#view-logs)
- [Connect a Database or other service](#connect-a-database-or-other-service)
  - [Connecting a Database](#connecting-a-database)
  - [Connect an external service](#connect-an-external-service)
  
[Links](#links)

[Advanced](#advanced)

[Acknowledgements](#acknowledgements)


## Install the CF CLI
Download and install the Cloud Foundry Command Line Interface (cf CLI):

https://github.com/cloudfoundry/cli/releases/tag/v6.49.0

CLI reference guide can be found at: https://docs.cloudfoundry.org/cf-cli/cf-help.html

## Deploying the Sample Application
### Download examples
For `go`, `python`, and `ReactJS` examples, clone or download this repository:

```
$ git clone git@github.com:xtreme-jon-ji/dell-aug-hackathon-sg.git
```

For `spring/java`, see https://github.com/cloudfoundry-samples/spring-music


### Accessing the Platform
First, you'll need to login to the platform:
```
$ cf login -a <cf_api_url> -u <username> -p <password>
```

The `cf_api_url`, `username`, `password` will be provided to you at the event.

In the meantime, if you're looking to experiment with cf and try the following examples, check out Pivotal Web Services:
- PWS: https://run.pivotal.io
- Tutorial (which this cheatsheet is based upon): https://pivotal.io/platform/pcf-tutorials/getting-started-with-pivotal-cloud-foundry/introduction

### Building and pushing applications to Cloud Foundry
**TLDR:** In most cases, just do `cf push`.


The sample applications contain the `manifest.yml` and supporting files necessary to deploy to Cloud Foundry.
By default `cf push` will use the `manifest.yml` in the current directory. You can specify a path to the manifest
that you want to use with the `-f` flag, or use `--no-manifest` to ignore the manifest file.

#### ReactJS app
1. Navigate to app directory
    ```
    $ cd demo-app-reactjs
    ```
1. Fetch dependencies
   ```
   $ npm install
   ```
1. Build
    ```
    $ npm run-script build
    ```
1. Push
   ```
   $ cf push
   ```

For more detailed notes, [click here](/examples/demo-app-reactjs/README.md)

#### Python app
1. Navigate to app directory
    ```
    $ cd demo-app-python
    ```
1. Push
    ```
    $ cf push
    ```

For more detailed notes, [click here](/examples/demo-app-python/README.md)

#### Go app
1. Navigate to app directory
    ```
    $ cd demo-app-go
    ```
1. Push
    ```
    $ cf push
    ```

For more detailed notes, [click here](/examples/demo-app-go/README.md)

#### Spring app
See https://github.com/cloudfoundry-samples/spring-music

### Once apps are running
#### View App status
View your apps:
```
cf apps
```

#### View Logs
View a snapshot of recent logs:

```
cf logs cf-demo --recent
```

Or, stream live logs:
```
cf logs cf-demo
```
#### Using the Apps Manager GUI
To view your apps, services, manage lifecycle using a GUI, you use your browser to navigate to:
```
    apps.<cf platform host>
```

The GUI is similar to the PWS web console.


### Connecting a Database or other service
#### Connect a Database
Pivotal Platform enables administrators to provide a variety of services on the platform that can easily be consumed by applications.

List available service offerings

```
$ cf marketplace
```

You'll see something like:

| service | plans | description | 
|---|---|---|                                                                                          
| p.mysql | db-small | Dedicated instances of MySQL |                                                                             

The available services will vary depending on the platform configuration. 

List the available MySQL plans:

```
$ cf marketplace -s p.mysql
```
Create a service instance with the free plan:

```
$ cf create-service p.mysql db-small my-mysql-instance-name
```
Bind the newly created service to the app:

```
$ cf bind-service <app name here> my-mysql-instance-name
```

Once a service is bound to an app, environment variables are stored that allow the app to connect to the service 
after a push, restage, or restart command.

Restage the app:

```
cf restage <app name here>
```

Verify the new service is bound to the app:

```
$ cf services
```

Verify the app has environment variables containing credentials to access the service:

```
$ cf env <app name here>
```

#### Connect an external service
If the platform marketplace doesn't have a service that you want, you can connect an external service
(eg. Amazon S3 bucket) by using the [create-user-provided-service](https://cli.cloudfoundry.org/en-US/cf/create-user-provided-service.html) command:

```
$ cf create-user-provided-service SERVICE_INSTANCE [-p CREDENTIALS] [-l SYSLOG_DRAIN_URL] [-r ROUTE_SERVICE_URL] [-t TAGS]
```

## Links

Topics to explore:
How Pivotal Platform Works
https://docs.pivotal.io/pivotalcf/concepts

Explore and download more Cloud Foundry sample apps
https://github.com/cloudfoundry-samples/

Cloud Foundry Developer Guide
https://docs.cloudfoundry.org/devguide/index.html

Cloud Foundry App Deployment Manifest Reference
https://docs.cloudfoundry.org/devguide/deploy-apps/manifest.html

Cloud Foundry CLI Reference
https://docs.cloudfoundry.org/cf-cli/cf-help.html

Buildpacks:
https://docs.cloudfoundry.org/buildpacks/

## Advanced
To view available buildpacks:
```
$ cf buildpacks
```

To map to specific routes:
https://docs.cloudfoundry.org/devguide/deploy-apps/manifest-attributes.html#routes

## Acknowledgements
Cheatsheet adapted from [Try Pivotal Platform on the Public Cloud](https://pivotal.io/platform/pcf-tutorials/getting-started-with-pivotal-cloud-foundry/introduction)