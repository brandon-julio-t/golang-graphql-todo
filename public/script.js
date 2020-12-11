$(() => {
    const tableBodyTodos = $("#table-body-todos")
    const inputTodoText = $("#input-todo-text")
    const formTodo = $("#form-todo")

    formTodo.submit(e => {
        e.preventDefault()

        const text = inputTodoText.val()

        $.ajax({
            url: "/graphql",
            contentType: "application/json",
            method: "POST",
            data: JSON.stringify({
                query: `
                    mutation createTodo {
                        createTodo(input: {
                            text: "${text}",
                            userId: ":v :v :v"
                        }) {
                            id
                            text
                            done
                            user {
                                id
                                name
                            }
                        }
                    }
                `
            })
        }).done(({data}) => {
            const {createTodo} = data

            const row = makeTodoRow(createTodo)
            tableBodyTodos.prepend(row)

            inputTodoText.val('')
            inputTodoText.focus()
        })
    })

    fetchTodos();

    function fetchTodos() {
        $.ajax({
            url: "/graphql",
            contentType: "application/json",
            method: "POST",
            data: JSON.stringify({
                query: `
                    query getTodos {
                         todos {
                            id
                            text
                            done
                            user {
                                id
                                name
                            }
                        }
                    }
            `
            }),
        }).done(({data}) => {
            const {todos} = data
            todos.forEach(todo => {
                const row = makeTodoRow(todo)
                tableBodyTodos.append(row)
            })
        })
    }

    function makeTodoRow(todo) {
        const row = $("<tr>")

        row.append($("<td>").html(todo.id))
            .append($("<td>").html(todo.text))
            .append($("<td>").html($(`<input type="checkbox">`).attr("checked", todo.done)))
            .append($("<td>").html(todo.user.id))
            .append($("<td>").html(todo.user.name))

        return row
    }
})