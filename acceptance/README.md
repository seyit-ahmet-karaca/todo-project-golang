# TodoList Acceptance Application

TodoList acceptance application is automated tests. It performs the specified tests.

The tests scenario written in Cucumber and positioned at ``features/todolist.feature`` file.

The steps related with scenario written in Puppeteer. 

### Scenario
Adding a TODO item

#### steps 
- [ ] Given: Empty ToDo List
  * Opens [TodoList test environment UI page]("http://34.75.50.180")
  * ensure todo list is empty.
- [ ] When : I write "buy some milk" to text box and click to add button
  * Search textbox can be enter some text
  * Click add button
- [ ] Then : I should see "buy some milk" item in ToDo list
  * See "buy some milk" item in Todo list

## Links
* [Cucumber](https://cucumber.io/)
* [Puppeeteer](https://pptr.dev/)
