<template>
    <button
      class="teal-button"
      @mouseover="isHovered = true"
      @mouseleave="isHovered = false"
      @mousedown="isActive = true"
      @mouseup="isActive = false"
      :style="buttonStyle"
    >
      <slot></slot>
    </button>
  </template>
  
  <script>
  export default {
    name: 'teal-button',
    data() {
      return {
        isHovered: false,
        isActive: false,
      };
    },
    computed: {
      buttonStyle() {
        let backgroundColor = '#20c997'; // Мягкий бирюзовый цвет
        if (this.isHovered) {
          backgroundColor = '#1aa179'; // Темнее при наведении
        }
        if (this.isActive) {
          backgroundColor = '#168a6b'; // Еще темнее при нажатии
        }
        return {
          backgroundColor,
          color: '#ffffff', // Белый текст
          borderRadius: '25px', // Закругленные углы
          padding: '10px 20px',
          border: 'none',
          cursor: 'pointer',
          transition: 'background-color 0.3s ease', // Плавное изменение цвета
          fontSize: '16px',
          outline: 'none',
        };
      },
    },
  };
  </script>
  
  <style scoped>
  .teal-button {
    /* Базовые стили, которые будут переопределены computed-свойством */
    background-color: #20c997;
    color: #ffffff;
    border-radius: 25px;
    padding: 10px 20px;
    border: none;
    cursor: pointer;
    transition: background-color 0.3s ease;
    font-size: 16px;
    outline: none;
  }
  </style>