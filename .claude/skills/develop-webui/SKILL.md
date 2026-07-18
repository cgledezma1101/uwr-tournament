---
name: develop-webui
description: Use this skill whenever an implementation task requires the development of a user interface that will be delivered through a web browser
---

This is skill should be used in composition with the develop-typescript skill. All of those principles apply, and we add some more that are specific to web interfaces

# Guiding principles

1. Our web UI framework will be ReactJS
2. Our state management library will be Redux
3. All react components should be split into two:
   1. A stateless functional component that only renders state that is passed through props (no hooks, no lifecycle)
   2. A stateful wrapper that performs all of the state management, eventing, etc. and renders the SFC in reaction to that
4. When unit testing React components, we use `enzyme` to shallow render components. Unit tests should not render component trees end to end. Instead they should test the behaviour of the specific component, and validate that sub-components are rendered with the correct props.
5. Our UI packaging framework is Webpack.