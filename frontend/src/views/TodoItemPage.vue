<template>
  <div id="todo-item-main-container">
    <div id="add-todo-item-container">
      <p id="title">Item Title</p>
      <input type="text" id="input-item" v-model="itemValue">
      <button id="add-item" @click="addItem">Add</button>
    </div>

    <to-do-item
        v-for="item of getTodoItems"
        :key="item.id"
        :item="item"
    />
  </div>

</template>

<script>
import ToDoItem from "../components/ToDoItem.vue";
import {mapGetters, mapActions} from "vuex"

export default {
  name: "TodoItemPage",
  components: {ToDoItem},
  data() {
    return {
      itemValue: ""
    }
  },
  methods: {
    ...mapActions(["fetchTodoItems", "createItem"]),
    addItem() {
      this.createItem(this.itemValue)
    }
  },
  computed: {
    ...mapGetters(["getTodoItems"])
  },
  created() {
    this.fetchTodoItems()
  }
}
</script>

<style scoped>

</style>