<script lang="ts">
  import Navbar from '$lib/components/navbar.svelte';

  let { children } = $props();
</script>

<svelte:head>
  <title>Eventually Assembled</title>
</svelte:head>

<div class="container">
  <Navbar />
</div>

<section class="section">
  <div class="container">
    <h1 class="title">Welcome to Eventually Assembled</h1>
    <h2 class="subtitle">Your furniture. Some assembly required. Eventually.</h2>

    {@render children()}
  </div>
</section>
