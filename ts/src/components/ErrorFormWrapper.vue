<template lang="pug">
  .wrapper(:class="hasError ? 'has-error' : null")
    slot
    ul.error-messages(v-if="hasError")
      li.text-danger(v-for="v in myErrors") {{ v }}
</template>

<script lang="ts">
  import Vue from 'vue'
  import { ErrorMessages } from '../types'

  export default Vue.extend({
    props: {
      errors: Object as () => ErrorMessages,
      name: String,
    },
    computed: {
      hasError (): number {
        return this.myErrors.length
      },
      myErrors (): string[] {
        return this.errors[this.name] || []
      },
    },
  })
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
  ul.error-messages {
    margin-top: 5px;
    padding: 0;
  }

  ul.error-messages li {
    margin-left: 20px;
    margin-top: 10px;
    font-weight: bold;
  }
</style>
