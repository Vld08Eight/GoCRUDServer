<template>
  <div class="task">
    <div class="task_left">
      <input type="checkbox" v-model="task.completed" @change="onCheckboxChange" />
    </div>
    <div class="task_center">
      <div><strong>{{ task.title }}</strong></div>
    </div>
    <div class="task_right">
      <button @click="toggleDescription" class="arrow-button">
        {{ isDescriptionVisible ? '▲' : '▼' }}
      </button>
    </div>
    <div v-if="isDescriptionVisible" class="task_description">
      <div class="description-content">
        <div><strong>Описание:</strong></div>
        <div>{{ task.body }}</div>
      </div>
      <div class="delete-button-container">
        <teal-button @click="deleteTask" class="delete-button">Delete</teal-button>
      </div>
    </div>
  </div>
</template>

<script>
import TealButton from './UI/TealButton.vue';
export default {
  components: {
    TealButton,
  },
  props: {
    task: {
      type: Object,
      required: true,
    }
  },
  data() {  
    return {
      isDescriptionVisible: false,
    };
  },
  methods: {
    deleteTask() {
      this.$emit("task-delete", this.task)
    },
    toggleDescription() {
      this.isDescriptionVisible = !this.isDescriptionVisible;
    },
    onCheckboxChange() {
      this.$emit('task-checked', this.task);
    }
  }
};
</script>

<style scoped>
.task {
  padding: 15px;
  border: 2px solid teal;
  margin-top: 15px;
  display: flex;
  align-items: center;
  justify-content: space-between;
  flex-wrap: wrap;
}
.task_left {
  flex: 0 0 auto;
}
.task_center {
  flex: 1;
  padding: 0 10px;
}
.task_right {
  flex: 0 0 auto;
}

.arrow-button {
  background: none;
  border: none;
  cursor: pointer;
  font-size: 16px;
}

.task_description {
  width: 100%;
  margin-top: 10px;
  padding: 10px;
  border-top: 1px solid #ccc;
}

.description-content {
  margin-bottom: 10px; /* Отступ между описанием и кнопкой */
}

.delete-button-container {
  text-align: right; /* Выравнивание кнопки по правому краю */
}

</style>