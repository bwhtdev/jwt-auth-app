import Alpine from 'alpinejs';
import { setToken, getToken, getUsername, setUsername } from './authUtils';
import { addNewMessage, updateMessage/*, removeMessage*/ } from '@components/messageStore';
import { username, loggedIn } from '@components/authStore';

document.addEventListener('alpine:init', () => {

  Alpine.store('modalOpen', false);
  
  Alpine.data('signUpData', () => ({
    signUp() {
      // Create new account for user and save cookies
      fetch('/api/v1/sign-up', {
        method: 'POST',
        body: JSON.stringify({ username: this.username, password: this.password }),
        headers: { 'Access-Control-Allow-Origin': '*' }
      })
        //.then(res => res.json())
        .then(() => {
          this.title = 'You are successfully signed up!!';
          this.type = 'success';
          this.popToast();

          Alpine.store('modalOpen', false);
          this.username = '';
          this.password = '';
        })
        .catch(err => {
          this.title = 'User creation unsuccessful!';
          this.description = err;
          this.type = 'danger';
          this.popToast();
        });
    },
  }));

  Alpine.data('logInData', () => ({
    logIn() {
      // Log in user and save cookies
      fetch('/api/v1/log-in', {
        method: 'POST',
        body: JSON.stringify({ username: this.username, password: this.password }),
        headers: { 'Access-Control-Allow-Origin': '*' }
      })
        .then(res => res.json())
        .then(data => {
          if (data.error) {
            loggedIn.set(true);
            username.set(this.username);

            this.title = 'User log in unsuccessful!';
            this.type = 'danger';
            this.popToast();
          } else {
            loggedIn.set(true);
            username.set(data.username);

            setToken(data.token);
            setUsername(data.username);
            
            this.title = 'User log in successful!';
            this.type = 'success';
            this.popToast();

            Alpine.store('modalOpen', false);
            this.username = '';
            this.password = '';
          }
        })
        .catch(err => {
          this.title = 'User log in unsuccessful!';
          this.description = err;
          this.type = 'danger';
          this.popToast();
        });
    },
  }));

  Alpine.data('createMessageData', () => ({
    createMessage() {
      const username = getUsername();
      const token = getToken();

      fetch(`/api/v1/message/new/${username}`, {
        method: 'POST',
        body: JSON.stringify({ text: this.messageText, username }),
        headers: {
          'Access-Control-Allow-Origin': '*',
          'Authorization': `Bearer ${token}`
        }
      })
        .then(res => res.json())
        .then(data => {
          if (data.error) {
            this.title = 'Message unsuccessful!';
            this.type = 'danger';
            this.popToast();
          } else {
            addNewMessage(data);
            
            this.title = 'Message created successfully!';
            this.type = 'success';
            this.popToast();

            Alpine.store('modalOpen', false);
            this.messageText = '';
          }
        })
        .catch(err => {
          this.title = 'Message unsuccessful!';
          this.description = err;
          this.type = 'danger';
          this.popToast();
        });
    },
  }));

  Alpine.data('editMessageData', () => ({
    editMessage() {
      const username = getUsername();
      const token = getToken();

      fetch(`/api/v1/message/${username}`, {
        method: 'POST',
        body: JSON.stringify({ id: this.messageId, text: this.messageText, username }),
        headers: {
          'Access-Control-Allow-Origin': '*',
          'Authorization': `Bearer ${token}`
        }
      })
        .then(res => res.json())
        .then(data => {
          if (data.error) {
            this.title = 'Message editing unsuccessful!';
            this.type = 'danger';
            this.popToast();
          } else {
            updateMessage(data.id, data);
            
            this.title = 'Message edited successfully!';
            this.type = 'success';
            this.popToast();

            Alpine.store('modalOpen', false);
          }
        })
        .catch(err => {
          this.title = 'Message editing unsuccessful!';
          this.description = err;
          this.type = 'danger';
          this.popToast();
        });
    },
  }));
});
