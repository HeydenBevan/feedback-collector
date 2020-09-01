# feedback-collector
A custom baked toolset that will collect and process feedback data from several sources

# Prior Warning
This is a super in-dev service that I'm baking with the use of Golang. I am still learning the ropes so there will no doubt be problems all over the place... Please bear with me and feel free to point out where I'm going wrong at any time.

When I start using this, it will only be collecting feedback for the `Sales & Purchases` SME Domain as this will be used for the Proof of Concept. Will be building out ways to create links for other Domains too.

# Get Up and Running on your Local Machine
We like the `GOPATH` in this house, take a look if you're not familiar: [How to Write Go Code (with GOPATH)](https://golang.org/doc/gopath_code.html)
Steps to set up are as follows:
1. Clone repo to local machine
2. cd into directory
3. Open directory in text editor/ide of choice
4. Profit

There is currently no config for a specific platform (IE: Linux/AMD64). So any machine should do as we are using the Std Lib.

Can Build using the `go build` command.

# Service design
The intial design will be to have a few services.

| Path | Subpath | Purpose |
| --- | --- | --- |
| `collector` | n/a | Feedback Collector Client; see the README in there for details |
| `myob` | n/a | MYOB-specific plumbing |
| `myob` | `contracts` | Contains JSON files that denote the contracts being used when sending to the `Processor` |
| `processor` | n/a | The Feedback Processing service. Data transformation happens here before passing it to an intended destination. |
| `processor` | `controllers` | Basic controllers for the server, contains state info |
| `processor` | `forum` | Forum package, provides the plumbing for processing Forum-based feedback |
| `processor` | `linker` | Linker package, provides the front-end interface to manage `DomainLinks` |
