import {shallowMount} from "@vue/test-utils";
import ToDoItem from "../ToDoItem.vue";

describe("ToDoItem.vue", () => {
    it("should render correctly", function () {
        const propItem = {
            id : 1,
            title: "testTitle"
        }

        const wrapper = shallowMount(ToDoItem, {
            propsData:{
                item: propItem
            }
        });
        expect(wrapper.exists()).toBeTruthy()

        const itemDiv = wrapper.find(".todo-item");
        expect(itemDiv.exists()).toBeTruthy()
        expect(itemDiv.text()).toStrictEqual(propItem.title)
    });
})