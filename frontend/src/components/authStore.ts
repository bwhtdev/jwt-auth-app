/**
 * Auth Store
 * Exports:
 * * username (store)
 * * loggedIn (store)
 */
import { writable } from 'svelte/store';

export const username = writable('');

export const loggedIn  = writable(false);
