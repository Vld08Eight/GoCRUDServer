<template>
  <div class="topbar">
    <div class="left-section">
      <slot name="left"></slot>
    </div>
    <div class="center-section">
      <slot name="center"></slot>
    </div>
    <div class="right-section">
      <slot name="right"></slot>
    </div>
  </div>
</template>

<script>
export default {
  name: 'TopBar',
};
</script>

<style scoped>
.topbar {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 10px 20px;
  background-color: #333;
  color: #fff;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
}

.left-section,
.center-section,
.right-section {
  display: flex;
  align-items: center;
}

.left-section {
  justify-content: flex-start;
}

.center-section {
  justify-content: center;
}

.right-section {
  justify-content: flex-end;
}
</style>