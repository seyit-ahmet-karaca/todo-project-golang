import Vuex from "vuex"
import Vue from "vue"
import API from "@/api"

Vue.use(Vuex)

export const state = {
    todoItems: []
}

export const getters = {
    getTodoItems(state) {
        return state.todoItems
    }
}

export const mutations = {
    setTodoItems(state, payload) {
        state.todoItems = payload;
    }
}

export const actions = {
    async fetchTodoItems({commit}) {
        const todoItems = await API.getTodoItems()
        commit("setTodoItems", todoItems)
    },
    async createItem(context, payload) {
        const todoItem = {title: payload}
        await API.createToDoItem(todoItem)
        await context.dispatch("fetchTodoItems")
    }
}

export default new Vuex.Store({
    state,
    getters,
    mutations,
    actions
})