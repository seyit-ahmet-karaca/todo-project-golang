import {shallowMount} from "@vue/test-utils";
import TodoItemPage from "../TodoItemPage.vue";
import ToDoItem from "../../components/ToDoItem.vue";

jest.mock("@/api")

function mountConfig(mockDispatch, customItems) {
    const items = []
    return {
        mocks: {
            $store: {
                getters: {
                    "getTodoItems": customItems === undefined ? items : customItems
                },
                dispatch: mockDispatch
            },
        }
    }
}


describe("TodoItemPage.vue", function () {
    it("should exists", function () {
        const mockDispatch = jest.fn()
        const wrapper = shallowMount(TodoItemPage, mountConfig(mockDispatch))

        expect(wrapper.exists()).toBeTruthy()

        const mainContainer = wrapper.find("#todo-item-main-container");
        expect(mainContainer.exists()).toBeTruthy()

        const addingItemContainer = mainContainer.find("#add-todo-item-container");
        expect(addingItemContainer.exists()).toBeTruthy()

        const title = addingItemContainer.find("#title");
        expect(title.exists()).toBeTruthy()
        expect(title.text()).toStrictEqual("Item Title")

        const inputItem = addingItemContainer.find("#input-item");
        expect(inputItem.exists()).toBeTruthy()

        const addButton = addingItemContainer.find("#add-item");
        expect(addButton.exists()).toBeTruthy()
        expect(addButton.text()).toStrictEqual("Add")

        const todoComponents = mainContainer.findAllComponents(ToDoItem);
        expect(todoComponents).toHaveLength(0)

        expect(mockDispatch).toHaveBeenCalled()
        expect(mockDispatch).toHaveBeenLastCalledWith("fetchTodoItems")
    });

    it("Todo Items list correctly", function () {
        const items = [
            {
                id: 1,
                title: "test"
            }
        ]

        const mockDispatch = jest.fn()
        const wrapper = shallowMount(TodoItemPage, mountConfig(mockDispatch, items))

        const todoItemComponents = wrapper.findAllComponents(ToDoItem);
        expect(todoItemComponents).toHaveLength(items.length)

        expect(mockDispatch).toHaveBeenCalled()
        expect(mockDispatch).toHaveBeenLastCalledWith("fetchTodoItems")
    });

    it("should Add item button clickable", async function () {
        const items = [
            {
                id: 1,
                title: "test"
            }
        ]

        const mockDispatch = jest.fn()
        const wrapper = shallowMount(TodoItemPage, mountConfig(mockDispatch, items))

        const mockAddItem = jest.fn()
        wrapper.setMethods({"addItem": mockAddItem})

        const addButton = wrapper.find("#add-item");
        await addButton.trigger("click")

        expect(mockAddItem).toHaveBeenCalled()
    });

    it("should addItem method functionality", async function () {
        const testItemValue = "test"
        const mockDispatch = jest.fn()
        const localThis = {
            itemValue: testItemValue,
            createItem : mockDispatch
        }

        TodoItemPage.methods.addItem.call(localThis)

        expect(mockDispatch).toHaveBeenLastCalledWith( testItemValue)
    });
})