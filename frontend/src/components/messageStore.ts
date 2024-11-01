/**
 * Message Store
 * Exports:
 * * Message (type)
 * * messageData (store)
 * * messageStatus (store)
 * * reloadMessageData()
 * * addNewMessage(Message)
 * * updateMessage(string, Message)
 * * removeMessage(string)
 */
import { get, writable } from 'svelte/store';

export type Message = {
  id: string;
  text: string;
  username: string;
  createdAt: string;
};

export const messageData = writable([]);

export const messageStatus = writable('complete');

export function reloadMessageData() {
  messageStatus.set('loading');
  
  fetch('/api/v1/messages', {
    headers: {'Access-Control-Allow-Origin': '*'}
  })
    .then(res => res.json())
    .then(data => {
      messageData.set(data);
      messageStatus.set('complete');
    })
    .catch(err => {
      console.log(err);
      messageStatus.set('error');
    });
}

export function addNewMessage(msg: Message) {
  messageData.set([ ...get(messageData), msg ]);
}

export function updateMessage(id: string, msg: Message) {
  let msgs = get(messageData);
  let index = msgs.findIndex((msg: Message) => msg.id == id);
  msgs[index] = msg;
  messageData.set(msgs);
}

export function removeMessage(id: string) {
  let msgs = get(messageData);
  let index = msgs.findIndex((msg: Message) => msg.id == id);
  msgs.splice(index, 1);
  messageData.set(msgs);
}
