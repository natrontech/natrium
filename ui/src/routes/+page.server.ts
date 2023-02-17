import { currentUser } from '$lib/pocketbase';
import { redirect } from '@sveltejs/kit';

// if user is not logged in, redirect to login page
export function load() {
	if (!currentUser) {
		throw redirect(303, '/login');
	}
}
