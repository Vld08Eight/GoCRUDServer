<template>
    <form @submit.prevent>
      <h4>Create task</h4>
      <my-input v-model="task.title" placeholder="Название" />
      <my-input v-model="task.body" placeholder="Описание" />
      <my-button class="btn" @click="createTask">Создать</my-button>
    </form>
  </template>
  
  <script>
  export default {
    data() {
      return {
        task: {
          title: '',
          body: '',
        },
      };
    },
    methods: {
      createTask() {
        // Assign the ID before emitting the task
        const newTask = {
          ...this.task,
          
          completed: false, // Add a default completed status
        };
        this.$emit('create', newTask); // Emit the new task with the ID
        // Reset the form
        this.task = {
          title: '',
          body: '',
        };
      },
    },
  };
  </script>

<style scoped>
form {
    display: flex;
    flex-direction: column;
}
.btn {
    align-self: flex-end;
    margin-top: 15px;
}

</style>