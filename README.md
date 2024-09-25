# OutageLab - Resilience testing you can actually use

OutageLab is a UI-driven, application-embedded resilience testing platform enabling remote-controlled outage simulation anywhere your code runs.

> [!NOTE]
> OutageLab is still in early development and considered experimental. Use only in low-risk environments.

## Focus

- Resilience testing for software developers and QA testers, not infrastructure specialists
- To run _anywhere_ your code does, without any platform-specific setup requirements
- Manual testing in running apps, especially in prod-like environments, for usecases including:
  - discovering single points of failure in full-stack apps and distributed systems
  - debugging and acceptance testing resilience and observability enhancements
- UI-driven and remote-controlled, enabling on-demand testing without hardcoding errors and deploying them
- Safer, more controlled resilience/chaos testing - no random failure in production

## Setup

A [5 minute quickstart](https://outagelab.com/docs/intro/quickstart/) can be found here.

TLDR on setup requirements:

- Adding support for OutageLab in your application requires installing the `outagelab` library.
- Supported languages currently include Python, Go, and Node.js. More to come.
- The default recommended host for the UI / backend is [app.outagelab.com](https://app.outagelab.com). This is the fastest way to start, and is forever free for personal use and evaluation. Paid licenses for teams will be introduced if the tool gets traction, to help fund development and hosting.
- You can optionally self-host the backend service, which is free and open source software.
