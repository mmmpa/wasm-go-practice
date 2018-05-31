export type Message = {
  name: string,
  email: string,
  body: string
}

export type FieldName = string

export type ErrorDetail = {
  field: FieldName,
  reason: string,
}

export type ValidationResult = {
  valid: boolean,
  message: string,
  errors: ErrorDetail[]
}
