# Task Manager

This question has been inspired by multiple task runners / background job
runners like [resque](https://github.com/resque/resque),
[sidekiq](https://sidekiq.org/), [celery](www.celeryproject.org/),
[faktory](https://contribsys.com/faktory/) - https://github.com/contribsys/faktory ,
and even schedulers like [nomad](https://nomadproject.io/),
[kubernetes](https://kubernetes.io/).

You are given a set of tasks to run in the form of a JSON. Some tasks are
independent and can be run anytime. Some tasks are dependent on one or more
tasks, in which case, they can be run only after running the tasks they are
dependent on. For example

```json
[
   {
      "name": "task1"
   },
   {
       "name": "task2"
   },
   {
       "name": "task3",
       "dependsOn": [ "task1", "task2" ]
   }
]
```

task3 can be executed only after executing task1 and task2. And task1 and task2
can be executed in any order, it doesn't have to be the order in which it
appears in the dependsOn array. 

1. Given the JSON, write a task manager which runs only one task at a time.
Running the task means printing the task name and that's it. So, for the above
example, the valid outputs are

```bash
task1
task2
task3
```

```bash
task2
task1
task3
```

We might need a separate program to check this output itself, as there are
multiple possible valid outputs.

2. Every task defines a command with arguments to execute, like this

```json
{
  "name": "task1",
  "command": "echo",
  "args": ["this is a task"]
}
```

Now, running a task means printing the task's name and then running the command
with the given array of arguments. 

3. Tasks manager can now run multiple tasks at a given time. It takes input on
how many tasks it should execute at a a given time and then runs tasks
accordingly in parallel. 

---

Thoughts from colleagues:
1. Think about if and how to bring in intelligence
2. Instead of running commands - the task can be simply given an input about
the duration of the task and no details about the command or what task, and it
can simply give a schedule of how to run the tasks to complete all the tasks
in as little time as possible. This is similar to CPU scheduling problem I
guess. And think about the duration and visualize it using Gantt charts
