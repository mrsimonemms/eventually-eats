<script lang="ts">
  import Navbar from '$lib/components/navbar.svelte';

  let { children } = $props();
</script>

<svelte:head>
  <title>Eventually Eats</title>
</svelte:head>

<div class="container">
  <Navbar />
</div>

<section class="section">
  <div class="container">
    <h1 class="title">Welcome to Eventually Eats</h1>
    <h2 class="subtitle">Your food. Exactly once. Eventually.</h2>

    {@render children()}
  </div>
</section>
