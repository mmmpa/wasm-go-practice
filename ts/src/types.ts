export type Message = {
  name: string,
  email: string,
  body: string
}

export type FieldName = string

export type ErrorMessages = { [key: string]: string[] }

export type CompilationResult = {
  fetchingTime: number,
  compilationTime: number,
  go: any,
}

export type MessageCreationResult = {
  id: number,
}
