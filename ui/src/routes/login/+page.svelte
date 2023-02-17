<script lang="ts">
	import { enhance } from '$app/forms';
	import toast from 'svelte-french-toast';
  export let form;
	let loading = false;
	const submitLogin = () => {
		loading = true;
		return async ({ result, update }: any) => {
			switch (result.type) {
				case 'success':
					await update();
					break;
				case 'invalid':
					toast.error('Invalid credentials');
					await update();
					break;
				case 'error':
					toast.error(result.error.message);
					break;
				default:
					await update();
			}
			loading = false;
		};
	};
</script>

<!-- <div class="max-w-xs">
			<img src="/images/natrium-logo.png" alt="Logo" width="100%" />
		</div> -->
<div class="card flex-shrink-0 w-full max-w-sm shadow-2xl bg-base-100">
	<div class="card-body">
		<form
      action="?/login"
			method="POST"
			class="card"
      use:enhance={() => {
        return submitLogin();
      }}
		>
			<div class="form-control">
				<label class="label">
					<span class="label-text">Email</span>
				</label>
				<input
					type="email"
					name="email"
					placeholder="Email"
					class="input input-bordered"
					autocomplete="off"
          disabled={loading}
				/>
			</div>
			<div class="form-control">
				<label class="label">
					<span class="label-text">Password</span>
				</label>
				<input
					type="password"
					name="password"
					placeholder="Password"
					class="input input-bordered"
					autocomplete="off"
          disabled={loading}
				/>
			</div>
			<div class="form-control mt-6">
				<button class="btn btn-primary" disabled={loading}>Login</button>
			</div>
		</form>
	</div>
</div>
