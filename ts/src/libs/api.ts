import ax from 'axios'
import { Message, MessageCreationResult } from '../types'
import { ValidationResult } from '../../validation/assembly/types'

export async function createMessage (message: Message): Promise<MessageCreationResult> {
  return ax.post(
    'http://localhost:1323/api/messages',
    message,
  )
    .then(({ data }) => data as MessageCreationResult)
    .catch(({ response: { data } }) => Promise.reject({ data }))
}
