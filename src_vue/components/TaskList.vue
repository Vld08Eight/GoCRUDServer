<template>
  <div>
    <div class="task-form">
      <task-form @create="addTask" />
    </div>
    <ul>
      <li v-for="task in tasks" :key="task.id">
        <task-item :key="task.id"
      :task="task"
      @task-checked="updateTask"
      @task-delete="deleteTask"
      />
      </li>
    </ul>
  </div>
</template>

<script>
import axios from 'axios';
import TaskItem from './TaskItem.vue';
import TaskForm from './TaskForm.vue'; // Adjust the path as necessary

export default {
  components: {
    TaskItem,
    TaskForm
  },
  data() {
    return {
      tasks: [],
      newTask: {
        title: '',
        body: '',
        completed: false,
      },
    };
  },
  created() {
    this.fetchTasks();
  },
  methods: {
    async addTask(newTask) {
      // Create a copy of the newTask object
      const taskToCreate = { ...newTask };
      await this.createTask(taskToCreate);
    },
    async fetchTasks() {
      const response = await axios.get('http://localhost:8080/tasks');
      this.tasks = response.data;
    },
    async createTask(task) {
      const response = await axios.post('http://localhost:8080/tasks', task);
      this.tasks.push(response.data);
      // Reset the newTask object after creation
      this.newTask = {
        title: '',
        body: '',
        completed: false,
      };
    },
    async updateTask(task) {
      await axios.put(`http://localhost:8080/tasks/${task.id}`, task);
      this.fetchTasks(); // Refresh the task list
    },
    async deleteTask(task) {
      await axios.delete(`http://localhost:8080/tasks/${task.id}`);
      this.fetchTasks(); // Refresh the task list
    },
  },
};
</script>

<style>
.task-form {
  max-width: 400px; /* Adjust the width as needed */
  margin: 15px 30px; /* Remove auto margin to align left */
  padding: 20px;
  border: 1px solid #ccc;
  border-radius: 8px;
  background-color: #f9f9f9;
  text-align: left; /* Ensure content inside the form is left-aligned */
}
</style>