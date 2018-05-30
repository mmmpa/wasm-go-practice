<template lang="pug">
  article
    .col-sm-4.col-sm-offset-4
      h1 Message Form
      section
        p(v-if="preparing")
          i.fa.fa-pulse.fa-spinner.space
          | loading wasm...
        P(v-else)
          | fetching time: {{ fetchingTime }} ms <br/>
          | compilation time: {{ compilationTime }} ms
      form(v-if="!sent" @submit.prevent="submit")
        ErrorFormWrapper(:errors="errors" name="name")
          .form-group
            label Name
            input.form-control(
            type="text"
            v-model="name"
            placeholder="ex: Tsuyoshi Kaneita"
            :disabled="disabled"
            )
        ErrorFormWrapper(:errors="errors" name="email")
          .form-group
            label Mail address
            input.form-control(
            type="text"
            v-model="email"
            placeholder="ex: tk@example.com"
            :disabled="disabled"
            )
        ErrorFormWrapper(:errors="errors" name="body")
          .form-group
            label Message
            input.form-control(
            type="text"
            v-model="body"
            placeholder="ex: This message form is pretty good."
            :disabled="disabled"
            )
        .form-group.top-space
          button.btn.btn-success(:disabled="disabled")
            i.fa.fa-paw
            | Submit
      section(v-else)
        h2 Thank you.
        .form-group.top-space
          button.btn.btn-success(@click="reset")
            i.fa.fa-envelope
            | Create new message
</template>

<script lang="ts">
  import Vue from 'vue';
  import ax from 'axios'
  import ErrorFormWrapper from './ErrorFormWrapper'
  import snakeCase from 'snake-case'
  import loadAndCompile from '../wasm'

  function initialData() {
    return {
      name: '',
      email: '',
      body: '',
      errors: {},
      status: '',
    }
  }

  export default Vue.extend({
    name: 'Form',
    components: {
      ErrorFormWrapper
    },
    async created() {
      this.status = 'preparation'
      Object.assign(this, await loadAndCompile())
      this.status = ''
    },
    data() {
      return {
        fetchingTime: null,
        compilationTime: null,
        go: null,
        ...initialData(),
      }
    },
    computed: {
      disabled(): Boolean {
        return this.preparing || this.sending
      },
      sent(): Boolean {
        return this.status === 'sent'
      },
      sending(): Boolean {
        return this.status === 'sending'
      },
      preparing(): Boolean {
        return this.status === 'preparation'
      },
    },
    methods: {
      reset() {
        Object.assign(this, initialData())
      },
      parseErrors(raw) {
        this.errors = raw.reduce((a, {field, reason}) => {
          const key   = snakeCase(field)
          const store = a[key] || (a[key] = [])
          store.push(reason)
          return a
        }, {})
      },
      async validate() {
        const {
                name,
                email,
                body,
              } = this

        return JSON.parse(
          await this.go.run(
            JSON.stringify({name, email, body})
          )
        )
      },
      async submit() {
        const {
                name,
                email,
                body,
              } = this

        const {valid, errors} = await this.validate()

        if (!valid) {
          this.parseErrors(errors)
          return
        }

        this.status = 'sending'
        try {
          await ax.post(
            'http://localhost:1323/api/messages',
            {name, email, body},
          )
          this.status = 'sent'
        } catch ({response: {data: {errors}}}) {
          this.parseErrors(errors)
          this.status = ''
        }
      },
    },
  })
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
  button {
    width: 100%
  }

  .space {
    margin-right: 0.5em;
  }

  .top-space {
    margin-top: 30px
  }

  button .fa {
    margin-right: 0.5em;
  }
</style>
