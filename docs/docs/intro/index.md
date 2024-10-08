# What is OutageLab?

OutageLab is an open source reliability/chaos testing platform for software developers. It allows you to simulate application level issues such as failed or slow HTTP requests, and is managed remotely via a UI. You can experiment with failures interactively, without restarting or redeploying your application, and run tests absolutely anywhere your code runs.

### Can I use it?

As long as you can install a library in your code, you can use OutageLab. Libraries are currently available for Go, Python, and Node.js, with more to come.

Note that OutageLab's technical approach involves a backend server that hosts the UI and client API, but self-hosting this is entirely optional. The default option is app.outagelab.com, a managed service with free personal accounts and reasonably priced team plans. The backend software is also open source and free to self-host for those interested.

### What kinds of failure can it simulate?

At the current stage OutageLab only supports intercepting and modifying your app's _outgoing_ HTTP requests, including:

- Responding with a chosen HTTP failure status code
- Adding latency to HTTP requests

Additional failure modes will continue to be added depending on user demand.

### Do I have to add code to my app?

Only a line or two of code to configure and start it. After that, all failure injection is automated and only needs to be configured in the UI.

### How does it compare to other reliability/chaos testing tools?

OutageLab aims to be the easiest tool for software devs to adopt.

The majority of alternative options are hard for most devs to use, for reasons such as:

- Limitations in which platforms or environments they are compatible with
- Focusing heavily on infrastructure problems and concepts, which makes them hard for devs to use
- Integrating at the infrastructure level, which may require your infra team's support to set up
- Having expensive pricing models requiring involved cost evaluations and contracts

### Am I ready for chaos testing? Isn't it risky?

OutageLab shouldn't be compared with high stakes tools like Netflix's Chaos Monkey, which built a reputation for its approach of intentionally killing servers at random in production. In contrast, OutageLab's primary focus is on-demand failure simulation in non-production environments.
