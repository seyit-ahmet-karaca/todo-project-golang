<template>
  <p class="todo-item">{{ item.title }}</p>
</template>

<script>
export default {
  name: "ToDoItem",
  props: {
    item: Object
  }
}
</script>

<style scoped>

</style>