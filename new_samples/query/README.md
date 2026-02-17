<!-- THIS IS A GENERATED FILE -->
<!-- PLEASE DO NOT EDIT -->

# Query Sample

## Prerequisites

0. Install Cadence CLI. See instruction [here](https://cadenceworkflow.io/docs/cli/).
1. Run the Cadence server:
    1. Clone the [Cadence](https://github.com/cadence-workflow/cadence) repository if you haven't done already: `git clone https://github.com/cadence-workflow/cadence.git`
    2. Run `docker compose -f docker/docker-compose.yml up` to start Cadence server
    3. See more details at https://github.com/uber/cadence/blob/master/README.md
2. Once everything is up and running in Docker, open [localhost:8088](localhost:8088) to view Cadence UI.
3. Register the `cadence-samples` domain:

```bash
cadence --domain cadence-samples domain register
```

Refresh the [domains page](http://localhost:8088/domains) from step 2 to verify `cadence-samples` is registered.

## Steps to run sample

Inside the folder this sample is defined, run the following command:

```bash
go run .
```

This will call the main function in main.go which starts the worker, which will be execute the sample workflow code

## Query Samples

This folder contains samples demonstrating how to use Cadence queries with MarkDoc-formatted responses. MarkDoc allows you to create interactive query responses with buttons that can signal workflows or start new workflows.

---

### Markdown Query Workflow

This workflow demonstrates basic MarkDoc query usage with signal buttons to control workflow execution.

```bash
cadence --domain cadence-samples \
  workflow start \
  --tl cadence-samples-worker \
  --et 1000 \
  --workflow_type cadence_samples.MarkdownQueryWorkflow
```

#### How to interact with the workflow

1. Go to the `cadence-samples` domain in cadence-web and click on this workflow
2. Click on the "Query" tab
3. Select the "Signal" query from the dropdown
4. Use the rendered buttons to:
   - **Complete** - Send a signal to complete the workflow
   - **Continue** - Send a signal to continue without completing
   - **Start Another** - Start a new instance of this workflow

---

### Lunch Vote Workflow

This workflow demonstrates a more complex MarkDoc query use case - an interactive lunch voting system. Users can vote for their preferred lunch option via signal buttons.

```bash
cadence --domain cadence-samples \
  workflow start \
  --tl cadence-samples-worker \
  --et 600 \
  --workflow_type cadence_samples.LunchVoteWorkflow
```

#### How to vote

1. Go to the `cadence-samples` domain in cadence-web and click on the workflow
2. Click on the "Query" tab
3. Select the "options" query from the dropdown
4. You'll see:
   - A table of current votes
   - A menu with food options and images
   - Three voting buttons (Farmhouse, Ethiopian, Ler Ros)
5. Click any button to cast your vote
6. Refresh the query to see updated vote counts

#### Key concepts demonstrated

- **Dynamic query responses**: The vote table updates as votes come in
- **Signal buttons with payloads**: Each vote button sends structured data (location, meal, requests)
- **MarkDoc formatting**: Tables, images, and interactive buttons in query responses

The workflow runs for 10 minutes, collecting votes before completing.

## References

* The website: https://cadenceworkflow.io
* Cadence's server: https://github.com/uber/cadence
* Cadence's Go client: https://github.com/uber-go/cadence-client

