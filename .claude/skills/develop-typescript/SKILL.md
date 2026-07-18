---
name: develop-typescript
description: Use this skill whenever a coding task requires implementing code using Typescript or JavaScript
---

This skill is triggered with both Typescript and JavaScript tasks. If the user wants to use JavaScript, just use Typescript instead. We don't want any JavaScript in this code base

# Guiding principles

1. We use Typescript in strict mode from the get-go. No `any`, no non-null assertions. If the type system says something can't be validated, then code must add the appropriate checks to make the type system pass. We don't cheat the type system into thinking things are ok because of external context. Think you're writing Kotlin
2. When untyped structures come at the boundaries, we use Zod to define a schema for it and validate the expected structure. We use Zod in strict mode. If something doesn't pass Zod validation, we fail loud and early.
3. Everything that the package needs to run and develop should be captured in `package.json`. We don't run dependencies using `npx` unless they're one-off. Common tasks should be captured as `npm` scripts for consistent evaluation
4. Our unit testing library is `jest`
5. Our network calls are done through `fetch`
6. When unit testing a package that has other package dependencies, prefer mocking the entire dependency package over testing e2e behaviour
7. When unit testing a package that does network calls, use `nock` to mock the networking layer, and validate the network call itself, rather than mocking `fetch`
8. When adding a new dependency, make sure you're bringing in the latest version (no migration tech-debt from the get-go), and use SemVer to allow auto-updates up to minor version.
9. When unit testing a package that relies on the runtime clock for some implementation, mock the clock using jest mock timers in your tests. Don't write tests that wait for wall-clock to pass.