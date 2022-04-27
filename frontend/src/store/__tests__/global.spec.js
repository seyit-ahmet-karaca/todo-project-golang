import {actions} from "@/store/store";
import API from "@/api";

jest.mock("@/api")

describe("store", () => {
    describe("actions", () => {
        it("fetchTodoItems", async function () {
            const context = {
                commit: jest.fn()
            }

            const todoItems = [{
                "title": "testTitle",
            }]

            API.getTodoItems.mockResolvedValue(todoItems)

            await actions.fetchTodoItems(context)
            expect(context.commit).toHaveBeenCalled()
            expect(context.commit).toHaveBeenLastCalledWith("setTodoItems", todoItems)
        });

        it("createTodoItems", async function () {
            const context = {
                commit: jest.fn(),
                dispatch: jest.fn()
            }

            const todoItems = [{
                "title": "testTitle",
            }]

            API.createToDoItem.mockResolvedValue(201)
            API.getTodoItems.mockResolvedValue(todoItems)

            const todoItem = "testTitle"
            await actions.createItem(context, todoItem)
            expect(context.dispatch).toHaveBeenCalled()
        });
    })
})