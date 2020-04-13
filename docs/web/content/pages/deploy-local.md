---
title: Local Deployment
keywords: local
tags: [local]
sidebar: home_sidebar
permalink: deploy-local.html
summary: 
---

Micro is incredibly simple to spin up locally

## Install

From source

```
go get github.com/micro/micro
```

Release binary

```
# MacOS
curl -fsSL https://raw.githubusercontent.com/micro/micro/master/scripts/install.sh | /bin/bash

# Linux
wget -q  https://raw.githubusercontent.com/micro/micro/master/scripts/install.sh -O - | /bin/bash

# Windows
powershell -Command "iwr -useb https://raw.githubusercontent.com/micro/micro/master/scripts/install.ps1 | iex"
```

## Run

Running micro is as simple as typing `micro`.

```
micro
```

To run the stack without connecting to the network

```
micro server
```

## Verify

Check everythings working by using a few commands

```
# list local services

micro list services

# list network nodes

micro network nodes

# call a service

micro call go.micro.network Debug.Health
```

{% include links.html %}
