<template>
  <div id="app">
    <TopBar>
      <template v-slot:left>
        <div style="display: flex; flex-direction: column; align-items: center;">
          <span style="font-size: 28px; font-weight: bold;">TASK</span>
          <span style="font-size: 18px; margin-top: -10px;">manager</span>
        </div>
</template>
      <template v-slot:center>
        <!-- Nothing in the center -->
      </template>
      <template v-slot:right>
        <teal-button @click="goToAbout" style="margin-right: 10px;">About</teal-button>
        <teal-button @click="goToContact">Contact</teal-button>
      </template>
    </TopBar>
    <TaskList />
  </div>
</template>

<script>
import TaskList from './components/TaskList.vue';
import TopBar from './components/TopBar.vue';
export default {
  components: {
    TaskList, TopBar
  },
};
</script>

<style>
/* Add your styles here */
</style>