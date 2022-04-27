<template>
  <div id="app">
    <todo-item-page></todo-item-page>
  </div>
</template>

<script>

import TodoItemPage from "@/views/TodoItemPage";

export default {
  name: 'App',
  components: {
    TodoItemPage
  }
}
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
