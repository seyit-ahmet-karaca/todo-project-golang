const {Given, When, Then} = require("cucumber")
const openUrl = require("../support/action/openUrl")
const checkElementExists = require("../support/check/checkElementExists")
const sendKeys = require("../support/action/sendKeys");
const clickElement = require("../support/action/clickElement");

Given("Empty ToDo List", async function () {
    await openUrl.call(this, "http://34.75.50.180")

    const addToDoItemSelector = "#add-todo-item-container"
    await checkElementExists.call(this, addToDoItemSelector)

    const itemSelector = ".todo-item"
    const toDoItemLength = await this.page.$$(itemSelector)

    if (toDoItemLength.length !== 0) {
        throw "ToDoList is not empty"
    }
})

When(/^I write "([^"]*)" to text box and click to add button$/, async function (itemTitle) {
    await sendKeys.call(this, "#input-item", itemTitle)
    await clickElement.call(this, "#add-item")
});

Then(/^I should see "([^"]*)" item in ToDo list$/, async function (itemTitle) {
    // I could not find method like waitForNetworkIdle. So I use waitForTimeout instead
    await this.page.waitForTimeout(1000)
    await this.page.$$eval(".todo-item",
        async (items, itemTitle) => {
            const itemFound = items.find(item => item.textContent === itemTitle);
            if (!itemFound) {
                throw "There is not item such title" + itemTitle
            }
        },
        itemTitle
    )
});