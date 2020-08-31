# mapper
This is a work in progress Web Front-End to manage `DomainLinks`.

# What is a `DomainLink`?
A `DomainLink` is a way to map the collection of Feedback Sources a Domain will be looking for and also keep a record of the intended Destinations.

See the current a `DomainLink` design in `myob/contracts/newlink.json`

# Design
The idea behind this is:
1. Have static pages that will allow for CRUD of `DomainLinks`
2. On `CREATE` functions, establish a `DomainLink` and call the `processor` to set up the new link.
3. On `UPDATE` function, this simply edits the link and saves it back.
4. On `DELETE` function, simply nukes the `DomainLink` in it's entirety.

It's intended to keep `DomainLinks` as a JSON blob a DynamoDB table.

These pages will be kept in the `pages` directory and be named for their function, ie: `create.gohtml` or `edit.gohtml`.
