const {Publisher} = require("@pact-foundation/pact")
const opts = {
    pactBroker: "https://sak-assignment.pactflow.io",
    pactBrokerToken: "oN-yv0nYlrYmxkcdpO-slQ",
    consumerVersion: "1.0.0",
    pactFilesOrDirs: ["./pact/pacts"],
    environments: ["environments"],
    tags: ["test"]
};

new Publisher(opts).publishPacts().then(
    response => console.log(response)
)