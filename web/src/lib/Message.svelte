<script>
  import { onMount } from "svelte";

  let message = {};

  onMount(async () => {
    const res = await fetch(`http://localhost:8080/api/v1/`);
    message = await res.json();
  });

  async function refresh() {
    const res = await fetch(`http://localhost:8080/api/v1/`);
    message = await res.json();
    console.log("refresh")
  }
</script>

<div>
  {#await message}
    <!-- promise is pending -->
    <p>waiting for the promise to resolve...</p>
  {:then value}
    <!-- promise was fulfilled -->
    <p>message: {value.message}</p>
    <p>version: {value.version}</p>
  {:catch error}
    <!-- promise was rejected -->
    <p>Something went wrong: {error.message}</p>
  {/await}
  <button on:click={refresh} >Refresh</button>
</div>
