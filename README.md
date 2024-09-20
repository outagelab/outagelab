# OutageLab - resilience testing for everyone

OutageLab is a UI-driven resilience testing platform enabling remote-controlled outage simulation in running applications.

Other resilience/chaos testing tools are generally one or several of the following:

- Focused on infrastructure resilience with infrastructure specialists as the intended users
- Limited to supporting specific kinds of hosting environments (only AWS, only Kubernetes, etc.)
- Designed to scare people into good engineering by breaking things at random in production
- Intended for automated integration testing of a single service in isolation

OutageLab's focus is:

- To be easy for any developer, QA tester, and semi-technical PM to use
- To enable end-to-end resilience testing in prod-like environments (prod **not** recommended)
- To simulate outages on-demand for multiple use cases, including exploratory testing, debugging, acceptance testing, etc.
- To support more workloads by running anywhere your code does, and to avoid reliance on infra teams to set up
- To be incrementally adoptable by individual applications and teams

## How it works

OutageLab uses an application-embedded agent that installs as a package, with support for multiple languages. These agents rely on a backend server that acts as a controller, and outage simulations are managed from a simple UI. This application-embedded approach makes it easy for any developer to set up and use anywhere, whether it be locally, in a container, on bare-metal, in a serverless function, a data pipeline - anywhere.

The backend server is open source and can be self-hosted, but to support accelerated adoption and experimentation, a managed service is available at [app.outagelab.com](https://app.outagelab.com/). This service will remain free for individuals, and if the project gets enough traction, a competitively priced paid tier for businesses may be created to support the project's development and hosting costs.

## Get started

[Learn more at outagelab.com](https://outagelab.com/)
