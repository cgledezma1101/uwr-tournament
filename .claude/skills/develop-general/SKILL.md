---
name: develop-general
description: Use this skill whenever the user asks to write an implementation that requires you to write code, irrespective of programming language
---

This skill provides the general guiding principles of developing software, independently of programming language. This skill should be used in tandem with the `develop-*` skill for the specific language you're writing code in

# Guiding principles

When writing code we want to stick to the following:

1. Avoid introducing new technologies or idioms. Take context from the repository to identify how this repo works and is built, and prefer using what exists, rather than building something new
2. Make scoped incremental changes. Avoid touching files adjacent to what you're doing, unless you're explicitly doing a refactoring task. You should write the smallest amount of code that gets you to the definition of done
3. Follow a strict test-driven development approach. When tackling a task, write the main abstractions that will be required to deliver the task. Then figure out if there are any tests that already cover the functionality you'll write, and update them so they become red for your use cases. Only if there are no applicable tests should you write new ones, which should also be red. If the tests you wrote are not red, then you wrote the wrong tests. Tests should be driven by behaviour, not implementation
4. Once you have red tests, write the smallest amount of code that will make them green
5. Run the tests again and validate that they're green. If they aren't go back to coding.
6. It's ok to have missed some tests in (3). If your implementation introduced functionality that is not covered with tests, write those tests and make sure they're green
7. Checkpoint often. Use git commits liberally to make sure work is constantly being stored. Token availability is an issue, so a task may have to be stopped half way. When git commit history is not enough, keep a markdown file with the tasks you need to complete, and context you've gathered, and update it as you go so work can be resumed seamlessly
