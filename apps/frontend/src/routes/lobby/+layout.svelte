<script>
  let { children } = $props();
</script>

<svelte:head>
  <link rel="preload" as="image" href="/logo.png" />
</svelte:head>

<header class="fixed top-0 left-0 right-0 bg-white border-b border-gray-200 z-50 h-24">
  <div class="max-w-6xl mx-auto px-6 h-full flex items-center justify-between">
    
    <button type="button" class="flex items-center gap-3 hover:opacity-90 transition" onclick={() => location.reload()}>
      <img 
        src="/logo.png" 
        alt="Ollm Logo"
        width="140" 
        height="56"
        class="h-14 w-auto object-contain"
      /> 
    </button>

    <div class="flex items-center gap-3">
      <span class="text-sm font-bold text-gray-500">Player1</span>
      <div class="w-10 h-10 rounded-full bg-gray-200 border border-gray-300 overflow-hidden">
         <svg class="w-full h-full text-gray-400 p-1" fill="currentColor" viewBox="0 0 24 24"><path d="M24 20.993V24H0v-2.996A14.977 14.977 0 0112.004 15c4.904 0 9.26 2.354 11.996 5.993zM16.002 8.999a4 4 0 11-8 0 4 4 0 018 0z" /></svg>
      </div>
    </div>

  </div>
</header>

<main class="pt-24 pb-12">
  {@render children()}
</main>