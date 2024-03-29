# Temporal + Benthos - Proof of Concept

This project demonstrates how Benthos can be used to drive Temporal workflow execution.

## Basic instructions

### Step 0: Temporal Server

Make sure [Temporal Server is running](https://docs.temporal.io/docs/server/quick-install/) first:

```bash
git clone https://github.com/temporalio/docker-compose.git
cd  docker-compose
docker-compose up
```

Leave it running. You can use the Temporal Web UI at [localhost:8080](localhost:8080) which is currently in Beta. To use the legacy Temporal Web UI, use the [localhost:8088](localhost:8088) URL instead. There should be no workflows visible in the dashboard right now.

### Step 1: Clone this Repository

In another terminal instance, clone this repo and run this application.

```bash
git clone https://github.com/disintegrator/benthos-temporal-poc.git
cd benthos-temporal-poc
```

### Step 2: Run the workflow with Benthos

```bash
go run benthos/main.go --log.level debug -c config.generate.yml
```

Observe that Temporal Web reflects the workflow, but it is still in "Running" status. This is because there is no Workflow or Activity Worker yet listening to the `TRANSFER_MONEY_TASK_QUEUE` task queue to process this work.

There is also an alternate pipeline configuration which reads data from a file and uses lines within it to trigger executions. You can run it like so:

```bash
go run benthos/main.go --log.level debug -c config.file.yml
```

### Step 3: Run the Worker

In YET ANOTHER terminal instance, run the worker. Notice that this worker hosts both Workflow and Activity functions.

```bash
go run worker/main.go
```

Now you can see the workflow run to completion. You can also see the worker polling for workflows and activities in the task queue at [http://localhost:8080/namespaces/default/task-queues/TRANSFER_MONEY_TASK_QUEUE](http://localhost:8080/namespaces/default/task-queues/TRANSFER_MONEY_TASK_QUEUE).

<img width="882" alt="CleanShot 2021-07-20 at 17 48 45@2x" src="https://user-images.githubusercontent.com/6764957/126413160-18663430-bb7a-4d3a-874e-80598e1fa07d.png">

## What Next?

You can run the Workflow code a few more times with `go run benthos/main.go --log.level debug -c config.yml` to understand how it interacts with the Worker and Temporal Server.
