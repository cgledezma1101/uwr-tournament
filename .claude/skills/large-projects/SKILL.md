---
name: large-projects
description: use this skill when the user wants to implement a large piece of software. This could be a major refactor of the whole codebase, implementing a large feature set, updating a framework that cuts across a large portion of the codebase, etc.
---

Goal: split up a large piece of work into manageable chunks that can be taken up by teams of agents so implementation is parallelised and efficient. We also want to make sure that the project can be tracked across interruptions, such as token exhaustion, network failures, or agent crashes.

Here's the framework you should follow:

# Step 1: Create a detailed plan, waterfall style

We want to start with a clear picture of requirements before jumping into planning. Don't make assumptions, always validate them with the user. Ask a lot of questions so that the output requires few iterations. Once you have a clear picture of the requirements, put together an implementation plan following these principles:

1. Implementation tasks should be small. The expected scope should be modifying no more than 10 files, or a few hundred lines of code.
2. We want to focus on capturing the definition of done of the tasks, rather than specific implementation. You can add implementation hints, but your goal is not to define an end to end implementation.
3. We should optimise task breakdown for parallelisation. Identify cross-cutting concerns early so they can be tackled soon and unblock large portions of coding. We don't want two tasks implementing the same thing as that can lead to conflicts in implementation

Once you have the tasks identified, present your plan for feedback. Iterate on this step as much as required until the user is happy with the plan

# Step 2: Build coordination framework

Once the plan is agreed on, you will capture the context required for each task, including a description for it, some context on why it's needed, and its definition of done, in the project's directory like so:

```
<project-name>/
  tasks/
    <task-id>-some-description.md
```

additional to that, you'll build a DAG representing the dependencies between the tasks, and store it in `<project-name>/task-dag.md`. We will use this DAG to track the progress of the project, and understand what tasks remain. We can encode the DAG as YAML using this format:

```
nodes:
  <task-id>:
    definitionFile: <path-to-defintion>
    status: # one of pending,in-progress,complete
edges:
  - source: <task-id>
    target: <task-id>
```

Validate that the DAG has no loops. If it does, break them by either merging tasks, changing their scope, or splitting further.

# Step 3: The build loop

Now that we have a coordination framework, you'll use teams of agents to deliver the implementation. The main guiding principle to build a large project is:

## Step 3.0: Find the next batch of tasks

Load the DAG definition and identify all the root nodes. Then go ahead with the next steps to perform those implementation tasks. In selecting how many tasks should be scheduled for development, you should be mindful of resources. Take into consideration how many tasks are currently in implementation, the pressure the system has, and the amount of tokens available. We prefer doing fewer tasks in parallel as long as we can complete them. Track your estimates of how much pressure the system can safely handle so you can be efficient

## Step 3.1: Start small

When tackling a task that's not similar to something you've seen before, solve it with a single agent for a single ticket.

## Step 3.2: Validate approach with the user

Once you have a solution ready, propose it to the user and ask for a review.

## Step 3.3: Learn and iterate

Update any skills that were loaded for the task with the feedback from the user. Capture any necessary context in configuration files, or in ticket definitions. The goal of this step is to make it less likely an agent will make the same mistake again.

Once you've done that, send an agent to implement the feedback with new artifacts. Then go back to 3.2 until the user approves the changes

## Step 3.4: Scale up

Once you have a validated approach, parallelise the implementation of all unblocked tickets. This step should be done independently, since we cannot expect the user to monitor teams of agents constantly. Here are some principles to follow:

1. If you find uncharted territory, go back to 3.1. Don't make assumptions, don't run wild with implementations. Be conservative, not wasteful
2. Every task should be implemented by an agent with an adversarial reviewer in its own git worktree.
    1. The agent should create an implementation loading appropriate skills and stop only when the definition of done is met
    2. Once that is done, it should hand over to an adversarial agent that should try to validate that the implementation meets the definition of done without looking at the actual implementation. Some strategies they can use are: make API calls to a local server, analyse completeness of the test suites, spin up a headless browser and validate user flows. They should attempt to break the implementation by finding issues around the limits of the parameters
    3. The adversarial agent should give any feedback to the agent, and start again on step 1. Only once the adversarial agent is content should the task be handed back to the main agent for reconciliation.

## Step 3.5: Reconcile and continue

Once an agent gives back a completed ticket, the main orchestrator should merge their changes into the main worktree, resolve any merge conflicts, validate that all automated tests pass, and that the application can be launched end to end. If there are failures, send back to the agent for rebasing and rework, starting step 3.4 again.

If all looks green, push the changes to main and update the task DAG to remove the nodes and edges that are complete. If there are still nodes remaining, then go back to 3.0. Otherwise, you're done!