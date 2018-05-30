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
        ErrorFormWrapper(:errors="parsedErrors" name="name")
          .form-group
            label Name
            input.form-control(
            type="text"
            v-model="name"
            placeholder="ex: Tsuyoshi Kaneita"
            :disabled="disabled"
            )
        ErrorFormWrapper(:errors="parsedErrors" name="email")
          .form-group
            label Mail address
            input.form-control(
            type="text"
            v-model="email"
            placeholder="ex: tk@example.com"
            :disabled="disabled"
            )
        ErrorFormWrapper(:errors="parsedErrors" name="body")
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
  import Vue from 'vue'
  import ax from 'axios'
  import ErrorFormWrapper from './ErrorFormWrapper'
  import snakeCase from 'snake-case'
  import { loadAndCompile } from '../wasm'
  import { ValidationResult, ErrorDetail, ErrorMessages, CompilationResult } from '../types'

  enum Status {
    Ready = 'Ready',
    Sent  = 'Sent',
    Sending  = 'Sending',
    Preparing = 'Preparing',
  }

  function initialData () {
    return {
      name: '',
      email: '',
      body: '',
      errors: [],
    }
  }

  function parseErrors (errors: ErrorDetail[]): ErrorMessages {
    return errors.reduce((a, { field, reason }) => {
      const key: string     = snakeCase(field)
      const store: string[] = a[key] || (a[key] = [])
      store.push(reason)
      return a
    }, {})
  }

  export default Vue.extend({
    name: 'Form',
    components: {
      ErrorFormWrapper,
    },
    async created () {
      this.status  = Status.Preparing
      const result = await loadAndCompile() as CompilationResult
      Object.assign(this, result)
      this.status = Status.Ready
    },
    data (): {
      name: string,
      email: string,
      body: string,
      errors: ErrorDetail[],
      status: Status,
      fetchingTime: number,
      compilationTime: number,
      go: any,
    } {
      return {
        fetchingTime: null,
        compilationTime: null,
        go: null,
        status: Status.Preparing,
        ...initialData(),
      }
    },
    computed: {
      disabled (): Boolean {
        return this.preparing || this.sending
      },
      sent (): Boolean {
        return this.status === Status.Sent
      },
      sending (): Boolean {
        return this.status === Status.Sending
      },
      preparing (): Boolean {
        return this.status === Status.Preparing
      },
      parsedErrors (): ErrorMessages {
        return parseErrors(this.errors)
      },
    },
    methods: {
      reset () {
        Object.assign(this, initialData())
      },
      async validate () {
        const {
                name,
                email,
                body,
              } = this

        return JSON.parse(
          await this.go.run(
            JSON.stringify({ name, email, body }),
          ),
        ) as ValidationResult
      },
      async submit () {
        const {
                name,
                email,
                body,
              } = this

        const { valid, errors }: ValidationResult = await this.validate()

        if (!valid) {
          this.errors = errors
          return
        }

        this.status = Status.Sending
        try {
          await ax.post(
            'http://localhost:1323/api/messages',
            { name, email, body },
          )
          this.errors = []
          this.status = Status.Sent
        } catch ({ response: { data: { errors } } }) {
          this.errors = errors
          this.status = Status.Ready
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
