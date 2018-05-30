<template lang="pug">
  Form
</template>

<script lang="ts">
  import Vue from 'vue';
  import Form from './components/Form'

  export default Vue.extend({
    name: 'App',
    components: {
      Form
    }
  })
</script>

<style>
  #app {
    font-family: 'Avenir', Helvetica, Arial, sans-serif;
    -webkit-font-smoothing: antialiased;
    -moz-osx-font-smoothing: grayscale;
    text-align: center;
    color: #2c3e50;
    margin-top: 60px;
  }
</style>
