import Alpine from 'alpinejs';

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
        .then(res => res.json())
        .then(data => {
          // set cookie here
          console.log(data);
          
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
          // set cookie here
          console.log(data);
          if (data.error) {
            this.title = 'User log in unsuccessful!';
            this.type = 'danger';
            this.popToast();
          } else {
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
});
