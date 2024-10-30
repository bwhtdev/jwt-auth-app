import Alpine from 'alpinejs';

document.addEventListener('alpine:init', () => {
  Alpine.store('modalOpen', false);
  
  Alpine.data('signUpData', () => ({
    signUp() {
      this.title = 'You are successfully signed up!!';
      this.popToast();
      Alpine.store('modalOpen', false);
    },
  }));

  Alpine.data('logInData', () => ({
    logIn() {
      this.title = 'You are successfully logged in!!';
      this.popToast();
      Alpine.store('modalOpen', false);
    }
  }));
});
