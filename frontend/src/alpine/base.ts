import Alpine from 'alpinejs';
import { setToken, getToken/*, eraseToken*/, getUsername, setUsername/*, eraseUsername*/ } from './authUtils';

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
            this.title = 'User log in unsuccessful!';
            this.type = 'danger';
            this.popToast();
          } else {
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
        body: JSON.stringify({ text: this.messageText }),
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
            // Refresh data??
            
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
});
