<template lang="pug">
  article(v-if="!sent")
    .col-sm-4.col-sm-offset-4
      h1 Message Form
      form(@submit.prevent="submit")
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
  article(v-else)
    .col-sm-4.col-sm-offset-4
      h1 Thank you.
      .form-group.top-space
        button.btn.btn-success(@click="reset")
          i.fa.fa-envelope
          | Create new message
</template>

<script>
  import ax from 'axios'
  import ErrorFormWrapper from './ErrorFormWrapper'
  import snakeCase from 'snake-case'
  import loadAndCompile from '../wasm'

  function initialData () {
    return {
      name: '',
      email: '',
      body: '',
      errors: {},
      status: false,
    }
  }

  export default {
    name: 'HelloWorld',
    components: {
      ErrorFormWrapper
    },
    async created () {
      this.status = 'wasm'
      this.go = await loadAndCompile().catch(console.error)
      this.status = ''
    },
    data () {
      return initialData()
    },
    computed: {
      disabled () {
        return this.preparing || this.sending
      },
      sent () {
        return this.status === 'sent'
      },
      sending () {
        return this.status === 'sending'
      },
      preparing () {
        return this.status === 'wasm'
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
            JSON.stringify({ name, email, body })
          )
        )
      },
      parseErrors (raw) {
        this.errors = raw.reduce((a, { field, reason }) => {
          const key = snakeCase(field)
          const store = a[key] || (a[key] = [])
          store.push(reason)
          return a
        }, {})
      },
      async submit () {
        const {
          name,
          email,
          body,
        } = this

        const { valid, errors } = await this.validate()

        if (!valid) {
          this.parseErrors(errors)
          return
        }

        this.status = 'sending'
        try {
          await ax.post(
            'http://localhost:1323/api/messages',
            { name, email, body },
          )
          this.status = 'sent'
        } catch ({ response: { data: { errors } } }) {
          this.parseErrors(errors)
          this.status = ''
        }
      },
    },
  }
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
  button {
    width: 100%
  }

  .top-space {
    margin-top: 30px
  }

  button .fa {
    margin-right: 0.5em;
  }
</style>
