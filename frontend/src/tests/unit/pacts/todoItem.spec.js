import {pactWith} from "jest-pact"
import {Matchers} from "@pact-foundation/pact"

const {eachLike, like} = Matchers
import {API} from "@/api";

pactWith({
    consumer: "todo-frontend",
    provider: "todo-backend"
}, provider => {
    describe("todo itemss", () => {
        let client;

        beforeEach(() => {
            client = new API(provider.mockService.baseUrl, false)
        });

        it("fetch to-do-items", async function () {
            await provider.addInteraction({
                state: "fetch items successfully",
                uponReceiving: "a request for fetch todo list",
                withRequest: {
                    method: "GET",
                    path: "/api/todo-items"
                },
                willRespondWith: {
                    status: 200,
                    headers: {
                        "Content-Type": "application/json; charset=UTF-8",
                    },
                    body: eachLike({
                        id: like(1),
                        title: like("Reading book"),
                    })
                }
            })

            const res = await client.getTodoItems()
            expect(res[0].id).toStrictEqual(1)
        });

        it("create todo list item", async function () {
            await provider.addInteraction({
                state: "create todo list item",
                uponReceiving: "a request to create todo list item",
                withRequest: {
                    method: "POST",
                    path: "/api/todo-item",
                    body: like({
                        title: like("do sport"),
                    }),
                    headers: {
                        "Content-Type": "application/json; charset=UTF-8",
                    }
                },
                willRespondWith: {
                    status: 201
                }
            })

            const toDoItem = {
                title: "do sport",
            }
            const res = await client.createToDoItem(toDoItem)
            expect(res).toStrictEqual(201)
        });

    })
})