# Quickstart

This quickstart will walk you through adding the OutageLab client library to your application and running a network outage simulation. It should take 5-10 minutes.

### Pre-requisites

OutageLab has no infrastructure or hosting requirements, and it's optimized for use by developers of all skill levels. As long as you can install the library in your application, you can follow this guide.

::: tip
If you don't have an application to test with, you can still follow this guide using our [demo app](https://github.com/outagelab/demo).
:::

### Steps

- [Add the client library to your app](#step-1)
- [Choose a server and get an API key](#step-2)
- [Initiate a simulated outage from the UI](#step-3)

## Add the client library to your app {#step-1}

Client libraries are currently available in Go, Python, and Node.js. Select your language below.

### Install

:::tabs key:lang
== Go

<!--@include: ./snippets/go.md#install-->

== Python

<!--@include: ./snippets/python.md#install-->

== Node.js

<!--@include: ./snippets/nodejs.md#install-->

:::

### Setup

:::tabs key:lang
== Go

<!--@include: ./snippets/go.md#setup-->

== Python

<!--@include: ./snippets/python.md#setup-->

== Node.js

<!--@include: ./snippets/nodejs.md#setup-->

:::

## Choose a server and get an API key {#step-2}

- Sign in to [app.outagelab.com](https://app.outagelab.com) and select **API Keys** on the left.
- Copy your API key value and save it in an environment variable named `OUTAGELAB_API_KEY`.

## Initiate a simulated outage from the UI {#step-3}

- Start your application, either locally or hosted
- Go to [the Applications page](https://app.outagelab.com/applications) in OutageLab.
- You should see the name of your running application on this page. If not, your app may be misconfigured or isn't running. Check your app logs for error details.
- Click **Edit** on the new application to open the configuration page.
- Toggle on the environment you are targeting.
- Next to **Outage Rules** click **Add**.
- This will generate a rule for controlling client HTTP requests sent by your app (not HTTP endpoints served by your app). This is the only supported outage type (for now).
- Enter a value in the **Host** field matching the host DNS name of the server that you want to break requests to.
- Choose to enter either or both a bad **Response Status** or **Response Delay** value.
- Make sure both the rule and the environment are enabled (toggle switch) and press **Save**.
- HTTP/S requests sent to the host you specified should now be breaking according to the rule you defined.
- You can change your outage rules in the UI on the fly without restarting your running app. Note that changes may take up to 30 seconds to take effect.
- Now is your opportunity to evaluate the impact and observability of this type of outage, design better ways to fail, and validate that they work.
- Deploy enhancements with the confidence that they will reduce business costs and improve customer experience during your next outage!
