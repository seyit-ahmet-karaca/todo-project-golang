Feature: TodoList
  As Product Owner I want to navigate to ToDoItem listing page
  So that I can see empty list and add a new todo item

  Scenario: Adding a TODO item
    Given Empty ToDo List
    When I write "buy some milk" to text box and click to add button
    Then I should see "buy some milk" item in ToDo list