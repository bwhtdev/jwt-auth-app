/**
 * Utils for Alpine auth scripts
 * Fns:
 * * isLoggedIn(): boolean
 * * logOut()
 *
 * * getToken(): string
 * * setToken(string)
 * * clearToken()
 *
 * * getUsername(): string
 * * setUsername(string)
 * * clearUsername()
 */

import { username, loggedIn } from '@components/authStore';

export function isLoggedIn(): boolean {
  const c = document.cookie;
  const parts = c.split(';');
  for (const part of parts)
    if (part.trim().startsWith('BEARER_TOKEN=') || part.trim().startsWith('USERNAME=')) {
      loggedIn.set(true);
      username.set(getUsername());
      return true;
    }
  loggedIn.set(false);
  return false;
}

export function logOut() {
  eraseToken();
  eraseUsername();
  loggedIn.set(false);
}

export function setToken(value: string) {
  var expires = "";
  var date = new Date();
  // 12 days expiration
  date.setTime(date.getTime() + (12 *24*60*60*1000));
  expires = "; expires=" + date.toUTCString();
  document.cookie = "BEARER_TOKEN=" + value  + expires + "; path=/";
}

export function getToken(): string | null {
  var nameEQ = "BEARER_TOKEN=";
  var ca = document.cookie.split(';');
  for (var i = 0; i < ca.length; i++) {
    var c = ca[i];
    while (c.charAt(0) == ' ') c = c.substring(1, c.length);
    if (c.indexOf(nameEQ) == 0) return c.substring(nameEQ.length, c.length);
  }
  return null;
}

export function eraseToken() {   
  document.cookie = 'BEARER_TOKEN=; Path=/; Expires=Thu, 01 Jan 1970 00:00:01 GMT;';
}

export function setUsername(value: string) {
  var expires = "";
  var date = new Date();
  // 12 days expiration
  date.setTime(date.getTime() + (12 *24*60*60*1000));
  expires = "; expires=" + date.toUTCString();
  document.cookie = "USERNAME=" + value  + expires + "; path=/";
}

export function getUsername(): string | null {
  var nameEQ = "USERNAME=";
  var ca = document.cookie.split(';');
  for (var i = 0; i < ca.length; i++) {
      var c = ca[i];
      while (c.charAt(0) == ' ') c = c.substring(1, c.length);
      if (c.indexOf(nameEQ) == 0) return c.substring(nameEQ.length, c.length);
  }
  return null;
}

export function eraseUsername() {   
  document.cookie = 'USERNAME=; Path=/; Expires=Thu, 01 Jan 1970 00:00:01 GMT;';
}
