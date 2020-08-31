# collector
This is an automated client that will continue to poll the feedback sources.

# Design
This is currently a Local Client _only_. This will eventually become it's own automated service, but that will be set up once the main plumbing is sorted.
During intial PoC this will only make calls the MYOB Community Forum to make sure the overall plumbing is present.

The idea is:
1. `DomainLinks` are set up using the `mapper` front-end
2. Collector will fetch a list of `DomainLinks` during intial startup
3. Collector will call the sources in the `DomainLinks`/`sources` object
4. Results will be sent immediately to the `processor` for transformation and passing to the intended destination.

The `collector` will only be adding a wrapper to the data before calling the processor. It should not be transforming any data.
