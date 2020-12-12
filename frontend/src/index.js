import("jquery").then($ => {
    $(() => {
        getAllTodo()

        const todoForm = $("#todo-form")
        const todoText = $("#todo-text")
        const todoId = $("#todo-id")
        const todoList = $("#todo-list")

        todoForm.submit(e => {
            e.preventDefault();

            const id = todoId.val()
            const text = todoText.val()

            if (id) return updateTodo(id, text);
            createTodo(text);
        })

        function getAllTodo() {
            graphql(`
                query getAllTodo {
                    allTodo {
                        id
                        text
                        done
                    }
                }
            `).done(({data}) => {
                const {allTodo} = data
                allTodo.forEach(todo => {
                    const item = makeTodoRow(todo)
                    todoList.append(item)
                })
            })
        }

        function makeTodoRow(todo) {
            const doneCheckbox = $("<input type='checkbox'>")
                .attr("checked", todo.done)
                .click(function () {
                    const elem = $(this)

                    graphql(`
                        mutation toggleTodoDoneStatus($id: ID!) {
                            toggleTodoDoneStatus(input: {id: $id}) {
                                id
                                text
                                done
                            }
                        }
                     `, {id: todo.id}).done(({data}) => {
                        const {toggleTodoDoneStatus} = data
                        elem.attr("checked", toggleTodoDoneStatus.done)
                    })
                });
            return $("<tr>")
                .data("id", todo.id)
                .append($("<td>").html(todo.text))
                .append($("<td>").append(doneCheckbox))
                .append($("<td>")
                    .append(makeUpdateTodoButton(todo))
                    .append(makeRemoveTodoButton(todo)))
        }

        function makeUpdateTodoButton(todo) {
            return $("<button>").html("Update").click(() => {
                todoId.val(todo.id)
                todoText.val(todo.text)
            })
        }

        function makeRemoveTodoButton(todo) {
            return $("<button>").html("Remove").click(() => {
                graphql(`
                    mutation deleteTodo($id: ID!) {
                        deleteTodo(input: { id: $id }) {
                            id
                            text
                            done
                        }
                    }
                `, {id: todo.id}).done(({data}) => {
                    const {deleteTodo} = data

                    todoList.children().each(function () {
                        const elem = $(this)
                        const id = elem.data("id")
                        if (id === deleteTodo.id) {
                            elem.remove()
                        }
                    })
                })
            })
        }

        function updateTodo(id, text) {
            graphql(`
                    mutation updateTodo($id: ID!, $text: String!) {
                        updateTodo(input: {
                            id: $id,
                            text: $text
                        }) {
                            id
                            text
                            done
                        }
                    }
                `, {id, text}).done(({data}) => {
                const {updateTodo} = data

                todoList.children().each(function () {
                    const elem = $(this)
                    const id = elem.data("id")
                    if (id === updateTodo.id) {
                        elem.children("td:eq(0)").html(updateTodo.text)
                    }
                })

                todoId.val('')
                todoText.val('')
            })
        }

        function createTodo(text) {
            graphql(`
                mutation createTodo($text: String!) {
                    createTodo(input: { text: $text }) {
                        id
                        text
                        done
                    }
                }
            `, {text}).done(({data}) => {
                const {createTodo} = data

                const item = makeTodoRow(createTodo)
                todoList.append(item)

                todoText.val('')
                todoText.focus()
            })
        }

        function graphql(query, variables) {
            return $.ajax({
                url: "/graphql",
                method: "POST",
                contentType: "application/json",
                data: JSON.stringify({query, variables})
            })
        }
    })
})
